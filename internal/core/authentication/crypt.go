package authentication

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"todo-jwt-mongo/internal/core/authentication/models"
)

func HashPassword(user models.RegisterUser) (hashedPass string, salt int, err error) {
	salt = rand.Int()
	plainStr := fmt.Sprintf("%s.%s.%d.%s", user.Username, user.Email, salt, user.Password)
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainStr), 14)
	if err != nil {
		return "", 0, err
	}

	return string(bytes), salt, err
}

func ValidatePassword(plainPassword string, user *models.User) bool {
	plainStr := fmt.Sprintf("%s.%s.%d.%s", user.Username, user.Profile.Email, user.Salt, plainPassword)
	err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(plainStr))
	return err == nil
}
