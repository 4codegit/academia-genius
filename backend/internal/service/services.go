package service

import (
	"academy-genius/internal/repository"
)

// Services объединяет все сервисы.
type Services struct {
	Auth     *AuthService
	News     *NewsService
	Problem  *ProblemService
	Course   *CourseService
	Book     *BookService
	Alumni   *AlumniService
	Schedule *ScheduleService
	User     *UserService
}

// NewServices создаёт все сервисы.
func NewServices(repos *repository.Repositories, jwtSecret string) *Services {
	return &Services{
		Auth:     NewAuthService(repos.User, jwtSecret),
		News:     NewNewsService(repos.News),
		Problem:  NewProblemService(repos.Problem),
		Course:   NewCourseService(repos.Course),
		Book:     NewBookService(repos.Book),
		Alumni:   NewAlumniService(repos.Alumni),
		Schedule: NewScheduleService(repos.Schedule),
		User:     NewUserService(repos.User),
	}
}
