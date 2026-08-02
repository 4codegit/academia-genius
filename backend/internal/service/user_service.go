package service

import (
	"context"

	"academy-genius/internal/models"
	"academy-genius/internal/repository"
)

type UserService struct {
	repo *repository.UserRepo
}

func NewUserService(repo *repository.UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetStats(ctx context.Context, userID int64) (*models.UserStats, error) {
	return s.repo.GetStats(ctx, userID)
}

func (s *UserService) GetByID(ctx context.Context, id int64) (*models.User, error) {
	return s.repo.GetByID(ctx, id)
}