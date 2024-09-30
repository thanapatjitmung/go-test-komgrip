package repository

import (
	"thanapatjitmung/go-test-komgrip/entities"
	"thanapatjitmung/go-test-komgrip/modules/beer/exception"
	_beerModel "thanapatjitmung/go-test-komgrip/modules/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type beerRepositoryImpl struct {
	mariaDb *gorm.DB
	logger  echo.Logger
}

func NewBeerRepositoryImpl(mariaDb *gorm.DB, logger echo.Logger) BeerRepository {
	return &beerRepositoryImpl{
		mariaDb: mariaDb,
		logger:  logger,
	}
}

func (r *beerRepositoryImpl) GetAll(itemFilter *_beerModel.BeerFilter) ([]*entities.Beer, error) {
	itemList := make([]*entities.Beer, 0)

	query := r.mariaDb.Model(&entities.Beer{})
	if itemFilter.Name != "" {
		query = query.Where("name LIKE ?", "%"+itemFilter.Name+"%")
	}

	offset := int((itemFilter.Page - 1) * itemFilter.Size)
	limit := int(itemFilter.Size)

	if err := query.Offset(offset).Limit(limit).Order("id asc").Find(&itemList).Error; err != nil {
		r.logger.Errorf("Get all item fail: %s", err.Error())
		return nil, &exception.BeerGetAll{}
	}
	return itemList, nil
}

func (r *beerRepositoryImpl) Counting(itemFilter *_beerModel.BeerFilter) (int64, error) {
	var count int64
	query := r.mariaDb.Model(&entities.Beer{}).Where("name LIKE ?", "%"+itemFilter.Name+"%")
	if err := query.Count(&count).Error; err != nil {
		r.logger.Errorf("Failed to counting item : %s", err.Error())
		return -1, &exception.BeerCounting{}
	}
	return count, nil
}

func (r *beerRepositoryImpl) FindById(itemId uint64) (*entities.Beer, error) {
	item := new(entities.Beer)
	if err := r.mariaDb.First(item, itemId).Error; err != nil {
		r.logger.Errorf("Failed to find beer by ID : %s", err.Error())
		return nil, &exception.BeerNotFound{ItemId: itemId}
	}
	return item, nil
}

func (r *beerRepositoryImpl) Creating(itemEntity *entities.Beer) (*entities.Beer, error) {
	item := new(entities.Beer)
	if err := r.mariaDb.Create(itemEntity).Scan(item).Error; err != nil {
		r.logger.Errorf("Creating item fail : %s", err.Error())
		return nil, &exception.BeerCreating{}
	}
	return item, nil
}

func (r *beerRepositoryImpl) Updating(itemId uint64, itemUpdateReq *_beerModel.BeerUpdateRequest) (uint64, error) {
	if err := r.mariaDb.Model(&entities.Beer{}).Where("id = ?", itemId).Updates(itemUpdateReq).Error; err != nil {
		r.logger.Errorf("Editing item fail : %s", err.Error())
		return 0, &exception.BeerUpdateing{ItemId: itemId}
	}
	return itemId, nil
}

func (r *beerRepositoryImpl) Delete(itemId uint64) error {
	if err := r.mariaDb.Delete(&entities.Beer{}, itemId).Error; err != nil {
		r.logger.Errorf("Editing item fail : %s", err.Error())
		return &exception.BeerDelete{ItemId: itemId}
	}
	return nil
}
