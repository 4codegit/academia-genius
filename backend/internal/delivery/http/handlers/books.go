package handlers

import (
	"net/http"

	"academy-genius/internal/service"

	"github.com/gin-gonic/gin"
)

// BookHandler обрабатывает запросы к книгам.
type BookHandler struct {
	svc *service.BookService
}

// NewBookHandler создаёт новый BookHandler.
func NewBookHandler(svc *service.BookService) *BookHandler {
	return &BookHandler{svc: svc}
}

// List возвращает список книг с фильтрацией по категории и пагинацией.
// Query-параметры:
//   - category — категория книги (необязательный)
//   - page     — номер страницы (по умолчанию 1)
//   - limit    — размер страницы (по умолчанию 12)
func (h *BookHandler) List(c *gin.Context) {
	category := c.Query("category")
	page, limit := parsePagination(c, 12)

	resp, err := h.svc.List(c.Request.Context(), category, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
