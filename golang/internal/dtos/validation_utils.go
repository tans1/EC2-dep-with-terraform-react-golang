package dtos

import (
	"gopkg.in/go-playground/validator.v9"
)

type DtoError struct {
	Param   string
	Message string
}

func msgForTag(vErr validator.FieldError) string {
	switch vErr.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "password":
		return "Invalid password"
	case "login":
		return "Invalid credential"
	default:
		return "Invalid data" // default
	}

}
func TranslateValidationError(err error) []*DtoError {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		out := make([]*DtoError, len(validationErrors))
		for i, vErr := range validationErrors {
			out[i] = &DtoError{vErr.Field(), msgForTag(vErr)}
		}
		return out
	}
	return make([]*DtoError, 0)
}
