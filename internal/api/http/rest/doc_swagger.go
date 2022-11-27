package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "todo-jwt-mongo/cmd/docs"
)

func RegisterSwaggerDoc(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)
}
