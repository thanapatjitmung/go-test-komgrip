package models

type (

	PaginateRequest struct {
		Page int64 `query:"page" validate:"required,min=1"`
		Size int64 `query:"size" validate:"required,min=1,max=20"`
	}

	PaginateResponse struct {
		Page      int64 `json:"page"`
		TotalPage int64 `json:"totalPage"`
	}
)
