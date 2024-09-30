package usecase

import "thanapatjitmung/go-test-komgrip/modules/models"

type BeerUsecase interface {
	GetAll(itemFilter *models.BeerFilter) (*models.BeerResponse, error)
	Create(item *models.BeerCreateRequest) (*models.Beer, error)
	Update(itemId uint64, itemUpdate *models.BeerUpdateRequest) (*models.Beer, error)
	Delete(itemId uint64) error

}
