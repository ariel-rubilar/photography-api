package reciperead

type RecipeRead struct {
	ID       string
	Name     string
	Settings RecipeReadSettings
	Link     string
}

func BuildRecipe(id string, name string, settings RecipeReadSettings, link string) *RecipeRead {
	return &RecipeRead{
		ID:       id,
		Name:     name,
		Settings: settings,
		Link:     link,
	}
}

type RecipeReadSettings struct {
	FilmSimulation       string
	DynamicRange         string
	Highlight            string
	Shadow               string
	Color                string
	NoiseReduction       string
	Sharpening           string
	Clarity              string
	GrainEffect          string
	ColorChromeEffect    string
	ColorChromeBlue      string
	WhiteBalance         string
	Iso                  string
	ExposureCompensation string
}

func BuildRecipeSettings(
	filmSimulation,
	dynamicRange,
	highlight,
	shadow,
	color,
	noiseReduction,
	sharpening,
	clarity,
	grainEffect,
	colorChromeEffect,
	colorChromeBlue,
	whiteBalance,
	iso,
	exposureCompensation string,
) RecipeReadSettings {
	return RecipeReadSettings{
		FilmSimulation:       filmSimulation,
		DynamicRange:         dynamicRange,
		Highlight:            highlight,
		Shadow:               shadow,
		Color:                color,
		NoiseReduction:       noiseReduction,
		Sharpening:           sharpening,
		Clarity:              clarity,
		GrainEffect:          grainEffect,
		ColorChromeEffect:    colorChromeEffect,
		ColorChromeBlue:      colorChromeBlue,
		WhiteBalance:         whiteBalance,
		Iso:                  iso,
		ExposureCompensation: exposureCompensation,
	}
}
