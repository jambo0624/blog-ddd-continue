package config

import (
    "fmt"
    "net/url"
    "os"
    "github.com/spf13/viper"
)

type Config struct {
    Database DatabaseConfig
    Server   ServerConfig
}

type DatabaseConfig struct {
    URL      string // 完整的数据库URL
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

    viper.SetConfigName(fmt.Sprintf(".env.%s", env))
    viper.SetConfigType("env")
    viper.AddConfigPath(".")
    viper.AutomaticEnv()

    // 设置默认值
    viper.SetDefault("SERVER_PORT", "8080")
    
    if err := viper.ReadInConfig(); err != nil {
        if env != "production" {
            return nil, fmt.Errorf("error reading config file: %s", err)
        }
    }

    config := &Config{
        Server: ServerConfig{
            Port: viper.GetString("SERVER_PORT"),
        },
    }

    // 优先使用 DATABASE_URL
    if dbURL := viper.GetString("DATABASE_URL"); dbURL != "" {
        config.Database = ParseDatabaseURL(dbURL)
    } else {
        // 如果没有 DATABASE_URL，使用独立的配置项
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
        DBName:   u.Path[1:], // 移除开头的 "/"
        SSLMode:  "disable",  // 默认值
    }
} 