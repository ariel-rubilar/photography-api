package recipereadrepository

import (
	"github.com/ariel-rubilar/photography-api/internal/projection/photoview/domain/reciperead"
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

func (r *repository) Get(id string) (*reciperead.RecipeRead, error) {
	// Implementation goes here
	return nil, nil
}
