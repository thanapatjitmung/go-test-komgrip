package usecase

import (
	"encoding/json"
	"strconv"
	"thanapatjitmung/go-test-komgrip/entities"
	_logRepo "thanapatjitmung/go-test-komgrip/modules/log/logRepository"
	"time"
)

type logUsecaseImpl struct {
	logRepo _logRepo.LogRepository
}

func NewLogUsecaseImpl(logRepo _logRepo.LogRepository) LogUsecase {
	return &logUsecaseImpl{logRepo: logRepo}
}

func (u *logUsecaseImpl) LogAction(action string, itemId uint64, newData interface{}, oldData interface{}) error {
	newDataByte, _ := u.stringToByte(newData)
	oldDataByte, _ := u.stringToByte(oldData)

	log := entities.Log{
		Action:    action,
		Timestamp: time.Now(),
		ItemId:    strconv.FormatUint(itemId, 10),
		Data:      string(newDataByte),
		OldData:   string(oldDataByte),
	}

	err := u.logRepo.Create(log)
	if err != nil {
		return err
	}

	return nil
}

func (u *logUsecaseImpl) stringToByte(data interface{}) ([]byte, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return dataBytes, nil
}
