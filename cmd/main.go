package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
	"time"
	"todo-jwt-mongo/internal/api/http/rest"
	"todo-jwt-mongo/internal/app"
	"todo-jwt-mongo/internal/core/authentication"
	"todo-jwt-mongo/internal/infra/database"
)

func main() {
	db, err := SetupMongoDB()
	if err != nil {
		log.Fatalln(err)
	}

	httpServerSetup(db)
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
