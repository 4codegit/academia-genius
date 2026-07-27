package repository

import (
        "context"
        "fmt"
        "strings"

        "academy-genius/internal/models"

        "github.com/jackc/pgx/v5"
        "github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
        pool *pgxpool.Pool
}

func NewUserRepo(pool *pgxpool.Pool) *UserRepo {
        return &UserRepo{pool: pool}
}

func (r *UserRepo) Create(ctx context.Context, user *models.User) error {
        query := `INSERT INTO users (username, email, password_hash, full_name, role)
                VALUES ($1, $2, $3, $4, $5)
                RETURNING id, created_at, updated_at`
        return r.pool.QueryRow(ctx, query,
                user.Username, user.Email, user.PasswordHash, user.FullName, user.Role,
        ).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
}

func (r *UserRepo) GetByEmail(ctx context.Context, email string) (*models.User, error) {
        u := &models.User{}
        query := `SELECT id, username, email, password_hash, full_name, role, created_at, updated_at
                FROM users WHERE email = $1`
        err := r.pool.QueryRow(ctx, query, email).Scan(
                &u.ID, &u.Username, &u.Email, &u.PasswordHash, &u.FullName, &u.Role, &u.CreatedAt, &u.UpdatedAt,
        )
        if err != nil {
                return nil, err
        }
        return u, nil
}

func (r *UserRepo) GetByID(ctx context.Context, id int64) (*models.User, error) {
        u := &models.User{}
        query := `SELECT id, username, email, full_name, role, created_at, updated_at
                FROM users WHERE id = $1`
        err := r.pool.QueryRow(ctx, query, id).Scan(
                &u.ID, &u.Username, &u.Email, &u.FullName, &u.Role, &u.CreatedAt, &u.UpdatedAt,
        )
        if err != nil {
                return nil, err
        }
        return u, nil
}

// GetStats возвращает статистику пользователя для кабинета.
func (r *UserRepo) GetStats(ctx context.Context, userID int64) (*models.UserStats, error) {
        stats := &models.UserStats{
                ByTopic:      make(map[string]int),
                ByDifficulty: make(map[string]int),
        }

        // Общее количество решённых задач
        err := r.pool.QueryRow(ctx,
                `SELECT COUNT(*) FROM problems_solved WHERE user_id = $1`, userID,
        ).Scan(&stats.TotalSolved)
        if err != nil {
                return nil, err
        }

        // По темам
        rows, err := r.pool.Query(ctx,
                `SELECT p.topic, COUNT(*)
                FROM problems_solved ps
                JOIN problems p ON p.id = ps.problem_id
                WHERE ps.user_id = $1
                GROUP BY p.topic`, userID,
        )
        if err != nil {
                return nil, err
        }
        defer rows.Close()
        for rows.Next() {
                var topic string
                var count int
                if err := rows.Scan(&topic, &count); err != nil {
                        return nil, err
                }
                stats.ByTopic[topic] = count
        }

        // По сложности
        rows, err = r.pool.Query(ctx,
                `SELECT p.difficulty, COUNT(*)
                FROM problems_solved ps
                JOIN problems p ON p.id = ps.problem_id
                WHERE ps.user_id = $1
                GROUP BY p.difficulty`, userID,
        )
        if err != nil {
                return nil, err
        }
        defer rows.Close()
        for rows.Next() {
                var diff string
                var count int
                if err := rows.Scan(&diff, &count); err != nil {
                        return nil, err
                }
                stats.ByDifficulty[diff] = count
        }

        // Карта знаний
        rows, err = r.pool.Query(ctx,
                `SELECT topic, progress FROM knowledge_map WHERE user_id = $1 ORDER BY topic`, userID,
        )
        if err != nil {
                return nil, err
        }
        defer rows.Close()
        for rows.Next() {
                var entry models.KnowledgeMapEntry
                if err := rows.Scan(&entry.Topic, &entry.Progress); err != nil {
                        return nil, err
                }
                stats.KnowledgeMap = append(stats.KnowledgeMap, entry)
        }

        // Последняя активность и стрик
        _ = r.pool.QueryRow(ctx,
                `SELECT COALESCE(MAX(solved_at::date)::text, '') FROM problems_solved WHERE user_id = $1`, userID,
        ).Scan(&stats.LastActive)

        // Простой стрик: количество уникальных дней за последние 30 дней подряд
        _ = r.pool.QueryRow(ctx,
                `SELECT COUNT(DISTINCT solved_at::date) FROM problems_solved
                WHERE user_id = $1 AND solved_at >= CURRENT_DATE - INTERVAL '30 days'`, userID,
        ).Scan(&stats.StreakDays)

        return stats, nil
}

// ExistsByEmail проверяет, есть ли пользователь с таким email.
func (r *UserRepo) ExistsByEmail(ctx context.Context, email string) (bool, error) {
        var exists bool
        err := r.pool.QueryRow(ctx, `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`, email).Scan(&exists)
        return exists, err
}

// ExistsByUsername проверяет, есть ли пользователь с таким username.
func (r *UserRepo) ExistsByUsername(ctx context.Context, username string) (bool, error) {
        var exists bool
        err := r.pool.QueryRow(ctx, `SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)`, username).Scan(&exists)
        return exists, err
}

func (r *UserRepo) UpdatePasswordByEmail(ctx context.Context, email, passwordHash string) error {
        tag, err := r.pool.Exec(ctx,
                `UPDATE users SET password_hash = $1, updated_at = now() WHERE email = $2`,
                passwordHash, email,
        )
        if err != nil {
                return err
        }
        if tag.RowsAffected() == 0 {
                return pgx.ErrNoRows
        }
        return nil
}

// GetSolvedProblemIDs возвращает множество ID решённых задач.
func (r *UserRepo) GetSolvedProblemIDs(ctx context.Context, userID int64) ([]int64, error) {
        rows, err := r.pool.Query(ctx,
                `SELECT problem_id FROM problems_solved WHERE user_id = $1`, userID,
        )
        if err != nil {
                return nil, err
        }
        defer rows.Close()
        var ids []int64
        for rows.Next() {
                var id int64
                if err := rows.Scan(&id); err != nil {
                        return nil, err
                }
                ids = append(ids, id)
        }
        return ids, nil
}

// topicsList используется для инициализации карты знаний нового пользователя.
var topicsList = []string{
        "Механика", "МКТ", "Термодинамика",
        "Электростатика", "Магнетизм", "Оптика",
        "СТО", "Квантовая",
}

// InitKnowledgeMap создаёт записи в knowledge_map для нового пользователя.
func (r *UserRepo) InitKnowledgeMap(ctx context.Context, userID int64) error {
        placeholders := make([]string, len(topicsList))
        args := make([]interface{}, 0, len(topicsList))
        for i, t := range topicsList {
                placeholders[i] = fmt.Sprintf("($%d, $%d)", 2*i+1, 2*i+2)
                args = append(args, userID, t)
        }
        query := fmt.Sprintf(
                `INSERT INTO knowledge_map (user_id, topic) VALUES %s ON CONFLICT DO NOTHING`,
                strings.Join(placeholders, ","),
        )
        _, err := r.pool.Exec(ctx, query, args...)
        return err
}
