package handlers

import (
	"net/http"

	"academy-genius/internal/service"

	"github.com/gin-gonic/gin"
)

// ProblemHandler обрабатывает запросы к задачам.
type ProblemHandler struct {
	svc *service.ProblemService
}

// NewProblemHandler создаёт новый ProblemHandler.
func NewProblemHandler(svc *service.ProblemService) *ProblemHandler {
	return &ProblemHandler{svc: svc}
}

// List возвращает список задач с фильтрацией (topics, difficulty) и пагинацией.
// Query-параметры:
//   - topics     — список тем через запятую (например "algebra,geometry")
//   - difficulty — сложность ("easy" | "medium" | "hard")
//   - page       — номер страницы (по умолчанию 1)
//   - limit      — размер страницы (по умолчанию 20)
func (h *ProblemHandler) List(c *gin.Context) {
	topics := c.Query("topics")
	difficulty := c.Query("difficulty")
	page, limit := parsePagination(c, 20)

	resp, err := h.svc.List(c.Request.Context(), topics, difficulty, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
