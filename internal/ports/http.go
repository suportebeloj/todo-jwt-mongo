package ports

import "github.com/gofiber/fiber/v2"

type HTTPAuthenticationAPIPort interface {
	Run()
	SignUp(c *fiber.Ctx) error
	SignIn(c *fiber.Ctx) error
}
