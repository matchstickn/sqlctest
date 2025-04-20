package server

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Tag   string
	Field string
	Value interface{}
	Error string
}

func NewValidator() *validator.Validate {
	return validator.New()
}

func Validate(v *validator.Validate, data interface{}) ([]ErrorResponse, validator.ValidationErrors) {
	validationErrors := []ErrorResponse{}

	errs := v.Struct(data)
	validateErrs := errs.(validator.ValidationErrors)
	if errs != nil {
		for _, err := range validateErrs {
			var elem ErrorResponse

			elem.Tag = err.Tag()
			elem.Field = err.Field()
			elem.Value = err.Value()
			elem.Error = err.Error()

			validationErrors = append(validationErrors, elem)
		}
	}
	return validationErrors, validateErrs
}

func HandleValidationErrors(errs validator.ValidationErrors, errResps []ErrorResponse) error {
	errMsgs := make([]string, 0)

	for i := 0; i < len(errResps); i++ {
		respErr := errResps[i].Error
		if respErr != "" {
			for _, err := range errs {
				field := err.Field()
				value := err.Value().(string)
				tag := err.Tag()
				errMsgs = append(errMsgs, "["+field+"]: "+value+" | Needs to implement "+tag)
			}
			return &fiber.Error{
				Code:    fiber.ErrBadRequest.Code,
				Message: strings.Join(errMsgs, " and "),
			}
		}
	}
	return nil
}

func RegisterValidations(v *validator.Validate) error {
	v.RegisterValidation("example", func(fl validator.FieldLevel) bool {
		return fl.Field().Bool()
	})
	return nil
}
