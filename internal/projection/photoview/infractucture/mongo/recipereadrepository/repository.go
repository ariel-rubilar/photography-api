package recipereadrepository

import (
	"context"

	"github.com/ariel-rubilar/photography-api/internal/projection/photoview/domain/reciperead"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type repository struct {
	client     *mongo.Client
	database   string
	collection string
}

var _ reciperead.Repository = (*repository)(nil)

func NewMongoRepository(client *mongo.Client) *repository {
	return &repository{
		client:     client,
		database:   "backoffice",
		collection: "recipes",
	}
}

func (r *repository) getCollection() *mongo.Collection {
	return r.client.Database(r.database).Collection(r.collection)
}

func (r *repository) Get(ctx context.Context, id string) (*reciperead.RecipeRead, error) {
	colllection := r.getCollection()
	filter := bson.M{
		"_id": id,
	}

	var document *recipeDocument

	err := colllection.FindOne(ctx, filter).Decode(document)
	if err != nil {
		return nil, err
	}

	return document.ToDomain(), nil
}
