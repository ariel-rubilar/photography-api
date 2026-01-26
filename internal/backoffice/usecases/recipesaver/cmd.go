package recipesaver

type SaveRecipeSettingsCommand struct {
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

type SaveRecipeCommand struct {
	ID       string
	Name     string
	Settings SaveRecipeSettingsCommand
	Link     string
}
