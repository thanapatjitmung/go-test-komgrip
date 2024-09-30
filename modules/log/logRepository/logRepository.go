package repository

import "thanapatjitmung/go-test-komgrip/entities"

type LogRepository interface {
	Create(logEntities entities.Log) error
}
