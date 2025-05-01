package lib

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

// Validator instance
var validate = validator.New()

// ValidateStruct checks if the struct fields are valid
func ValidateStruct(data interface{}) map[string]string {
	err := validate.Struct(data)
	if err != nil {
		errors := make(map[string]string)
		for _, e := range err.(validator.ValidationErrors) {
			key := strings.ToLower(e.Field())
			errors[key] = e.Tag()
		}
		return errors
	}
	return nil
}
