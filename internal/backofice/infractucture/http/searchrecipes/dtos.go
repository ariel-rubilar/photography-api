package searchrecipes

import (
	"github.com/ariel-rubilar/photography-api/internal/backofice/recipe"
)

type settingDTO struct {
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

type recipeDTO struct {
	Name     string     `json:"name"`
	Settings settingDTO `json:"settings"`
	Link     string     `json:"link"`
}

type searchPhotosResponse struct {
	Data []recipeDTO `json:"data"`
}

func newSearchRecipesResponse(recipes []*recipe.Recipe) searchPhotosResponse {

	recipeDTOs := make([]recipeDTO, 0, len(recipes))

	for _, recipe := range recipes {

		primitives := recipe.ToPrimitives()

		recipeDTOs = append(recipeDTOs, recipeDTO{
			Name: primitives.Name,
			Settings: settingDTO{
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
		})
	}

	return searchPhotosResponse{
		Data: recipeDTOs,
	}
}
