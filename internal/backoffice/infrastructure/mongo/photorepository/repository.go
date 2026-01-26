package photorepository

import (
	"context"
	"fmt"
	"strings"

	"github.com/ariel-rubilar/photography-api/internal/backoffice/photo"
	"github.com/ariel-rubilar/photography-api/internal/backoffice/usecases/photoquery"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type repository struct {
	client     *mongo.Client
	database   string
	collection string
}

type repo interface {
	photo.Repository
	photoquery.Repository
}

var _ repo = (*repository)(nil)

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

func (r *repository) Search(ctx context.Context, c photoquery.Criteria) ([]*photoquery.PhotoDTO, error) {
	filter := bson.M{}

	for _, f := range c.Filters {
		switch f.Op {
		case photoquery.OpEq:
			key := f.Field

			if key == photoquery.FieldID {
				value, err := bson.ObjectIDFromHex(f.Value.(string))
				if err != nil {
					return nil, fmt.Errorf("invalid ObjectID format: %w", err)
				}
				filter["_id"] = value
				continue
			}

			filter[string(key)] = f.Value

		case photoquery.OpContains:
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

	photos := make([]*photoquery.PhotoDTO, len(*documents))

	for i, doc := range *documents {
		new := doc.ToDomain()

		photos[i] = new
	}

	return photos, nil
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
