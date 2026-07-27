package config

import (
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
        return &Config{
                Port:      getEnv("PORT", "8080"),
                DBURL:     getEnv("DATABASE_URL", "postgres://z@localhost:5432/academy_genius"),
                JWTSecret: getEnv("JWT_SECRET", "academy-genius-secret-key-change-in-production"),
        }
}

func getEnv(key, fallback string) string {
        if v := os.Getenv(key); v != "" {
                return v
        }
        return fallback
}
