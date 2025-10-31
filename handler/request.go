package handler

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type CreateOpeningRequest struct {
	Role     string `json:"role" validate:"required"`
	Company  string `json:"company" validate:"required"`
	Location string `json:"location" validate:"required"`
	Remote   *bool  `json:"remote" validate:"required"`
	Link     string `json:"link" validate:"required"`
	Salary   int64  `json:"salary" validate:"required,gt=0"`
}

var validate = validator.New(validator.WithRequiredStructEnabled())

func (r *CreateOpeningRequest) Validate() error {
	if err := validate.Struct(r); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			return fmt.Errorf("param: %s is invalid (rule: %s)", e.Field(), e.Tag())
		}
	}
	return nil
}
