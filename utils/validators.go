package utils

import (
    "regexp"
    "github.com/go-playground/validator/v10"
)

func ValidatePassword(fl validator.FieldLevel) bool {
    password := fl.Field().String()
    re := regexp.MustCompile(`[A-Z]`)
    return re.MatchString(password) 
}
