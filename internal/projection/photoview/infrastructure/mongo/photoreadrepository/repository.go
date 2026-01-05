package photoreadrepository

import (
	"context"
	"fmt"

	"github.com/ariel-rubilar/photography-api/internal/projection/photoview/domain/photoread"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type repository struct {
	client     *mongo.Client
	database   string
	collection string
}

var _ photoread.Repository = (*repository)(nil)

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

func (r *repository) Get(ctx context.Context, id string) (*photoread.PhotoRead, error) {
	colllection := r.getCollection()

	idObj, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"_id": idObj,
	}

	var document photoDocument

	err = colllection.FindOne(ctx, filter).Decode(&document)
	if err != nil {
		return nil, fmt.Errorf("error finding photo document: %w", err)
	}

	return document.ToDomain(), nil
}
