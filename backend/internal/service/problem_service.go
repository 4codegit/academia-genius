package service

import (
    "context"
    "math"
    "strings"

    "academy-genius/internal/models"
    "academy-genius/internal/repository"
)

type ProblemService struct {
    repo *repository.ProblemRepo
}

func NewProblemService(repo *repository.ProblemRepo) *ProblemService {
    return &ProblemService{repo: repo}
}

func (s *ProblemService) List(ctx context.Context, topicsStr, difficulty string, page, limit int) (*models.Pagination, error) {
    if page < 1 {
        page = 1
    }
    if limit < 1 || limit > 50 {
        limit = 20
    }

    var topics []string
    if topicsStr != "" {
        topics = strings.Split(topicsStr, ",")
        // Тримим пробелы
        for i := range topics {
            topics[i] = strings.TrimSpace(topics[i])
        }
    }

    items, total, err := s.repo.List(ctx, topics, difficulty, page, limit)
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

func (s *ProblemService) GetByID(ctx context.Context, id int64) (*models.Problem, error) {
    return s.repo.GetByID(ctx, id)
}
