package validation

import (
	"reflect"
	"strconv"
	"strings"
)

func ValidateStruct(data interface{}) []string {
	var validationErrors []string

	value := reflect.ValueOf(data)
	valueType := value.Type()

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fieldName := valueType.Field(i).Name
		fieldTag := valueType.Field(i).Tag.Get("validate")

		if fieldTag != "" {
			tags := strings.Split(fieldTag, ",")
			for _, tag := range tags {
				validationError := validateTag(tag, fieldName, field)
				if validationError != "" {
					validationErrors = append(validationErrors, validationError)
				}
			}
		}
	}

	return validationErrors
}

func validateTag(tag string, fieldName string, field reflect.Value) string {
	parts := strings.Split(tag, "=")
	switch parts[0] {
	case "required":
		if field.Interface() == reflect.Zero(field.Type()).Interface() {
			return fieldName + " is required"
		}
	case "min":
		minValue, _ := strconv.Atoi(parts[1])
		if field.Int() < int64(minValue) {
			return fieldName + " should be at least " + parts[1]
		}
	case "max":
		maxValue, _ := strconv.Atoi(parts[1])
		if field.Int() > int64(maxValue) {
			return fieldName + " should be at most " + parts[1]
		}
	}
	return ""
}
