package auth

import (
	"context"
	"grpc-serveci-ref/internal/domain/models"
	"log/slog"
	"time"
)

type Auth struct {
	log          *slog.Logger
	userSaver    UserSaver
	userProvider UserProvider
	appProvider  AppProvider
	tokenTTL     time.Duration
}

type UserSaver interface {
	SaveUser(ctx context.Context, email string, passHash []byte) (uid int64, err error)
}

type UserProvider interface {
	User(ctx context.Context, email string) (models.User, error)
	IsAdmin(ctx context.Context, userID int64) (bool, error)
}

type AppProvider interface {
	App(ctx context.Context, appID int) (models.App, error)
}

// New returns a new instance of the Auth service
func New(
	log *slog.Logger,
	userSaver UserSaver,
	userProvider UserProvider,
	appProvider AppProvider,
	tokenTTL time.Duration,
) *Auth {
	return &Auth{
		log:          log,
		userSaver:    userSaver,
		userProvider: userProvider,
		appProvider:  appProvider,
		tokenTTL:     tokenTTL,
	}
}

// Login проверяет, если user с переданными параметрами существует в системе, возвращает токен доступа
//
// Если user существует, но логин или пароль невалидны, возвращается ошибка
// Если user не существует, возвращается ошибка
func (a Auth) Login(ctx context.Context, email string, password string, appID int) (string, error) {
	panic("not implemented")
}

func (a Auth) RegisterNewUser(ctx context.Context, email string, password string) (int64, error) {
	panic("not implemented")
}

func (a Auth) IsAdmin(ctx context.Context, userID int64) (bool, error) {
	panic("not implemented")
}
