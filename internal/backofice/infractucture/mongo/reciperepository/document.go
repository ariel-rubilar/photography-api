package reciperepository

import (
	"github.com/ariel-rubilar/photography-api/internal/backofice/recipe"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type settings struct {
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

type recipeDocument struct {
	ID       bson.ObjectID `bson:"_id,omitempty"`
	Name     string        `bson:"name"`
	Settings settings      `bson:"settings"`
	Link     string        `bson:"link"`
}

func DocumentFromDomain(r *recipe.Recipe) (recipeDocument, error) {
	primitives := r.ToPrimitives()
	id, err := bson.ObjectIDFromHex(primitives.ID)
	if err != nil {
		return recipeDocument{}, err
	}
	return recipeDocument{
		ID:   id,
		Name: primitives.Name,
		Settings: settings{
			FilmSimulation:       primitives.Settings.FilmSimulation,
			DynamicRange:         primitives.Settings.DynamicRange,
			Highlight:            primitives.Settings.Highlight,
			Shadow:               primitives.Settings.Shadow,
			Color:                primitives.Settings.Color,
			NoiseReduction:       primitives.Settings.NoiseReduction,
			Sharpening:           primitives.Settings.Sharpening,
			Clarity:              primitives.Settings.Clarity,
			GrainEffect:          primitives.Settings.GrainEffect,
			ColorChromeEffect:    primitives.Settings.ColorChromeEffect,
			ColorChromeBlue:      primitives.Settings.ColorChromeBlue,
			WhiteBalance:         primitives.Settings.WhiteBalance,
			Iso:                  primitives.Settings.Iso,
			ExposureCompensation: primitives.Settings.ExposureCompensation,
		},
		Link: primitives.Link,
	}, nil
}

func (p recipeDocument) toDomain() *recipe.Recipe {

	settings := recipe.NewRecipeSettings(
		p.Settings.FilmSimulation,
		p.Settings.DynamicRange,
		p.Settings.Highlight,
		p.Settings.Shadow,
		p.Settings.Color,
		p.Settings.NoiseReduction,
		p.Settings.Sharpening,
		p.Settings.Clarity,
		p.Settings.GrainEffect,
		p.Settings.ColorChromeEffect,
		p.Settings.ColorChromeBlue,
		p.Settings.WhiteBalance,
		p.Settings.Iso,
		p.Settings.ExposureCompensation,
	)

	recipe := recipe.NewRecipe(
		p.ID.Hex(),
		p.Name,
		settings,
		p.Link,
	)

	return recipe
}
