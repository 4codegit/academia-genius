package handlers

import (
	"net/http"
	"strconv"

	"academy-genius/internal/service"

	"github.com/gin-gonic/gin"
)

// NewsHandler обрабатывает запросы к новостям.
type NewsHandler struct {
	svc *service.NewsService
}

// NewNewsHandler создаёт новый NewsHandler.
func NewNewsHandler(svc *service.NewsService) *NewsHandler {
	return &NewsHandler{svc: svc}
}

// parsePagination извлекает page и limit из query-параметров.
// Если параметр отсутствует или некорректен — используется значение по умолчанию.
func parsePagination(c *gin.Context, defaultLimit int) (page, limit int) {
	page = 1
	if p := c.Query("page"); p != "" {
		if v, err := strconv.Atoi(p); err == nil && v > 0 {
			page = v
		}
	}

	limit = defaultLimit
	if l := c.Query("limit"); l != "" {
		if v, err := strconv.Atoi(l); err == nil && v > 0 {
			limit = v
		}
	}

	return page, limit
}

// List возвращает список новостей с пагинацией.
func (h *NewsHandler) List(c *gin.Context) {
	page, limit := parsePagination(c, 10)

	resp, err := h.svc.List(c.Request.Context(), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetByID возвращает новость по идентификатору.
func (h *NewsHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный ID новости"})
		return
	}

	news, err := h.svc.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "новость не найдена"})
		return
	}

	c.JSON(http.StatusOK, news)
}
