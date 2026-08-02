package repository

import (
        "context"
        "fmt"
        "strings"

        "academy-genius/internal/models"

        "github.com/jackc/pgx/v5/pgxpool"
)

type ProblemRepo struct {
        pool *pgxpool.Pool
}

func NewProblemRepo(pool *pgxpool.Pool) *ProblemRepo {
        return &ProblemRepo{pool: pool}
}

// List возвращает задачи с фильтрацией по темам и сложности.
func (r *ProblemRepo) List(ctx context.Context, topics []string, difficulty string, page, limit int) ([]models.Problem, int64, error) {
        var conditions []string
        var args []interface{}
        argNum := 1

        if len(topics) > 0 {
                placeholders := make([]string, len(topics))
                for i, t := range topics {
                        placeholders[i] = fmt.Sprintf("$%d", argNum)
                        args = append(args, t)
                        argNum++
                }
                conditions = append(conditions, fmt.Sprintf("topic IN (%s)", strings.Join(placeholders, ",")))
        }
        if difficulty != "" {
                conditions = append(conditions, fmt.Sprintf("difficulty = $%d", argNum))
                args = append(args, difficulty)
                argNum++
        }

        where := ""
        if len(conditions) > 0 {
                where = "WHERE " + strings.Join(conditions, " AND ")
        }

        // Считаем total
        var total int64
        countQuery := fmt.Sprintf("SELECT COUNT(*) FROM problems %s", where)
        err := r.pool.QueryRow(ctx, countQuery, args...).Scan(&total)
        if err != nil {
                return nil, 0, err
        }

        offset := (page - 1) * limit
        args = append(args, limit, offset)
        query := fmt.Sprintf(
                `SELECT id, title, topic, difficulty, content, solution, image_url, created_at
                FROM problems %s
                ORDER BY created_at DESC
                LIMIT $%d OFFSET $%d`,
                where, argNum, argNum+1,
        )

        rows, err := r.pool.Query(ctx, query, args...)
        if err != nil {
                return nil, 0, err
        }
        defer rows.Close()

        var items []models.Problem
        for rows.Next() {
                var p models.Problem
                if err := rows.Scan(
                        &p.ID, &p.Title, &p.Topic, &p.Difficulty,
                        &p.Content, &p.Solution, &p.ImageURL, &p.CreatedAt,
                ); err != nil {
                        return nil, 0, err
                }
                items = append(items, p)
        }
        return items, total, nil
}

func (r *ProblemRepo) GetByID(ctx context.Context, id int64) (*models.Problem, error) {
        p := &models.Problem{}
        query := `SELECT id, title, topic, difficulty, content, solution, image_url, created_at
                FROM problems WHERE id = $1`
        err := r.pool.QueryRow(ctx, query, id).Scan(
                &p.ID, &p.Title, &p.Topic, &p.Difficulty,
                &p.Content, &p.Solution, &p.ImageURL, &p.CreatedAt,
        )
        if err != nil {
                return nil, err
        }
        return p, nil
}
