package ports

import (
	"context"
	"todo-jwt-mongo/internal/core/authentication/models"
)

type UserDatabasePort interface {
	Save(ctx context.Context, user models.UserData) (*primitive.ObjectID, error)
	GetByName(ctx context.Context, username string) (*models.User, error)
}
