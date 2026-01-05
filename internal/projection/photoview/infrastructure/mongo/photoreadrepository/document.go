package photoreadrepository

import (
	"github.com/ariel-rubilar/photography-api/internal/projection/photoview/domain/photoread"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type photoDocument struct {
	ID       bson.ObjectID `bson:"_id,omitempty"`
	Title    string        `bson:"title"`
	URL      string        `bson:"url"`
	RecipeID string        `bson:"recipeId"`
}

func (p photoDocument) ToDomain() *photoread.PhotoRead {
	return photoread.Build(
		p.ID.Hex(),
		p.Title,
		p.URL,
		p.RecipeID,
	)
}
