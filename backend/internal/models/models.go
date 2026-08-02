package models

import (
	"time"
)

// ===================== АВТОРИЗАЦИЯ =====================

type User struct {
	ID          int64     `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	PasswordHash string   `json:"-"`
	FullName    string    `json:"full_name"`
	Role        string    `json:"role"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=100"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required,min=1,max=200"`
}

type ResetPasswordRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

// ===================== НОВОСТИ =====================

type News struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug"`
	Summary     string    `json:"summary"`
	Content     string    `json:"content"`
	ImageURL    string    `json:"image_url"`
	PublishedAt time.Time `json:"published_at"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
}

// ===================== ЗАДАЧИ =====================

type Problem struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Topic     string    `json:"topic"`
	Difficulty string   `json:"difficulty"`
	Content   string    `json:"content"`
	Solution  string    `json:"solution"`
	ImageURL  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
}

// ===================== КУРСЫ =====================

type Course struct {
	ID         int64   `json:"id"`
	Title      string  `json:"title"`
	Description string `json:"description"`
	Instructor string  `json:"instructor"`
	ImageURL   string  `json:"image_url"`
	Price      float64 `json:"price"`
	Duration   string  `json:"duration"`
	IsActive   bool    `json:"is_active"`
}

// ===================== КНИГИ =====================

type Book struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Category    string `json:"category"`
	Description string `json:"description"`
	CoverURL    string `json:"cover_url"`
	Year        int    `json:"year"`
	DownloadURL string `json:"download_url"`
}

// ===================== ВЫПУСКНИКИ =====================

type Alumnus struct {
	ID             int64  `json:"id"`
	FullName       string `json:"full_name"`
	Bio            string `json:"bio"`
	PhotoURL       string `json:"photo_url"`
	GraduationYear int    `json:"graduation_year"`
	University     string `json:"university"`
	IsFeatured     bool   `json:"is_featured"`
	SortOrder      int    `json:"sort_order"`
}

type AlumniAward struct {
	ID          int64  `json:"id"`
	AlumniID    int64  `json:"alumni_id"`
	AwardTitle  string `json:"award_title"`
	Competition string `json:"competition"`
	Year        int    `json:"year"`
	Description string `json:"description"`
	SortOrder   int    `json:"sort_order"`
}

type FeaturedAlumnus struct {
	Alumnus
	Awards []AlumniAward `json:"awards"`
}

// ===================== РАСПИСАНИЕ =====================

type Webinar struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Speaker     string    `json:"speaker"`
	EventDate   time.Time `json:"event_date"`
	DurationMin int       `json:"duration_min"`
	PlatformURL string    `json:"platform_url"`
	IsActive    bool      `json:"is_active"`
}

// ===================== КАБИНЕТ =====================

type KnowledgeMapEntry struct {
	Topic    string `json:"topic"`
	Progress int    `json:"progress"`
}

type UserStats struct {
	TotalSolved  int                  `json:"total_solved"`
	ByTopic      map[string]int       `json:"by_topic"`
	ByDifficulty map[string]int       `json:"by_difficulty"`
	KnowledgeMap []KnowledgeMapEntry  `json:"knowledge_map"`
	StreakDays   int                  `json:"streak_days"`
	LastActive   string               `json:"last_active"`
}

// ===================== ПАГИНАЦИЯ =====================

type Pagination struct {
	Data       interface{} `json:"data"`
	Total      int64       `json:"total"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	TotalPages int         `json:"total_pages"`
}
