package middleware

import (
    "net/http"
    "strings"

    "academy-genius/internal/service"

    "github.com/gin-gonic/gin"
)

// Auth проверяет JWT-токен в заголовке Authorization.
// Если токен валиден — записывает user_id в контекст.
func Auth(authService *service.AuthService) gin.HandlerFunc {
    return func(c *gin.Context) {
        header := c.GetHeader("Authorization")
        if header == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "требуется авторизация"})
            c.Abort()
            return
        }

        parts := strings.SplitN(header, " ", 2)
        if len(parts) != 2 || parts[0] != "Bearer" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "неверный формат токена"})
            c.Abort()
            return
        }

        userID, err := authService.ParseToken(parts[1])
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "невалидный или просроченный токен"})
            c.Abort()
            return
        }

        c.Set("user_id", userID)
        c.Next()
    }
}