package config

import (
	"fmt"
	"github.com/spf13/viper"
	"net/url"
	"os"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
}

type DatabaseConfig struct {
	URL      string // Full database URL
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type ServerConfig struct {
	Port string
}

func LoadConfig() (*Config, error) {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "development"
	}

	// Add all possible config paths
	viper.AddConfigPath(".")
	viper.AddConfigPath("../../..")  // For tests
	viper.AddConfigPath("../../../..") // For deeper test directories
	
	viper.SetConfigName(fmt.Sprintf(".env.%s", env))
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	// Set default values
	viper.SetDefault("SERVER_PORT", "8080")
	viper.SetDefault("DATABASE_URL", "postgres://username:password@localhost:5432/databasename")

	if err := viper.ReadInConfig(); err != nil {
		if env != "production" {
			fmt.Printf("Warning: Config file not found, using defaults. Error: %v\n", err)
			// not production environment, use default value instead of returning error
			return &Config{
				Database: ParseDatabaseURL(viper.GetString("DATABASE_URL")),
				Server: ServerConfig{
					Port: viper.GetString("SERVER_PORT"),
				},
			}, nil
		}
		return nil, fmt.Errorf("error reading config file: %s", err)
	}

	config := &Config{
		Server: ServerConfig{
			Port: viper.GetString("SERVER_PORT"),
		},
	}

	// Use DATABASE_URL if available
	if dbURL := viper.GetString("DATABASE_URL"); dbURL != "" {
		config.Database = ParseDatabaseURL(dbURL)
	} else {
		// Use separate configuration items if DATABASE_URL is not available
		config.Database = DatabaseConfig{
			Host:     viper.GetString("DB_HOST"),
			Port:     viper.GetString("DB_PORT"),
			User:     viper.GetString("DB_USER"),
			Password: viper.GetString("DB_PASSWORD"),
			DBName:   viper.GetString("DB_NAME"),
			SSLMode:  viper.GetString("DB_SSL_MODE"),
		}
	}

	return config, nil
}

func ParseDatabaseURL(dbURL string) DatabaseConfig {
	u, err := url.Parse(dbURL)
	if err != nil {
		return DatabaseConfig{}
	}

	password, _ := u.User.Password()
	return DatabaseConfig{
		URL:      dbURL,
		Host:     u.Hostname(),
		Port:     u.Port(),
		User:     u.User.Username(),
		Password: password,
		DBName:   u.Path[1:], // Remove leading "/"
		SSLMode:  "disable",  // Default value
	}
}
