package reciperepository

import (
	"context"
	"fmt"
	"strings"

	"github.com/ariel-rubilar/photography-api/internal/backoffice/recipe"
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

func (r *repository) Search(ctx context.Context, c recipe.Criteria) ([]*recipe.Recipe, error) {

	filter := bson.M{}

	for _, f := range c.Filters {
		switch f.Op {
		case recipe.OpEq:
			key := f.Field

			if key == recipe.FieldID {
				value, err := bson.ObjectIDFromHex(f.Value.(string))
				if err != nil {
					return nil, fmt.Errorf("invalid ObjectID format: %w", err)
				}
				filter["_id"] = value
				continue
			}

			filter[string(key)] = f.Value

		case recipe.OpContains:
			value := strings.TrimSpace(f.Value.(string))

			filter[string(f.Field)] = bson.M{
				"$regex":   value,
				"$options": "i",
			}
		}
	}

	cursor, err := r.client.Database(r.database).Collection(r.collection).Find(ctx, filter)
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
