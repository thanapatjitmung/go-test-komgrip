package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	_beerUsecase "thanapatjitmung/go-test-komgrip/modules/beer/beerUsecase"
	"thanapatjitmung/go-test-komgrip/modules/custom"
	_logUsecase "thanapatjitmung/go-test-komgrip/modules/log/logUsecase"
	"thanapatjitmung/go-test-komgrip/modules/models"
	"time"

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

	// แยกการอ่านข้อมูล JSON และไฟล์
	if err := pctx.Request().ParseMultipartForm(32 << 20); err != nil { // ขนาดสูงสุด 32MB
		return custom.ErrResponse(pctx, http.StatusBadRequest, err)
	}

	if data := pctx.FormValue("data"); data != "" {
		if err := json.Unmarshal([]byte(data), item); err != nil {
			return custom.ErrResponse(pctx, http.StatusBadRequest, err)
		}
	} else {
		return custom.ErrResponse(pctx, http.StatusBadRequest, fmt.Errorf("missing data"))
	}

	// กำหนด path สำหรับโฟลเดอร์ uploads
	uploadsDir := "./uploads"

	// ตรวจสอบว่าโฟลเดอร์ uploads มีอยู่หรือไม่ หากไม่ให้สร้าง
	if _, err := os.Stat(uploadsDir); os.IsNotExist(err) {
		err := os.Mkdir(uploadsDir, 0755)
		if err != nil {
			return custom.ErrResponse(pctx, http.StatusInternalServerError, err)
		}
	}

	// อ่านไฟล์จากฟอร์ม
	file, err := pctx.FormFile("image_url")
	if err != nil {
		return custom.ErrResponse(pctx, http.StatusBadRequest, err)
	}

	src, err := file.Open()
	if err != nil {
		return custom.ErrResponse(pctx, http.StatusInternalServerError, err)
	}
	defer src.Close()

	fileByte, err := ioutil.ReadAll(src)
	if err != nil {
		return custom.ErrResponse(pctx, http.StatusInternalServerError, err)
	}

	fileType := http.DetectContentType(fileByte)
	var fileName string

	if fileType == "application/pdf" {
		fileName = strconv.FormatInt(time.Now().Unix(), 10) + ".pdf"
	} else {
		fileName = strconv.FormatInt(time.Now().Unix(), 10) + ".jpg"
	}

	// เขียนไฟล์ไปที่ uploads directory
	fullFilePath := filepath.Join(uploadsDir, fileName)
	err = ioutil.WriteFile(fullFilePath, fileByte, 0777)
	if err != nil {
		return custom.ErrResponse(pctx, http.StatusInternalServerError, err)
	}

	// เก็บแค่ชื่อไฟล์ใน item.ImageURL
	item.ImageURL = fileName // เปลี่ยนเป็นแค่ชื่อไฟล์

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

	if data := pctx.FormValue("data"); data != "" {
		if err := json.Unmarshal([]byte(data), item); err != nil {
			return custom.ErrResponse(pctx, http.StatusBadRequest, err)
		}
	} else {
		return custom.ErrResponse(pctx, http.StatusBadRequest, fmt.Errorf("missing data"))
	}

	uploadsDir := "./uploads"

	if _, err := os.Stat(uploadsDir); os.IsNotExist(err) {
		err := os.Mkdir(uploadsDir, 0755)
		if err != nil {
			return custom.ErrResponse(pctx, http.StatusInternalServerError, err)
		}
	}

	file, err := pctx.FormFile("image_url")
	if err == nil {
		src, err := file.Open()
		if err != nil {
			return custom.ErrResponse(pctx, http.StatusInternalServerError, err)
		}
		defer src.Close()

		fileByte, err := ioutil.ReadAll(src)
		if err != nil {
			return custom.ErrResponse(pctx, http.StatusInternalServerError, err)
		}

		fileType := http.DetectContentType(fileByte)
		var fileName string

		if fileType == "application/pdf" {
			fileName = strconv.FormatInt(time.Now().Unix(), 10) + ".pdf"
		} else {
			fileName = strconv.FormatInt(time.Now().Unix(), 10) + ".jpg"
		}

		fullFilePath := filepath.Join(uploadsDir, fileName)
		err = ioutil.WriteFile(fullFilePath, fileByte, 0777)
		if err != nil {
			return custom.ErrResponse(pctx, http.StatusInternalServerError, err)
		}
		item.ImageURL = fileName
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
