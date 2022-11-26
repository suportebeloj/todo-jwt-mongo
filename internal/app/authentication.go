package app

import (
	"context"
	"todo-jwt-mongo/internal/app/utils"
	"todo-jwt-mongo/internal/core/authentication"
	"todo-jwt-mongo/internal/core/authentication/models"
	"todo-jwt-mongo/internal/ports"
)

type Adapter struct {
	usersRepository ports.UserDatabasePort
}

func (a Adapter) Authenticate(ctx context.Context, user models.AuthUser) (string, error) {
	userModel, err := a.usersRepository.GetByName(ctx, user.Username)
	if err != nil {
		return "", err
	}

	if ok := authentication.ValidatePassword(user.Password, userModel); !ok {
		return "", err
	}
	tokenString, err := utils.NewToken(userModel)
	return tokenString, nil
}

func NewAppAuthentication(usersRepository ports.UserDatabasePort) *Adapter {
	return &Adapter{usersRepository: usersRepository}
}
