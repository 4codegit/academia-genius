package handlers

import (
	"net/http"

	"academy-genius/internal/service"

	"github.com/gin-gonic/gin"
)

// AlumniHandler обрабатывает запросы к выпускникам.
type AlumniHandler struct {
	svc *service.AlumniService
}

// NewAlumniHandler создаёт новый AlumniHandler.
func NewAlumniHandler(svc *service.AlumniService) *AlumniHandler {
	return &AlumniHandler{svc: svc}
}

// GetAll возвращает избранного выпускника (с наградами) и список остальных.
func (h *AlumniHandler) GetAll(c *gin.Context) {
	resp, err := h.svc.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
