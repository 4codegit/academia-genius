package service

import (
	"context"
	"math"

	"academy-genius/internal/models"
	"academy-genius/internal/repository"
)

type BookService struct {
	repo *repository.BookRepo
}

func NewBookService(repo *repository.BookRepo) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) List(ctx context.Context, category string, page, limit int) (*models.Pagination, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 50 {
		limit = 12
	}

	items, total, err := s.repo.List(ctx, category, page, limit)
	if err != nil {
		return nil, err
	}

	return &models.Pagination{
		Data:       items,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: int(math.Ceil(float64(total) / float64(limit))),
	}, nil
}
