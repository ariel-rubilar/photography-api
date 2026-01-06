package photoview

type RecipeSettings struct {
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

func NewRecipeSettings(
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
) RecipeSettings {
	return RecipeSettings{
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
