package utils

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

// A custom validator for checking if a field is in valid email format
func NameValidator(fl validator.FieldLevel) bool {
	name := fl.Field().String()
	if len(name) < 1 || len(name) > 20 {
		return false
	}

	// Allow letters, digits, underscores, and Chinese characters
	for _, char := range name {
		if !(unicode.IsLetter(char) || unicode.IsDigit(char) || char == '_' || (char >= '\u4e00' && char <= '\u9fff')) {
			return false
		}
	}
	return true
}
