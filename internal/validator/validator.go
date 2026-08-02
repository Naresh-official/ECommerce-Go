package validator

import (
	"errors"
	"strings"

	v "github.com/go-playground/validator/v10"
)

var Validate = v.New()

type ValidationErrors = v.ValidationErrors

func ValidateRequest(req any) error {
	if err := Validate.Struct(req); err != nil {
		validationErrors, ok := err.(ValidationErrors)
		if !ok {
			return err
		}

		messages := make([]string, 0, len(validationErrors))
		for _, fieldError := range validationErrors {
			messages = append(messages, ValidationMessage(fieldError))
		}

		return errors.New(strings.Join(messages, "\n"))
	}

	return nil
}

func ValidationMessage(fe v.FieldError) string {
	switch fe.Tag() {

	case "required":
		return fe.Field() + " is required"

	case "email":
		return "Invalid email address"

	case "min":
		return fe.Field() + " must be at least " + fe.Param() + " characters"

	case "max":
		return fe.Field() + " cannot exceed " + fe.Param() + " characters"

	case "len":
		return fe.Field() + " must be exactly " + fe.Param() + " characters"

	default:
		return "Invalid " + fe.Field()
	}
}
