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
