package photorepository

import (
	"context"
	"fmt"

	"github.com/ariel-rubilar/photography-api/internal/backofice/photo"
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
