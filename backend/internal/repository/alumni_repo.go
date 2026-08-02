package repository

import (
	"context"

	"academy-genius/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AlumniRepo struct {
	pool *pgxpool.Pool
}

func NewAlumniRepo(pool *pgxpool.Pool) *AlumniRepo {
	return &AlumniRepo{pool: pool}
}

// GetFeatured возвращает featured выпускника с наградами (Пиров Далер).
func (r *AlumniRepo) GetFeatured(ctx context.Context) (*models.FeaturedAlumnus, error) {
	a := &models.Alumnus{}
	err := r.pool.QueryRow(ctx,
		`SELECT id, full_name, bio, photo_url, graduation_year, university, is_featured, sort_order
		FROM alumni WHERE is_featured = true
		ORDER BY sort_order LIMIT 1`,
	).Scan(&a.ID, &a.FullName, &a.Bio, &a.PhotoURL, &a.GraduationYear, &a.University, &a.IsFeatured, &a.SortOrder)
	if err != nil {
		return nil, err
	}

	rows, err := r.pool.Query(ctx,
		`SELECT id, alumni_id, award_title, competition, year, description, sort_order
		FROM alumni_awards WHERE alumni_id = $1
		ORDER BY year, sort_order`, a.ID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var awards []models.AlumniAward
	for rows.Next() {
		var aw models.AlumniAward
		if err := rows.Scan(&aw.ID, &aw.AlumniID, &aw.AwardTitle, &aw.Competition, &aw.Year, &aw.Description, &aw.SortOrder); err != nil {
			return nil, err
		}
		awards = append(awards, aw)
	}

	return &models.FeaturedAlumnus{
		Alumnus: *a,
		Awards:  awards,
	}, nil
}

// List возвращает всех остальных выпускников (не featured).
func (r *AlumniRepo) List(ctx context.Context) ([]models.Alumnus, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, full_name, bio, photo_url, graduation_year, university, is_featured, sort_order
		FROM alumni WHERE is_featured = false
		ORDER BY sort_order, graduation_year DESC`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Alumnus
	for rows.Next() {
		var a models.Alumnus
		if err := rows.Scan(&a.ID, &a.FullName, &a.Bio, &a.PhotoURL, &a.GraduationYear, &a.University, &a.IsFeatured, &a.SortOrder); err != nil {
			return nil, err
		}
		items = append(items, a)
	}
	return items, nil
}
