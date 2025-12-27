package photo

type RecipeSettingsPrimitives struct {
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

type RecipeSettings struct {
	filmSimulation       string
	dynamicRange         string
	highlight            string
	shadow               string
	color                string
	noiseReduction       string
	sharpening           string
	clarity              string
	grainEffect          string
	colorChromeEffect    string
	colorChromeBlue      string
	whiteBalance         string
	iso                  string
	exposureCompensation string
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
		filmSimulation:       filmSimulation,
		dynamicRange:         dynamicRange,
		highlight:            highlight,
		shadow:               shadow,
		color:                color,
		noiseReduction:       noiseReduction,
		sharpening:           sharpening,
		clarity:              clarity,
		grainEffect:          grainEffect,
		colorChromeEffect:    colorChromeEffect,
		colorChromeBlue:      colorChromeBlue,
		whiteBalance:         whiteBalance,
		iso:                  iso,
		exposureCompensation: exposureCompensation,
	}
}

func RecipeSettingsFromPrimitives(rsp RecipeSettingsPrimitives) RecipeSettings {
	return RecipeSettings{
		filmSimulation:       rsp.FilmSimulation,
		dynamicRange:         rsp.DynamicRange,
		highlight:            rsp.Highlight,
		shadow:               rsp.Shadow,
		color:                rsp.Color,
		noiseReduction:       rsp.NoiseReduction,
		sharpening:           rsp.Sharpening,
		clarity:              rsp.Clarity,
		grainEffect:          rsp.GrainEffect,
		colorChromeEffect:    rsp.ColorChromeEffect,
		colorChromeBlue:      rsp.ColorChromeBlue,
		whiteBalance:         rsp.WhiteBalance,
		iso:                  rsp.Iso,
		exposureCompensation: rsp.ExposureCompensation,
	}
}

func (rs RecipeSettings) ToPrimitives() RecipeSettingsPrimitives {
	return RecipeSettingsPrimitives{
		FilmSimulation:       rs.filmSimulation,
		DynamicRange:         rs.dynamicRange,
		Highlight:            rs.highlight,
		Shadow:               rs.shadow,
		Color:                rs.color,
		NoiseReduction:       rs.noiseReduction,
		Sharpening:           rs.sharpening,
		Clarity:              rs.clarity,
		GrainEffect:          rs.grainEffect,
		ColorChromeEffect:    rs.colorChromeEffect,
		ColorChromeBlue:      rs.colorChromeBlue,
		WhiteBalance:         rs.whiteBalance,
		Iso:                  rs.iso,
		ExposureCompensation: rs.exposureCompensation,
	}
}
