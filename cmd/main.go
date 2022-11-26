package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"todo-jwt-mongo/internal/api/http/rest"
)

func main() {
	fiberApp := fiber.New()
	fiberApp.Use(logger.New())

	apiRest := rest.NewAuthenticationAdapter(fiberApp)
	apiRest.Run()
}
