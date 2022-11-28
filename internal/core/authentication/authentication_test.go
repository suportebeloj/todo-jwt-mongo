package authentication

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"testing"
	"time"
	"todo-jwt-mongo/internal/core/authentication/models"
	"todo-jwt-mongo/internal/infra/database"
	"todo-jwt-mongo/internal/ports"
)

type AuthenticationTestSuit struct {
	suite.Suite
	Repository ports.UserDatabasePort
	Collection *mongo.Collection
}

func (s *AuthenticationTestSuit) SetupTest() {
	databaseTestClient, err := database.NewMongoDBClient(context.TODO())
	if err != nil {
		log.Fatalln(err)
	}

	coll := databaseTestClient.Database("test").Collection("users")
	s.Collection = coll

	repo := NewUsersRepository(coll)
	s.Repository = repo
}

func (s *AuthenticationTestSuit) TearDownTest() {
	_, _ = s.Collection.DeleteMany(context.TODO(), bson.D{})
}

func (s *AuthenticationTestSuit) TestRegisterNewUser_WithValidData_AndNotReceiveErrors() {
	id := primitive.NewObjectID()
	registerData := models.RegisterUser{
		Username: fmt.Sprintf("%s-test", id.Hex()),
		Password: "testPassword",
		Email:    fmt.Sprintf("%s@mail.com", id.Hex()),
	}
	pwdHash, salt, err := HashPassword(registerData)

	profile := models.Profile{
		Email: registerData.Email,
	}
	userData := models.UserData{
		Username:       registerData.Username,
		HashedPassword: pwdHash,
		Salt:           salt,
		Profile:        profile,
		CreatedAt:      time.Now(),
	}

	object, err := s.Repository.Save(context.TODO(), userData)
	s.NoError(err)
	s.True(primitive.IsValidObjectID(object.ID.Hex()))
}

func (s *AuthenticationTestSuit) TestGivenAValidUser_WhenICallGetByUsername_AndNotHasError() {
	id := primitive.NewObjectID()
	registerData := models.RegisterUser{
		Username: fmt.Sprintf("%s-test", id.Hex()),
		Password: "testPassword",
		Email:    fmt.Sprintf("%s@mail.com", id.Hex()),
	}

	pwdHash, salt, err := HashPassword(registerData)
	if err != nil {
		panic(err)
	}

	profile := models.Profile{
		Email: registerData.Email,
	}
	userData := models.UserData{
		Username:       registerData.Username,
		HashedPassword: pwdHash,
		Salt:           salt,
		Profile:        profile,
		CreatedAt:      time.Now(),
	}

	_, err = s.Repository.Save(context.TODO(), userData)
	s.NoError(err)

	found, err := s.Repository.GetByName(context.TODO(), userData.Username)
	s.NoError(err)
	s.Equal(userData.Username, found.Username)
	s.Equal(userData.HashedPassword, found.HashedPassword)
	s.Equal(userData.Salt, found.Salt)
	s.Equal(userData.Profile, found.Profile)
	s.Equal(userData.CreatedAt.Unix(), found.CreatedAt.Unix())
}

func TestRunSuit(t *testing.T) {
	suite.Run(t, new(AuthenticationTestSuit))
}
