package repository

import (
        "context"
        "fmt"

        "academy-genius/internal/models"

        "github.com/jackc/pgx/v5/pgxpool"
)

type BookRepo struct {
        pool *pgxpool.Pool
}

func NewBookRepo(pool *pgxpool.Pool) *BookRepo {
        return &BookRepo{pool: pool}
}

func (r *BookRepo) List(ctx context.Context, category string, page, limit int) ([]models.Book, int64, error) {
        var total int64
        var args []interface{}
        argNum := 1

        where := ""
        if category != "" {
                where = "WHERE category = $1"
                args = append(args, category)
                argNum++
        }

        countQuery := "SELECT COUNT(*) FROM books " + where
        err := r.pool.QueryRow(ctx, countQuery, args...).Scan(&total)
        if err != nil {
                return nil, 0, err
        }

        offset := (page - 1) * limit
        args = append(args, limit, offset)

        query := "SELECT id, title, author, category, description, cover_url, year, download_url " +
                "FROM books " + where + " ORDER BY year DESC " +
                "LIMIT $" + fmt.Sprintf("%d", argNum) + " OFFSET $" + fmt.Sprintf("%d", argNum+1)

        rows, err := r.pool.Query(ctx, query, args...)
        if err != nil {
                return nil, 0, err
        }
        defer rows.Close()

        var items []models.Book
        for rows.Next() {
                var b models.Book
                if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Category, &b.Description, &b.CoverURL, &b.Year, &b.DownloadURL); err != nil {
                        return nil, 0, err
                }
                items = append(items, b)
        }
        return items, total, nil
}
