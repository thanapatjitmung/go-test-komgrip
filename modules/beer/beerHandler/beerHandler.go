package handler

import "github.com/labstack/echo/v4"

type BeerHandler interface {
	GetAll(pctx echo.Context) error
	Create(pctx echo.Context) error
	Update(pctx echo.Context) error
	Delete(pctx echo.Context) error

}
