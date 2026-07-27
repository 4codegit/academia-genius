package config

import (
        "fmt"
        "net/url"
        "os"
)

// Config хранит все конфигурационные параметры приложения.
type Config struct {
        Port      string
        DBURL     string
        JWTSecret string
}

// Load загружает конфигурацию из переменных окружения.
// Если переменная не задана — используется значение по умолчанию.
func Load() *Config {
        port := getEnv("PORT", "8080")
        jwtSecret := getEnv("JWT_SECRET", "academy-genius-secret-key-change-in-production")

        dbURL := os.Getenv("DATABASE_URL")
        if dbURL == "" {
                dbHost := getEnv("DB_HOST", "localhost")
                dbPort := getEnv("DB_PORT", "5432")
                dbUser := getEnv("DB_USER", "academy_genius")
                dbPassword := getEnv("DB_PASSWORD", "academia_genius_123")
                dbName := getEnv("DB_NAME", "academy_genius_db")
                sslMode := getEnv("DB_SSLMODE", "disable")

                dbURL = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
                        url.QueryEscape(dbUser),
                        url.QueryEscape(dbPassword),
                        dbHost,
                        dbPort,
                        dbName,
                        sslMode,
                )
        }

        return &Config{
                Port:      port,
                DBURL:     dbURL,
                JWTSecret: jwtSecret,
        }
}

func getEnv(key, fallback string) string {
        if v := os.Getenv(key); v != "" {
                return v
        }
        return fallback
}
