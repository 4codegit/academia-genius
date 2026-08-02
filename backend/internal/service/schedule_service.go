package service

import (
    "context"

    "academy-genius/internal/models"
    "academy-genius/internal/repository"
)

type ScheduleService struct {
    repo *repository.ScheduleRepo
}

func NewScheduleService(repo *repository.ScheduleRepo) *ScheduleService {
    return &ScheduleService{repo: repo}
}

func (s *ScheduleService) List(ctx context.Context) ([]models.Webinar, error) {
    return s.repo.List(ctx)
}