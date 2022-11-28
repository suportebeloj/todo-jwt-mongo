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

func (a Adapter) Register(ctx context.Context, user models.RegisterUser) (string, error) {
	var data models.UserData

	hashedPwd, salt, err := authentication.HashPassword(user)
	if err != nil {
		return "", err
	}
	data.Username = user.Username
	data.Profile = models.Profile{Email: user.Email}
	data.HashedPassword = hashedPwd
	data.Salt = salt

	savedUser, err := a.usersRepository.Save(ctx, data)
	if err != nil {
		return "", err
	}

	token, err := utils.NewToken(savedUser)
	if err != nil {
		return "", err
	}

	return token, nil
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
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func NewAppAuthentication(usersRepository ports.UserDatabasePort) *Adapter {
	return &Adapter{usersRepository: usersRepository}
}
