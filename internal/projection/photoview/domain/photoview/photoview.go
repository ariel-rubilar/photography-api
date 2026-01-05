package photoview

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

type Recipe struct {
	id       string
	name     string
	settings RecipeSettings
	link     string
}

func BuildRecipe(id, name string, settings RecipeSettings, link string) Recipe {
	return Recipe{
		id:       id,
		name:     name,
		settings: settings,
		link:     link,
	}
}

type PhotoView struct {
	id     string
	title  string
	url    string
	recipe Recipe
}

func BuildPhotoView(id, title, url string, recipe Recipe) *PhotoView {
	return &PhotoView{
		id:     id,
		title:  title,
		url:    url,
		recipe: recipe,
	}
}
