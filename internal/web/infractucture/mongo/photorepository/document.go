package photorepository

import (
	"github.com/ariel-rubilar/photography-api~/internal/web/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type photoDocument struct {
	ID    bson.ObjectID `bson:"_id,omitempty"`
	Title string        `bson:"title"`
	URL   string        `bson:"url"`
}

func (p photoDocument) toDomain() *domain.Photo {
	return &domain.Photo{
		ID:    p.ID.Hex(),
		Title: p.Title,
		URL:   p.URL,
	}
}
