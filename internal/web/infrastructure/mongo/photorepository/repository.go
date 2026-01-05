package photorepository

import (
	"context"
	"fmt"

	"github.com/ariel-rubilar/photography-api/internal/web/photo"
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
		database:   "web",
		collection: "photos",
	}
}

func (r *repository) Search(ctx context.Context) ([]*photo.Photo, error) {

	cursor, err := r.client.Database(r.database).Collection(r.collection).Find(ctx, bson.D{})
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
		photos[i] = doc.toDomain()
	}

	return photos, nil
}
