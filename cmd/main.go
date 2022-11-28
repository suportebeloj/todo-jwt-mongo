package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"todo-jwt-mongo/internal/api/http/rest"
)

func main() {
	httpServerSetup()
}

func SetupMongoDB() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	client, err := database.NewMongoDBClient(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func httpServerSetup(db *mongo.Client) {
	fiberApp := fiber.New()
	fiberApp.Use(logger.New())

	dbName := os.Getenv("DB_NAME")
	coll := db.Database(dbName).Collection("users")
	repo := authentication.NewUsersRepository(coll)

	authApp := app.NewAppAuthentication(repo)

	apiRest := rest.NewAuthenticationAdapter(fiberApp, authApp)
	rest.RegisterSwaggerDoc(fiberApp)
	apiRest.Run()
}
