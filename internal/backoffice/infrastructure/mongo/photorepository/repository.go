package photorepository

import (
	"context"
	"fmt"
	"strings"

	"github.com/ariel-rubilar/photography-api/internal/backoffice/photo"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type repository struct {
	client     *mongo.Client
	database   string
	collection string
}

var _ photo.Repository = (*repository)(nil)

func NewMongoRepository(client *mongo.Client) *repository {
	return &repository{
		client:     client,
		database:   "backoffice",
		collection: "photos",
	}
}

func (r *repository) getCollection() *mongo.Collection {
	return r.client.Database(r.database).Collection(r.collection)
}

func (r *repository) Save(ctx context.Context, new *photo.Photo) error {
	collection := r.getCollection()

	d, error := DocumentFromDomain(new)
	if error != nil {
		return fmt.Errorf("failed to convert domain to document: %w", error)
	}
	_, err := collection.InsertOne(ctx, d)

	if err != nil {
		return fmt.Errorf("failed to insert document: %w", err)
	}

	return nil
}

func (r *repository) Search(ctx context.Context, c photo.Criteria) ([]*photo.Photo, error) {
	filter := bson.M{}

	for _, f := range c.Filters {
		switch f.Op {
		case photo.OpEq:
			key := f.Field

			if key == photo.FieldID {
				value, err := bson.ObjectIDFromHex(f.Value.(string))
				if err != nil {
					return nil, fmt.Errorf("invalid ObjectID format: %w", err)
				}
				filter["_id"] = value
				continue
			}

			filter[string(key)] = f.Value

		case photo.OpContains:
			value := strings.TrimSpace(f.Value.(string))

			filter[string(f.Field)] = bson.M{
				"$regex":   value,
				"$options": "i",
			}
		}
	}

	collection := r.getCollection()

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to execute find query: %w", err)
	}

	defer func() {
		_ = cursor.Close(ctx)
	}()

	documents := &[]photoDocument{}

	err = cursor.All(ctx, documents)
	if err != nil {
		return nil, fmt.Errorf("failed to decode documents: %w", err)
	}

	photos := make([]*photo.Photo, len(*documents))

	for i, doc := range *documents {
		new, err := doc.ToDomain()

		if err != nil {
			return []*photo.Photo{}, err
		}
		photos[i] = new
	}

	return photos, nil
}
