package usecase

type LogUsecase interface {
	LogAction(action string, itemId uint64, newData interface{}, oldData interface{}) error
}
