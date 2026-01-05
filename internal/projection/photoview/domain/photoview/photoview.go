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

type Recipe struct {
	ID       string
	Name     string
	Settings RecipeSettings
	Link     string
}

func BuildRecipe(id, name string, settings RecipeSettings, link string) Recipe {
	return Recipe{
		ID:       id,
		Name:     name,
		Settings: settings,
		Link:     link,
	}
}

type PhotoView struct {
	ID     string
	Title  string
	Url    string
	Recipe Recipe
}

func BuildPhotoView(id, title, url string, recipe Recipe) *PhotoView {
	return &PhotoView{
		ID:     id,
		Title:  title,
		Url:    url,
		Recipe: recipe,
	}
}
