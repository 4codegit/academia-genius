package service

import (
	"context"
	"math"

	"academy-genius/internal/models"
	"academy-genius/internal/repository"
)

type NewsService struct {
	repo *repository.NewsRepo
}

func NewNewsService(repo *repository.NewsRepo) *NewsService {
	return &NewsService{repo: repo}
}

func (s *NewsService) List(ctx context.Context, page, limit int) (*models.Pagination, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 50 {
		limit = 10
	}

	items, total, err := s.repo.List(ctx, page, limit)
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

func (s *NewsService) GetByID(ctx context.Context, id int64) (*models.News, error) {
	return s.repo.GetByID(ctx, id)
}
