package reciperepository

import (
	"context"
	"fmt"

	"github.com/ariel-rubilar/photography-api/internal/backofice/recipe"
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
		database:   "backoffice",
		collection: "recipes",
	}
}

func (r *repository) Search(ctx context.Context) ([]*recipe.Recipe, error) {

	cursor, err := r.client.Database(r.database).Collection(r.collection).Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("failed to execute find query: %w", err)
	}

	defer func() {
		_ = cursor.Close(ctx)
	}()

	documents := &[]recipeDocument{}

	err = cursor.All(ctx, documents)
	if err != nil {
		return nil, fmt.Errorf("failed to decode documents: %w", err)
	}

	recipes := make([]*recipe.Recipe, len(*documents))

	for i, doc := range *documents {
		recipes[i] = doc.toDomain()
	}

	return recipes, nil
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
