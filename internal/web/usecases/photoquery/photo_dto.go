package photoquery

type PhotoDTORecipeSettings struct {
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

type PhotoDTORecipe struct {
	Name     string
	Settings PhotoDTORecipeSettings
	Link     string
	ID       string
}

type PhotoDTO struct {
	ID     string         `json:"id"`
	Title  string         `json:"title"`
	URL    string         `json:"url"`
	Recipe PhotoDTORecipe `json:"recipe"`
}
