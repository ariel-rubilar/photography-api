package searchphotos

import "github.com/ariel-rubilar/photography-api/internal/web/photo"

type PhotoRecipeSettings struct {
	FilmSimulation       string `json:"filmSimulation"`
	DynamicRange         string `json:"dynamicRange"`
	Highlight            string `json:"highlight"`
	Shadow               string `json:"shadow"`
	Color                string `json:"color"`
	NoiseReduction       string `json:"noiseReduction"`
	Sharpening           string `json:"sharpening"`
	Clarity              string `json:"clarity"`
	GrainEffect          string `json:"grainEffect"`
	ColorChromeEffect    string `json:"colorChromeEffect"`
	ColorChromeBlue      string `json:"colorChromeBlue"`
	WhiteBalance         string `json:"whiteBalance"`
	Iso                  string `json:"iso"`
	ExposureCompensation string `json:"exposureCompensation"`
}

type PhotoRecipe struct {
	Name     string              `json:"name"`
	Settings PhotoRecipeSettings `json:"settings"`
	Link     string              `json:"link"`
}

type PhotoDTO struct {
	ID     string      `json:"id"`
	Title  string      `json:"title"`
	URL    string      `json:"url"`
	Recipe PhotoRecipe `json:"recipe"`
}

func newSearchPhotosData(photos []*photo.Photo) []PhotoDTO {

	var photoDTOs []PhotoDTO

	for _, photo := range photos {

		primitives := photo.ToPrimitives()

		photoDTOs = append(photoDTOs, PhotoDTO{
			ID:    primitives.ID,
			Title: primitives.Title,
			URL:   primitives.URL,
			Recipe: PhotoRecipe{
				Name: primitives.Recipe.Name,
				Settings: PhotoRecipeSettings{
					FilmSimulation:       primitives.Recipe.Settings.FilmSimulation,
					DynamicRange:         primitives.Recipe.Settings.DynamicRange,
					Highlight:            primitives.Recipe.Settings.Highlight,
					Shadow:               primitives.Recipe.Settings.Shadow,
					Color:                primitives.Recipe.Settings.Color,
					NoiseReduction:       primitives.Recipe.Settings.NoiseReduction,
					Sharpening:           primitives.Recipe.Settings.Sharpening,
					Clarity:              primitives.Recipe.Settings.Clarity,
					GrainEffect:          primitives.Recipe.Settings.GrainEffect,
					ColorChromeEffect:    primitives.Recipe.Settings.ColorChromeEffect,
					ColorChromeBlue:      primitives.Recipe.Settings.ColorChromeBlue,
					WhiteBalance:         primitives.Recipe.Settings.WhiteBalance,
					Iso:                  primitives.Recipe.Settings.Iso,
					ExposureCompensation: primitives.Recipe.Settings.ExposureCompensation,
				},
				Link: primitives.Recipe.Link,
			},
		})
	}

	return photoDTOs
}
