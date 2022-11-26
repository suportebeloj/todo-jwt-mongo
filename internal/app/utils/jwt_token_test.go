package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"todo-jwt-mongo/internal/core/authentication/models"
)

func TestGenerateValidToken_WhenCallNewTokenFunction_AndNotHaveErrors(t *testing.T) {
	user := models.User{
		Username:   "test",
		Profile:    models.Profile{Email: "test@mail.com"},
		Permission: models.Permissions{Group: "user", Level: 6},
	}
	token, err := NewToken(&user)
	assert.NoError(t, err)
	err = VerifyToken(token)
	assert.NoError(t, err)
}
