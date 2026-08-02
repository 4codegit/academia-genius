package service

import (
    "context"

    "academy-genius/internal/models"
    "academy-genius/internal/repository"
)

type AlumniService struct {
    repo *repository.AlumniRepo
}

func NewAlumniService(repo *repository.AlumniRepo) *AlumniService {
    return &AlumniService{repo: repo}
}

type AlumniResponse struct {
    Featured models.FeaturedAlumnus `json:"featured"`
    Others   []models.Alumnus      `json:"others"`
}

func (s *AlumniService) GetAll(ctx context.Context) (*AlumniResponse, error) {
    featured, err := s.repo.GetFeatured(ctx)
    if err != nil {
        return nil, err
    }

    others, err := s.repo.List(ctx)
    if err != nil {
        return nil, err
    }

    return &AlumniResponse{
        Featured: *featured,
        Others:   others,
    }, nil
}