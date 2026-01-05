package photoviewdrepository

import (
	"github.com/ariel-rubilar/photography-api/internal/projection/photoview/domain/photoview"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type photoRecipeSettings struct {
	FilmSimulation       string `bson:"filmSimulation"`
	DynamicRange         string `bson:"dynamicRange"`
	Highlight            string `bson:"highlight"`
	Shadow               string `bson:"shadow"`
	Color                string `bson:"color"`
	NoiseReduction       string `bson:"noiseReduction"`
	Sharpening           string `bson:"sharpening"`
	Clarity              string `bson:"clarity"`
	GrainEffect          string `bson:"grainEffect"`
	ColorChromeEffect    string `bson:"colorChromeEffect"`
	ColorChromeBlue      string `bson:"colorChromeBlue"`
	WhiteBalance         string `bson:"whiteBalance"`
	Iso                  string `bson:"iso"`
	ExposureCompensation string `bson:"exposureCompensation"`
}

type photoRecipe struct {
	ID       string              `bson:"id"`
	Name     string              `bson:"name"`
	Settings photoRecipeSettings `bson:"settings"`
	Link     string              `bson:"link"`
}

type photoDocument struct {
	ID     bson.ObjectID `bson:"_id,omitempty"`
	Title  string        `bson:"title"`
	URL    string        `bson:"url"`
	Recipe photoRecipe   `bson:"recipe"`
}

func DocumentFromDomain(p *photoview.PhotoView) (photoDocument, error) {
	id, err := bson.ObjectIDFromHex(p.ID)
	if err != nil {
		return photoDocument{}, err
	}
	return photoDocument{
		ID:    id,
		Title: p.Title,
		URL:   p.Url,
		Recipe: photoRecipe{
			ID:   p.Recipe.ID,
			Name: p.Recipe.Name,
			Settings: photoRecipeSettings{
				FilmSimulation:       p.Recipe.Settings.FilmSimulation,
				DynamicRange:         p.Recipe.Settings.DynamicRange,
				Highlight:            p.Recipe.Settings.Highlight,
				Shadow:               p.Recipe.Settings.Shadow,
				Color:                p.Recipe.Settings.Color,
				NoiseReduction:       p.Recipe.Settings.NoiseReduction,
				Sharpening:           p.Recipe.Settings.Sharpening,
				Clarity:              p.Recipe.Settings.Clarity,
				GrainEffect:          p.Recipe.Settings.GrainEffect,
				ColorChromeEffect:    p.Recipe.Settings.ColorChromeEffect,
				ColorChromeBlue:      p.Recipe.Settings.ColorChromeBlue,
				WhiteBalance:         p.Recipe.Settings.WhiteBalance,
				Iso:                  p.Recipe.Settings.Iso,
				ExposureCompensation: p.Recipe.Settings.ExposureCompensation,
			},
			Link: p.Recipe.Link,
		},
	}, nil
}
