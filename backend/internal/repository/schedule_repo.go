package repository

import (
	"context"

	"academy-genius/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ScheduleRepo struct {
	pool *pgxpool.Pool
}

func NewScheduleRepo(pool *pgxpool.Pool) *ScheduleRepo {
	return &ScheduleRepo{pool: pool}
}

func (r *ScheduleRepo) List(ctx context.Context) ([]models.Webinar, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, title, description, speaker, event_date, duration_min, platform_url, is_active
		FROM schedule WHERE is_active = true
		ORDER BY event_date ASC`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Webinar
	for rows.Next() {
		var w models.Webinar
		if err := rows.Scan(&w.ID, &w.Title, &w.Description, &w.Speaker, &w.EventDate, &w.DurationMin, &w.PlatformURL, &w.IsActive); err != nil {
			return nil, err
		}
		items = append(items, w)
	}
	return items, nil
}
