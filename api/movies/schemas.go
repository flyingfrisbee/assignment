package movies

import (
	"assessment/api/common"
	"assessment/db/model"
)

type movieRequest struct {
	Title       string  `json:"title" validate:"required,gte=1"`
	Description string  `json:"description" validate:"required,gte=1"`
	Rating      float32 `json:"rating" validate:"required,gt=0"`
	Image       string  `json:"image" validate:"required,gte=1"`
}

func (mr *movieRequest) parseTo() *model.Movie {
	return &model.Movie{
		Title:       mr.Title,
		Description: mr.Description,
		Rating:      mr.Rating,
		Image:       mr.Image,
	}
}

type movie struct {
	ID          int               `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Rating      float32           `json:"rating"`
	Image       string            `json:"image"`
	CreatedAt   common.CustomTime `json:"created_at"`
	UpdatedAt   common.CustomTime `json:"updated_at"`
}

func (m *movie) parseFrom(movieDB *model.Movie) {
	m.ID = int(movieDB.ID)
	m.Title = movieDB.Title
	m.Description = movieDB.Description
	m.Rating = movieDB.Rating
	m.Image = movieDB.Image
	m.CreatedAt = common.CustomTime(movieDB.CreatedAt)
	m.UpdatedAt = common.CustomTime(movieDB.UpdatedAt)
}
