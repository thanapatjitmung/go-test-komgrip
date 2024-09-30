package entities

import "thanapatjitmung/go-test-komgrip/modules/models"

type Beer struct {
	ID       int64  `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Details  string `json:"details"`
	ImageURL string `json:"image_url"`
}

func (b *Beer) ToModel() *models.Beer {
	return &models.Beer{
		ID:       b.ID,
		Name:     b.Name,
		Type:     b.Type,
		Details:  b.Details,
		ImageURL: b.ImageURL,
	}
}
