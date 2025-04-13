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

func Validate(v *validator.Validate, data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	errs := v.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem ErrorResponse

			elem.Tag = err.Tag()     // Export struct field name
			elem.Field = err.Field() // Export struct tag
			elem.Value = err.Value() // Export field value
			elem.Error = err.Error()

			validationErrors = append(validationErrors, elem)
		}
	}
	return validationErrors
}

func HandleValidationErrors(errs validator.ValidationErrors, errResp []ErrorResponse) error {
	errMsgs := make([]string, 0)

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

func RegisterValidations(v *validator.Validate) error {
	v.RegisterValidation("example", func(fl validator.FieldLevel) bool {
		return fl.Field().Bool()
	})
	return nil
}
