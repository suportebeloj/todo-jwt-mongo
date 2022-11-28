package authentication

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"todo-jwt-mongo/internal/core/authentication/models"
)

type Adapter struct {
	collection *mongo.Collection
}

func NewUsersRepository(collection *mongo.Collection) *Adapter {
	return &Adapter{collection: collection}
}

func (a Adapter) Save(ctx context.Context, user models.UserData) (*models.User, error) {
	res, err := a.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	if id, ok := res.InsertedID.(primitive.ObjectID); ok {
		var user models.User
		if err := a.collection.FindOne(ctx, bson.D{{"_id", id}}).Decode(&user); err != nil {
			return nil, err
		}
		return &user, nil
	} else {
		return nil, errors.New("invalid objectId")
	}
}

func (a Adapter) GetByName(ctx context.Context, username string) (*models.User, error) {
	filter := bson.D{{"username", username}}

	var result models.User
	err := a.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
