package server

import (
	_beerHandler "thanapatjitmung/go-test-komgrip/modules/beer/beerHandler"
	_beerRepository "thanapatjitmung/go-test-komgrip/modules/beer/beerRepository"
	_beerUsecase "thanapatjitmung/go-test-komgrip/modules/beer/beerUsecase"
	_logRepository "thanapatjitmung/go-test-komgrip/modules/log/logRepository"
	_logUsecase "thanapatjitmung/go-test-komgrip/modules/log/logUsecase"
)

func (s *echoServer) initBeerRouter() {
	router := s.app.Group("/beer")

	beerRepo := _beerRepository.NewBeerRepositoryImpl(s.mariaDb, s.app.Logger)
	logRepo := _logRepository.NewLogRepositoryImpl(s.mongoDb)

	beerUsecase := _beerUsecase.NewBeerUsecaseImpl(beerRepo)
	logUsecase := _logUsecase.NewLogUsecaseImpl(logRepo)

	beerHandler := _beerHandler.NewBeerHandlerImpl(beerUsecase, logUsecase)

	router.GET("", beerHandler.GetAll)
	router.POST("", beerHandler.Create)
	router.PUT("/:id", beerHandler.Update)
	router.DELETE("/:id", beerHandler.Delete)
}
