package repository

import (
	"context"

	"academy-genius/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CourseRepo struct {
	pool *pgxpool.Pool
}

func NewCourseRepo(pool *pgxpool.Pool) *CourseRepo {
	return &CourseRepo{pool: pool}
}

func (r *CourseRepo) List(ctx context.Context) ([]models.Course, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, title, description, instructor, image_url, price, duration, is_active
		FROM courses WHERE is_active = true
		ORDER BY created_at DESC`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Course
	for rows.Next() {
		var c models.Course
		if err := rows.Scan(&c.ID, &c.Title, &c.Description, &c.Instructor, &c.ImageURL, &c.Price, &c.Duration, &c.IsActive); err != nil {
			return nil, err
		}
		items = append(items, c)
	}
	return items, nil
}
