package http

import (
	"academy-genius/internal/delivery/http/handlers"
	"academy-genius/internal/delivery/http/middleware"
	"academy-genius/internal/service"

	"github.com/gin-gonic/gin"
)

// NewRouter настраивает все маршруты API и возвращает настроенный *gin.Engine.
func NewRouter(services *service.Services, jwtSecret string) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORS())

	// Инициализация обработчиков
	authHandler := handlers.NewAuthHandler(services.Auth)
	newsHandler := handlers.NewNewsHandler(services.News)
	problemHandler := handlers.NewProblemHandler(services.Problem)
	courseHandler := handlers.NewCourseHandler(services.Course)
	bookHandler := handlers.NewBookHandler(services.Book)
	alumniHandler := handlers.NewAlumniHandler(services.Alumni)
	scheduleHandler := handlers.NewScheduleHandler(services.Schedule)
	userHandler := handlers.NewUserHandler(services.User)

	// Публичные маршруты
	api := r.Group("/api")
	{
		api.POST("/auth/register", authHandler.Register)
		api.POST("/auth/login", authHandler.Login)
		api.POST("/auth/forgot-password", authHandler.ForgotPassword)

		api.GET("/news", newsHandler.List)
		api.GET("/news/:id", newsHandler.GetByID)

		api.GET("/problems", problemHandler.List)
		api.GET("/courses", courseHandler.List)
		api.GET("/books", bookHandler.List)
		api.GET("/alumni", alumniHandler.GetAll)
		api.GET("/schedule", scheduleHandler.List)

		// Защищённые маршруты (требуют JWT)
		protected := api.Group("/")
		protected.Use(middleware.Auth(services.Auth))
		{
			protected.GET("/user/stats", userHandler.GetStats)
			protected.GET("/auth/me", authHandler.Me)
			protected.POST("/auth/logout", authHandler.Logout)
		}
	}

	return r
}
