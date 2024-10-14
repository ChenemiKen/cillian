package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validationMessages = map[string]string{
	"required": "%s is required",
	"max":      "%s should not be more than %s characters",
}

type fieldError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

func Validate(s interface{}) []fieldError {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(s)
	if err != nil {
		fieldErrors := []fieldError{}
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err)
			fieldName := err.Field()
			tag := err.Tag()
			param := err.Param()

			msg := validationMessages[tag]
			if param == "" {
				msg = fmt.Sprintf(msg, fieldName)
			} else {
				msg = fmt.Sprintf(msg, fieldName, param)
			}

			fieldErrors = append(fieldErrors, fieldError{
				Field: fieldName,
				Error: msg,
			})
		}
		return fieldErrors
	}
	return nil
}
