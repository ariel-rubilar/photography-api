package photorepository

import (
	"github.com/ariel-rubilar/photography-api/internal/backoffice/photo"
	"github.com/ariel-rubilar/photography-api/internal/backoffice/usecases/photoquery"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type photoDocument struct {
	ID       bson.ObjectID `bson:"_id,omitempty"`
	Title    string        `bson:"title"`
	URL      string        `bson:"url"`
	RecipeID string        `bson:"recipeId"`
}

func (p *photoDocument) ToDomain() *photoquery.PhotoDTO {

	primitives := photoquery.PhotoDTO{
		ID:       p.ID.Hex(),
		Title:    p.Title,
		URL:      p.URL,
		RecipeID: p.RecipeID,
	}

	return &primitives
}

func DocumentFromDomain(r *photo.Photo) (photoDocument, error) {
	primitives := r.ToPrimitives()
	id, err := bson.ObjectIDFromHex(primitives.ID)
	if err != nil {
		return photoDocument{}, err
	}
	return photoDocument{
		ID:       id,
		Title:    primitives.Title,
		URL:      primitives.URL,
		RecipeID: id.Hex(),
	}, nil
}
