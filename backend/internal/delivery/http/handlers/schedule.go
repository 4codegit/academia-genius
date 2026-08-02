package handlers

import (
	"net/http"

	"academy-genius/internal/service"

	"github.com/gin-gonic/gin"
)

// ScheduleHandler обрабатывает запросы к расписанию вебинаров.
type ScheduleHandler struct {
	svc *service.ScheduleService
}

// NewScheduleHandler создаёт новый ScheduleHandler.
func NewScheduleHandler(svc *service.ScheduleService) *ScheduleHandler {
	return &ScheduleHandler{svc: svc}
}

// List возвращает список запланированных вебинаров.
func (h *ScheduleHandler) List(c *gin.Context) {
	webinars, err := h.svc.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, webinars)
}
