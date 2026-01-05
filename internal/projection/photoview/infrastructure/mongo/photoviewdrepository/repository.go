package photoviewdrepository

import (
	"context"

	"github.com/ariel-rubilar/photography-api/internal/projection/photoview/domain/photoview"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type repository struct {
	client     *mongo.Client
	database   string
	collection string
}

var _ photoview.Repository = (*repository)(nil)

func NewMongoRepository(client *mongo.Client) *repository {
	return &repository{
		client:     client,
		database:   "web",
		collection: "photos",
	}
}

func (r *repository) getCollection() *mongo.Collection {
	return r.client.Database(r.database).Collection(r.collection)
}

func (r *repository) Save(ctx context.Context, photoview *photoview.PhotoView) error {

	doc, err := DocumentFromDomain(photoview)
	if err != nil {
		return err
	}

	_, err = r.getCollection().InsertOne(ctx, doc)

	if err != nil {
		return err
	}

	return nil
}
