package service

import (
	"context"

	"academy-genius/internal/models"
	"academy-genius/internal/repository"
)

type CourseService struct {
	repo *repository.CourseRepo
}

func NewCourseService(repo *repository.CourseRepo) *CourseService {
	return &CourseService{repo: repo}
}

func (s *CourseService) List(ctx context.Context) ([]models.Course, error) {
	return s.repo.List(ctx)
}
