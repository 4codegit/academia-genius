package repository

import (
	"context"

	"academy-genius/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type NewsRepo struct {
	pool *pgxpool.Pool
}

func NewNewsRepo(pool *pgxpool.Pool) *NewsRepo {
	return &NewsRepo{pool: pool}
}

func (r *NewsRepo) List(ctx context.Context, page, limit int) ([]models.News, int64, error) {
	var total int64
	err := r.pool.QueryRow(ctx, `SELECT COUNT(*) FROM news WHERE is_active = true`).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	rows, err := r.pool.Query(ctx,
		`SELECT id, title, slug, summary, content, image_url, published_at, is_active, created_at
		FROM news WHERE is_active = true
		ORDER BY published_at DESC
		LIMIT $1 OFFSET $2`, limit, offset,
	)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []models.News
	for rows.Next() {
		var n models.News
		if err := rows.Scan(
			&n.ID, &n.Title, &n.Slug, &n.Summary, &n.Content,
			&n.ImageURL, &n.PublishedAt, &n.IsActive, &n.CreatedAt,
		); err != nil {
			return nil, 0, err
		}
		items = append(items, n)
	}
	return items, total, nil
}

func (r *NewsRepo) GetByID(ctx context.Context, id int64) (*models.News, error) {
	n := &models.News{}
	query := `SELECT id, title, slug, summary, content, image_url, published_at, is_active, created_at
		FROM news WHERE id = $1`
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&n.ID, &n.Title, &n.Slug, &n.Summary, &n.Content,
		&n.ImageURL, &n.PublishedAt, &n.IsActive, &n.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return n, nil
}
