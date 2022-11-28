package ports

import (
	"context"
	"todo-jwt-mongo/internal/core/authentication/models"
)

type AuthenticationPort interface {
	Authenticate(ctx context.Context, user models.AuthUser) (string, error)
	Register(ctx context.Context, user models.RegisterUser) (string, error)
}
