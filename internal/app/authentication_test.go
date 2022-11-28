package app

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/suite"
	"log"
	"math/rand"
	"testing"
	"todo-jwt-mongo/internal/core/authentication"
	"todo-jwt-mongo/internal/core/authentication/models"
	"todo-jwt-mongo/internal/infra/database"
	"todo-jwt-mongo/internal/ports"
)

type AuthenticationTestSuite struct {
	suite.Suite
	AuthApp         ports.AuthenticationPort
	usersRepository ports.UserDatabasePort
}

func (s *AuthenticationTestSuite) SetupTest() {
	dbTestClient, err := database.NewMongoDBClient(context.TODO())
	if err != nil {
		log.Fatalln(err)
	}

	coll := dbTestClient.Database("test").Collection("users")
	repo := authentication.NewUsersRepository(coll)
	s.usersRepository = repo

	app := NewAppAuthentication(repo)
	s.AuthApp = app
}

func (s *AuthenticationTestSuite) TestAuthenticationAnExistentUser_WhenICallFunctionAuthenticate_AndReturnAValidJWTToken() {
	id := rand.Int()
	profile := models.Profile{
		Email: fmt.Sprintf("%d@mail.com", id),
	}

	user := models.UserData{
		Username: fmt.Sprintf("%d-test", id),
		Profile:  profile,
	}

	plainPassword := "test-password"
	hashedPass, salt, err := authentication.HashPassword(models.RegisterUser{
		Username: user.Username,
		Password: plainPassword,
		Email:    user.Profile.Email,
	})
	s.NoError(err)

	user.HashedPassword = hashedPass
	user.Salt = salt

	_, err = s.usersRepository.Save(context.TODO(), user)
	s.NoError(err)

	res, err := s.AuthApp.Authenticate(context.TODO(), models.AuthUser{Username: user.Username, Password: plainPassword})
	s.NoError(err)
	s.NotEqual("", res)
}

func (s *AuthenticationTestSuite) TestGiverToken_AfterRegisterUser() {
	registerData := models.RegisterUser{
		Username: "testuser",
		Password: "testpassword",
		Email:    "test@mail.com",
	}

	token, err := s.AuthApp.Register(context.TODO(), registerData)
	s.NoError(err)
	s.NotEqual("", token)
}

func TestRunTestSuite(t *testing.T) {
	suite.Run(t, new(AuthenticationTestSuite))
}
