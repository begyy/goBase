package utils

import (
	"github.com/go-playground/validator/v10"
	"reflect"
)

var validate = validator.New()

type ValidationError struct {
	Field string `json:"field_name"`
	Error string `json:"error"`
}

type LogicError struct {
	Message string `json:"message"`
}

func FormatValidationErrors(err error, dto interface{}) []ValidationError {
	var errors []ValidationError

	val := reflect.ValueOf(dto).Elem()
	typ := val.Type()

	for _, err := range err.(validator.ValidationErrors) {
		field, _ := typ.FieldByName(err.Field())
		jsonTag := field.Tag.Get("json")

		ve := ValidationError{
			Field: jsonTag,
			Error: err.Tag(),
		}
		errors = append(errors, ve)
	}

	return errors
}

func ValidateAndFormat(dto interface{}) []ValidationError {
	err := validate.Struct(dto)
	if err != nil {
		return FormatValidationErrors(err, dto)
	}
	return nil
}
