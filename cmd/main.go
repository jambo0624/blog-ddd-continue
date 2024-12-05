package main

import (
    "log"
    "github.com/jambo0624/blog/internal/shared/infrastructure/config"
    "github.com/jambo0624/blog/internal/shared/infrastructure/persistence"
    "github.com/jambo0624/blog/internal/bootstrap"
)

func main() {
    // 加载配置
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    // 初始化数据库
    db, err := persistence.InitDB(cfg)
    if err != nil {
        log.Fatalf("Failed to initialize database: %v", err)
    }

    // 初始化各层
    repos := bootstrap.SetupRepositories(db)
    services := bootstrap.SetupServices(repos)
    handlers := bootstrap.SetupHandlers(services)
    router := bootstrap.SetupRouter(handlers)

    // 启动服务器
    log.Fatal(router.Run(":" + cfg.Server.Port))
} 