package searchphotos

import "github.com/ariel-rubilar/photography-api/internal/web/photo"

type photoRecipeSettings struct {
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

type photoRecipe struct {
	Name     string              `json:"name"`
	Settings photoRecipeSettings `json:"settings"`
	Link     string              `json:"link"`
}

type photoDTO struct {
	ID     string      `json:"id"`
	Title  string      `json:"title"`
	URL    string      `json:"url"`
	Recipe photoRecipe `json:"recipe"`
}

type searchPhotosResponse struct {
	Data []photoDTO `json:"data"`
}

func newSearchPhotosResponse(photos []*photo.Photo) searchPhotosResponse {

	var photoDTOs []photoDTO

	for _, photo := range photos {

		primitives := photo.ToPrimitives()

		photoDTOs = append(photoDTOs, photoDTO{
			ID:    primitives.ID,
			Title: primitives.Title,
			URL:   primitives.URL,
			Recipe: photoRecipe{
				Name: primitives.Recipe.Name,
				Settings: photoRecipeSettings{
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

	return searchPhotosResponse{
		Data: photoDTOs,
	}
}
