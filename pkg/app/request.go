package app

import (
	"github.com/go-playground/validator/v10"

	"github.com/EDDYCJY/go-gin-example/pkg/logging"
)

// MarkErrors logs error logs
func MarkErrors(err error) {
	if err == nil {
		return
	}

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		logging.Error(err)
		return
	}

	for _, err := range validationErrors {
		logging.Info(err.Field(), err.Tag(), err.Param())
	}
}
