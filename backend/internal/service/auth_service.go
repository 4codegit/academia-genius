package service

import (
	"context"
	"errors"
	"time"

	"academy-genius/internal/models"
	"academy-genius/internal/repository"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserAlreadyExists   = errors.New("пользователь уже существует")
	ErrInvalidCredentials  = errors.New("неверный email или пароль")
	ErrUserNotFound        = errors.New("пользователь не найден")
)

type AuthService struct {
	repo      *repository.UserRepo
	jwtSecret string
}

func NewAuthService(repo *repository.UserRepo, jwtSecret string) *AuthService {
	return &AuthService{repo: repo, jwtSecret: jwtSecret}
}

// Register создаёт нового пользователя и возвращает JWT-токен.
func (s *AuthService) Register(ctx context.Context, req models.RegisterRequest) (*models.AuthResponse, error) {
	exists, err := s.repo.ExistsByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrUserAlreadyExists
	}

	exists, err = s.repo.ExistsByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrUserAlreadyExists
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: string(hash),
		FullName:     req.FullName,
		Role:         "student",
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	// Инициализируем карту знаний
	if err := s.repo.InitKnowledgeMap(ctx, user.ID); err != nil {
		// Не фатально, логируем но продолжаем
	}

	token, err := s.generateToken(user)
	if err != nil {
		return nil, err
	}

	user.PasswordHash = "" // Не отправляем хеш
	return &models.AuthResponse{Token: token, User: *user}, nil
}

// Login проверяет учётные данные и возвращает JWT-токен.
func (s *AuthService) Login(ctx context.Context, req models.LoginRequest) (*models.AuthResponse, error) {
	user, err := s.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	token, err := s.generateToken(user)
	if err != nil {
		return nil, err
	}

	user.PasswordHash = ""
	return &models.AuthResponse{Token: token, User: *user}, nil
}

// ParseToken парсит JWT-токен и возвращает userID.
func (s *AuthService) ResetPassword(ctx context.Context, req models.ResetPasswordRequest) error {
	user, err := s.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		return ErrUserNotFound
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	if err := s.repo.UpdatePasswordByEmail(ctx, user.Email, string(hash)); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrUserNotFound
		}
		return err
	}

	return nil
}

func (s *AuthService) Me(ctx context.Context, userID int64) (*models.User, error) {
	user, err := s.repo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	user.PasswordHash = ""
	return user, nil
}

func (s *AuthService) ParseToken(tokenStr string) (int64, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("неожиданный метод подписи")
		}
		return []byte(s.jwtSecret), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errors.New("невалидный токен")
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("невалидный токен")
	}

	return int64(userID), nil
}

func (s *AuthService) generateToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(7 * 24 * time.Hour).Unix(),
		"iat":      time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtSecret))
}
