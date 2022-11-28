package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"todo-jwt-mongo/internal/api/http/rest"
)

func main() {
	httpServerSetup()
}

func httpServerSetup() {
	fiberApp := fiber.New()
	fiberApp.Use(logger.New())

	apiRest := rest.NewAuthenticationAdapter(fiberApp)
	rest.RegisterSwaggerDoc(fiberApp)
	apiRest.Run()
}
