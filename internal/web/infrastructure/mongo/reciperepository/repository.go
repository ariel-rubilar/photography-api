package reciperepository

import (
	"context"
	"fmt"

	"github.com/ariel-rubilar/photography-api/internal/web/recipe"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type repository struct {
	client     *mongo.Client
	database   string
	collection string
}

var _ recipe.Repository = (*repository)(nil)

func NewMongoRepository(client *mongo.Client) *repository {
	return &repository{
		client:     client,
		database:   "web",
		collection: "recipes",
	}
}

func (r *repository) getCollection() *mongo.Collection {
	return r.client.Database(r.database).Collection(r.collection)
}

func (r *repository) Save(ctx context.Context, rec *recipe.Recipe) error {

	document, err := DocumentFromDomain(rec)

	if err != nil {
		return fmt.Errorf("failed to convert recipe to document: %w", err)
	}

	_, err = r.client.Database(r.database).Collection(r.collection).InsertOne(ctx, document)

	if err != nil {
		return fmt.Errorf("failed to insert document: %w", err)
	}

	return nil
}

func (r *repository) Exists(ctx context.Context, id string) (bool, error) {
	collection := r.getCollection()

	filter := bson.M{}

	value, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return false, fmt.Errorf("invalid ObjectID format: %w", err)
	}

	filter["_id"] = value

	result := collection.FindOne(ctx, filter)

	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
