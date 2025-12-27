package photorepository

import (
	"github.com/ariel-rubilar/photography-api/internal/web/photo"
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

func (p photoDocument) toDomain() *photo.Photo {

	settings := photo.NewRecipeSettings(
		p.Recipe.Settings.FilmSimulation,
		p.Recipe.Settings.DynamicRange,
		p.Recipe.Settings.Highlight,
		p.Recipe.Settings.Shadow,
		p.Recipe.Settings.Color,
		p.Recipe.Settings.NoiseReduction,
		p.Recipe.Settings.Sharpening,
		p.Recipe.Settings.Clarity,
		p.Recipe.Settings.GrainEffect,
		p.Recipe.Settings.ColorChromeEffect,
		p.Recipe.Settings.ColorChromeBlue,
		p.Recipe.Settings.WhiteBalance,
		p.Recipe.Settings.Iso,
		p.Recipe.Settings.ExposureCompensation,
	)

	recipe := photo.NewRecipe(
		p.Recipe.Name,
		settings,
		p.Recipe.Link,
	)

	return photo.NewPhoto(
		p.ID.Hex(),
		p.Title,
		p.URL,
		recipe,
	)
}
