package dto

import "github.com/go-playground/validator/v10"

var validate = validator.New()

type QR struct {
	Content string `json:"content" validate:"required"`
}

func (q *QR) Validate() error {
	return validate.Struct(q)
}
