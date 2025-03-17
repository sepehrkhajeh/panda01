package DomainApp

import "github.com/go-playground/validator/v10"

func ValidateData(req DomainCreateRequest) error {
	validate := validator.New()
	return validate.Struct(req)
}
