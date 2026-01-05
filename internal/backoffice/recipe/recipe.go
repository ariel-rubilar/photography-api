package recipe

import (
	"fmt"
	"strings"
)

type RecipeName struct {
	value string
}

func NewRecipeName(value string) (RecipeName, error) {
	value = strings.TrimSpace(value)
	value = strings.ToLower(value)

	if value == "" {
		return RecipeName{}, fmt.Errorf("recipe name cannot be empty")
	}
	return RecipeName{value: value}, nil
}

func (rn RecipeName) Value() string {
	return rn.value
}

type RecipePrimitives struct {
	Name     string
	Settings RecipeSettingsPrimitives
	Link     string
	ID       string
}

type Recipe struct {
	id       string
	name     RecipeName
	settings RecipeSettings
	link     string
}

func NewRecipe(id string, name string, settings RecipeSettings, link string) (*Recipe, error) {
	nameObj, err := NewRecipeName(name)
	if err != nil {
		return nil, err
	}
	return &Recipe{
		id:       id,
		name:     nameObj,
		settings: settings,
		link:     link,
	}, nil
}

func CreateRecipe(id string, name string, settings RecipeSettings, link string) (*Recipe, error) {
	return NewRecipe(id, name, settings, link)
}

func RecipeFromPrimitives(rp RecipePrimitives) (*Recipe, error) {
	settings := RecipeSettingsFromPrimitives(rp.Settings)

	return NewRecipe(rp.ID, rp.Name, settings, rp.Link)
}

func (r Recipe) ID() string {
	return r.id
}

func (r Recipe) Name() string {
	return r.name.Value()
}

func (r Recipe) ToPrimitives() RecipePrimitives {
	return RecipePrimitives{
		ID:       r.id,
		Name:     r.name.value,
		Settings: r.settings.ToPrimitives(),
		Link:     r.link,
	}
}
