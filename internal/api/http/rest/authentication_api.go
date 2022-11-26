package rest

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"todo-jwt-mongo/internal/ports"
)

type AuthenticationAdapter struct {
	App             *fiber.App
	UsersRepository ports.UserDatabasePort
}

func NewAuthenticationAdapter(app *fiber.App) *AuthenticationAdapter {
	return &AuthenticationAdapter{App: app}
}

func (a *AuthenticationAdapter) Run() {
	authGroup := a.App.Group("/auth/")
	authGroup.Post("/login", a.SignIn)
	log.Fatalln(a.App.Listen(":3000"))
}

func (a *AuthenticationAdapter) SignUp(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (a *AuthenticationAdapter) SignIn(c *fiber.Ctx) error {
	return c.Status(200).Send([]byte("welcome"))
}
