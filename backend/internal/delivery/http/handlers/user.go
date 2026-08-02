package handlers

import (
	"net/http"

	"academy-genius/internal/service"

	"github.com/gin-gonic/gin"
)

// UserHandler обрабатывает запросы к личному кабинету пользователя.
type UserHandler struct {
	svc *service.UserService
}

// NewUserHandler создаёт новый UserHandler.
func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

// GetStats возвращает статистику текущего авторизованного пользователя.
// Идентификатор пользователя берётся из gin-контекста (устанавливается middleware.Auth).
func (h *UserHandler) GetStats(c *gin.Context) {
	val, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "пользователь не авторизован"})
		return
	}

	userID, ok := val.(int64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "невалидный идентификатор пользователя"})
		return
	}

	stats, err := h.svc.GetStats(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stats)
}
