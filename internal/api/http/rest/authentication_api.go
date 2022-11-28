package rest

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"todo-jwt-mongo/internal/ports"
)

type AuthenticationAdapter struct {
	App     *fiber.App
	AuthApp ports.AuthenticationPort
}

func NewAuthenticationAdapter(app *fiber.App, authApp ports.AuthenticationPort) *AuthenticationAdapter {
	return &AuthenticationAdapter{App: app, AuthApp: authApp}
}

func (a *AuthenticationAdapter) Run() {
	authGroup := a.App.Group("/auth/")
	authGroup.Post("/login", a.SignIn)
	authGroup.Post("/register", a.SignUp)

	port := os.Getenv("PORT")
	log.Fatalln(a.App.Listen(fmt.Sprintf(":%s", port)))
}

func (a *AuthenticationAdapter) SignUp(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (a *AuthenticationAdapter) SignIn(c *fiber.Ctx) error {
	return c.Status(200).Send([]byte("welcome"))
}
