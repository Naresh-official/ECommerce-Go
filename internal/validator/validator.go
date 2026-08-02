package validator

import v "github.com/go-playground/validator/v10"

var Validate = v.New()

type ValidationErrors = v.ValidationErrors

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
