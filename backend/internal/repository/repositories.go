package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Repositories объединяет все репозитории.
type Repositories struct {
	User     *UserRepo
	News     *NewsRepo
	Problem  *ProblemRepo
	Course   *CourseRepo
	Book     *BookRepo
	Alumni   *AlumniRepo
	Schedule *ScheduleRepo
}

// NewRepositories создаёт все репозитории с общим пулом соединений.
func NewRepositories(pool *pgxpool.Pool) *Repositories {
	return &Repositories{
		User:     NewUserRepo(pool),
		News:     NewNewsRepo(pool),
		Problem:  NewProblemRepo(pool),
		Course:   NewCourseRepo(pool),
		Book:     NewBookRepo(pool),
		Alumni:   NewAlumniRepo(pool),
		Schedule: NewScheduleRepo(pool),
	}
}

// Ping проверяет доступность БД.
func (r *Repositories) Ping(ctx context.Context) error {
	return r.User.pool.Ping(ctx)
}
