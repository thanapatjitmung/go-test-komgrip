package models

type (
	Beer struct {
		ID       int64  `json:"id"`
		Name     string `json:"name"`
		Type     string `json:"type"`
		Details  string `json:"details"`
		ImageURL string `json:"image_url"`
	}

	BeerCreateRequest struct {
		Name     string `json:"name" validate:"required,max=64"`
		Type     string `json:"type" validate:"required,max=64"`
		Details  string `json:"details" validate:"required,max=128"`
		ImageURL string `json:"image_url"`
	}

	BeerUpdateRequest struct {
		Name     string `json:"name" validate:"omitempty,max=64"`
		Type     string `json:"type" validate:"omitempty,max=64"`
		Details  string `json:"details" validate:"omitempty,max=128"`
		ImageURL string `json:"image_url"`
	}

	BeerFilter struct {
		Name string `query:"name" validate:"omitempty,max=64"`
		PaginateRequest
	}

	BeerResponse struct {
		Beers    []*Beer          `json:"beers"`
		Paginate PaginateResponse `json:"paginate"`
	}
)
