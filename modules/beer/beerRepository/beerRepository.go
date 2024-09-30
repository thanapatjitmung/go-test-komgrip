package repository

import (
	"thanapatjitmung/go-test-komgrip/entities"
	_beerModel "thanapatjitmung/go-test-komgrip/modules/models"
)

type BeerRepository interface {
	GetAll(itemFilter *_beerModel.BeerFilter) ([]*entities.Beer, error)
	Counting(itemFilter *_beerModel.BeerFilter) (int64, error)
	FindById(itemId uint64) (*entities.Beer, error)
	Creating(itemEntity *entities.Beer) (*entities.Beer, error)
	Updating(itemId uint64, itemUpdateReq *_beerModel.BeerUpdateRequest) (uint64, error)
	Delete(itemId uint64) error
}
