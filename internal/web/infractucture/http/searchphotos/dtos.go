package searchphotos

import "github.com/ariel-rubilar/photography-api~/internal/web/domain"

type PhotoDTO struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
}

type SearchPhotosResponse struct {
	Data []PhotoDTO `json:"data"`
}

func NewSearchPhotosResponse(photos []*domain.Photo) *SearchPhotosResponse {

	var photoDTOs []PhotoDTO

	for _, photo := range photos {
		photoDTOs = append(photoDTOs, PhotoDTO{
			ID:    photo.ID,
			Title: photo.Title,
			URL:   photo.URL,
		})
	}

	return &SearchPhotosResponse{
		Data: photoDTOs,
	}
}
