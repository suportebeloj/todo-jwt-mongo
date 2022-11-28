package rest

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"time"
	"todo-jwt-mongo/internal/core/authentication/models"
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
	var registerData models.RegisterUser
	if err := c.BodyParser(&registerData); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("body parser error, invalid data, check request body and try again.")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	token, err := a.AuthApp.Register(ctx, registerData)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"token": token,
	})
}

func (a *AuthenticationAdapter) SignIn(c *fiber.Ctx) error {
	var cred models.AuthUser
	if err := c.BodyParser(&cred); err != nil {
		if cred.Username == "" || cred.Password == "" {
			return c.Status(fiber.StatusBadRequest).Send([]byte("username and password are required"))
		}
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	token, err := a.AuthApp.Authenticate(ctx, cred)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("username or password invalid.")
	}

	return c.Status(200).JSON(fiber.Map{
		"token": token,
	})
}
