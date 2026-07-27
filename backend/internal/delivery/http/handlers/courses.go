package handlers

import (
	"net/http"

	"academy-genius/internal/service"

	"github.com/gin-gonic/gin"
)

// CourseHandler обрабатывает запросы к курсам.
type CourseHandler struct {
	svc *service.CourseService
}

// NewCourseHandler создаёт новый CourseHandler.
func NewCourseHandler(svc *service.CourseService) *CourseHandler {
	return &CourseHandler{svc: svc}
}

// List возвращает список всех доступных курсов.
func (h *CourseHandler) List(c *gin.Context) {
	courses, err := h.svc.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, courses)
}
