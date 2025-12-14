package appCommon

import (
	"errors"
	"log"
)

var (
	RecordNotFound = errors.New("record not found")
)

func ErrCannotDeleteSessionID(err error) *AppError {
	return NewCustomError(err, "Can not delete session ID", "ErrCannotDeleteSessionID")
}

func AppRecover() {
	if err := recover(); err != nil {
		log.Println("Recovery error:", err)
	}
}
