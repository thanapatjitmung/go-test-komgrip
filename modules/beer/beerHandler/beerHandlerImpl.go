package handler

import (
	"net/http"
	"strconv"
	_beerUsecase "thanapatjitmung/go-test-komgrip/modules/beer/beerUsecase"
	"thanapatjitmung/go-test-komgrip/modules/custom"
	_logUsecase "thanapatjitmung/go-test-komgrip/modules/log/logUsecase"
	"thanapatjitmung/go-test-komgrip/modules/models"

	"github.com/labstack/echo/v4"
)

type beerHandlerImpl struct {
	beerUsecase _beerUsecase.BeerUsecase
	logUsecase  _logUsecase.LogUsecase
}

func NewBeerHandlerImpl(beerUsecase _beerUsecase.BeerUsecase, logUsecase _logUsecase.LogUsecase) BeerHandler {
	return &beerHandlerImpl{
		beerUsecase: beerUsecase,
		logUsecase:  logUsecase,
	}
}

func (h *beerHandlerImpl) GetAll(pctx echo.Context) error {
	itemFilter := new(models.BeerFilter)
	customEchoRequest := custom.NewCustomEchoRequest(pctx)
	if err := customEchoRequest.Bind(itemFilter); err != nil {
		return custom.ErrResponse(pctx, http.StatusBadRequest, err)
	}

	itemModelList, err := h.beerUsecase.GetAll(itemFilter)
	if err != nil {
		return custom.ErrResponse(pctx, http.StatusInternalServerError, err)
	}

	return custom.SuccessResponse(pctx, http.StatusOK, itemModelList)
}

func (h *beerHandlerImpl) Create(pctx echo.Context) error {
	item := new(models.BeerCreateRequest)
	customEchoRequest := custom.NewCustomEchoRequest(pctx)
	if err := customEchoRequest.Bind(item); err != nil {
		return custom.ErrResponse(pctx, http.StatusBadRequest, err)
	}
	itemModel, err := h.beerUsecase.Create(item)
	if err != nil {
		return custom.ErrResponse(pctx, http.StatusBadRequest, err)
	}

	_ = h.logUsecase.LogAction("Create", uint64(itemModel.ID), item, nil)

	return custom.SuccessResponse(pctx, http.StatusOK, itemModel)
}
func (h *beerHandlerImpl) Update(pctx echo.Context) error {
	item := new(models.BeerUpdateRequest)

	itemIdStr := pctx.Param("id")
	itemId, err := strconv.ParseUint(itemIdStr, 10, 64)
	if err != nil {
		return custom.ErrResponse(pctx, http.StatusBadRequest, err)
	}
	customEchoRequset := custom.NewCustomEchoRequest(pctx)
	if err := customEchoRequset.Bind(item); err != nil {
		return custom.ErrResponse(pctx, http.StatusBadRequest, err)
	}

	itemModel, err := h.beerUsecase.Update(itemId, item)
	if err != nil {
		return custom.ErrResponse(pctx, http.StatusInternalServerError, err)
	}

	_ = h.logUsecase.LogAction("Update", itemId, item, itemModel)

	return custom.SuccessResponse(pctx, http.StatusOK, itemModel)
}
func (h *beerHandlerImpl) Delete(pctx echo.Context) error {
	itemIdStr := pctx.Param("id")
	itemId, err := strconv.ParseUint(itemIdStr, 10, 64)
	if err != nil {
		return custom.ErrResponse(pctx, http.StatusBadRequest, err)
	}
	err = h.beerUsecase.Delete(itemId)
	if err != nil {
		return custom.ErrResponse(pctx, http.StatusInternalServerError, err)
	}

	_ = h.logUsecase.LogAction("Delete", itemId, nil, nil)

	return custom.SuccessResponse(pctx, http.StatusNoContent, "")
}
