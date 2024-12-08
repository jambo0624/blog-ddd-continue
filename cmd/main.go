package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/getsentry/sentry-go"

	"github.com/jambo0624/blog/internal/bootstrap"
	"github.com/jambo0624/blog/internal/shared/infrastructure/config"
	"github.com/jambo0624/blog/internal/shared/infrastructure/persistence"
)

const sentryFlushTimeout = 2 * time.Second

// handleFatalError captures the error in Sentry and exits the program.
func handleFatalError(err error, msg string) {
	sentry.CaptureException(err)
	sentry.Flush(sentryFlushTimeout)
	log.Printf("%s: %v", msg, err)
	os.Exit(1)
}

func initSentry() {
	// To initialize Sentry's handler, you need to initialize Sentry itself beforehand
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:           "https://1655538451d9814e2fc548771815430d@o4507962276446208.ingest.de.sentry.io/4507962286473296",
		EnableTracing: true,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for tracing.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	}); err != nil {
		log.Printf("Sentry initialization failed: %v\n", err)
	}
}

func main() {
	initSentry()

	// Set up graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		handleFatalError(err, "Failed to load config")
	}

	// Initialize database
	db, err := persistence.InitDB(cfg)
	if err != nil {
		handleFatalError(err, "Failed to initialize database")
	}

	// Initialize each layer
	repos := bootstrap.SetupRepositories(db)
	services := bootstrap.SetupServices(repos)
	handlers := bootstrap.SetupHandlers(services)
	router := bootstrap.SetupRouter(handlers)

	// Start server in a new goroutine
	go func() {
		if err := router.Run(":" + cfg.Server.Port); err != nil {
			handleFatalError(err, "Failed to start server")
		}
	}()

	// Wait for interrupt signal
	<-quit
	log.Println("Shutting down server...")

	// Log only when exiting normally
	log.Println("Server exited")
}
