package usecase

import (
	"thanapatjitmung/go-test-komgrip/entities"
	_beerRepo "thanapatjitmung/go-test-komgrip/modules/beer/beerRepository"
	"thanapatjitmung/go-test-komgrip/modules/models"
)

type beerUsecaseImpl struct {
	beerRepo _beerRepo.BeerRepository
}

func NewBeerUsecaseImpl(beerRepo _beerRepo.BeerRepository) BeerUsecase {
	return &beerUsecaseImpl{beerRepo: beerRepo}
}

func (u *beerUsecaseImpl) GetAll(itemFilter *models.BeerFilter) (*models.BeerResponse, error) {
	item, err := u.beerRepo.GetAll(itemFilter)
	if err != nil {
		return nil, err
	}
	itemCount, err := u.beerRepo.Counting(itemFilter)
	if err != nil {
		return nil, err
	}

	size := itemFilter.Size
	page := itemFilter.Page
	totalPage := u.totalPageCalculation(itemCount, size)

	result := u.toResultResponse(item, page, totalPage)
	return result, nil
}

func (u *beerUsecaseImpl) Create(item *models.BeerCreateRequest) (*models.Beer, error) {
	itemEntity := &entities.Beer{
		Name:     item.Name,
		Type:     item.Type,
		Details:  item.Details,
		ImageURL: item.ImageURL,
	}
	result, err := u.beerRepo.Creating(itemEntity)
	if err != nil {
		return nil, err
	}
	return result.ToModel(), nil
}

func (u *beerUsecaseImpl) Update(itemId uint64, itemUpdate *models.BeerUpdateRequest) (*models.Beer, error) {
	id, err := u.beerRepo.Updating(itemId, itemUpdate)
	if err != nil {
		return nil, err
	}
	result, err := u.beerRepo.FindById(id)
	if err != nil {
		return nil, err
	}
	return result.ToModel(), nil
}

func (u *beerUsecaseImpl) Delete(itemId uint64) error {
	err := u.beerRepo.Delete(itemId)
	if err != nil {
		return err
	}
	return nil
}

func (u *beerUsecaseImpl) totalPageCalculation(totalItems int64, size int64) int64 {
	totalPage := totalItems / size

	if totalItems%size != 0 {
		totalPage++
	}

	return totalPage
}

func (u *beerUsecaseImpl) toResultResponse(itemEntity []*entities.Beer, page, totalPage int64) *models.BeerResponse {
	itemModel := make([]*models.Beer, 0)
	for _, item := range itemEntity {
		itemModel = append(itemModel, item.ToModel())
	}
	return &models.BeerResponse{
		Beers: itemModel,
		Paginate: models.PaginateResponse{
			Page:      page,
			TotalPage: totalPage,
		},
	}
}
