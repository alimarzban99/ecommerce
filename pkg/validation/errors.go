package validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

// ValidationError represents a field-level validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidationErrors represents a collection of validation errors
type ValidationErrors struct {
	Errors []ValidationError `json:"errors"`
}

// FormatValidationErrors converts Gin/validator errors to user-friendly messages
func FormatValidationErrors(err error) *ValidationErrors {
	var validationErrors ValidationErrors

	// Check if it's a validator.ValidationErrors type
	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrs {
			validationErrors.Errors = append(validationErrors.Errors, ValidationError{
				Field:   getFieldName(fieldError),
				Message: getErrorMessage(fieldError),
			})
		}
	} else {
		// If it's not a validation error, return a generic error
		validationErrors.Errors = append(validationErrors.Errors, ValidationError{
			Field:   "general",
			Message: err.Error(),
		})
	}

	return &validationErrors
}

// getFieldName extracts the field name from the validation error
// Converts struct field name (PascalCase) to JSON field name (snake_case)
func getFieldName(fieldError validator.FieldError) string {
	field := fieldError.Field()

	// Convert PascalCase to snake_case
	var result strings.Builder
	for i, r := range field {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result.WriteRune('_')
		}
		result.WriteRune(r)
	}

	fieldName := strings.ToLower(result.String())

	// Handle special cases (ID, URL, UUID, etc.)
	fieldName = strings.ReplaceAll(fieldName, "_i_d", "_id")
	fieldName = strings.ReplaceAll(fieldName, "_u_r_l", "_url")
	fieldName = strings.ReplaceAll(fieldName, "_u_u_i_d", "_uuid")

	return fieldName
}

// getErrorMessage converts validation tag to user-friendly message
func getErrorMessage(fieldError validator.FieldError) string {
	tag := fieldError.Tag()
	param := fieldError.Param()

	switch tag {
	case "required":
		return "The field is required"
	case "email":
		return "Must have a valid email address"
	case "min":
		return fmt.Sprintf("Should be more than %s", param)
	case "max":
		return fmt.Sprintf("Should be less than %s", param)
	case "len":
		return fmt.Sprintf("Must be exactly %s characters", param)
	case "gte":
		return fmt.Sprintf("Must be greater than or equal to %s", param)
	case "lte":
		return fmt.Sprintf("Must be less than or equal to %s", param)
	case "gt":
		return fmt.Sprintf("Must be greater than %s", param)
	case "lt":
		return fmt.Sprintf("Must be less than %s", param)
	case "oneof":
		return fmt.Sprintf("Must be one of: %s", strings.ReplaceAll(param, " ", ", "))
	case "numeric":
		return "Must be a valid number"
	case "alpha":
		return "Must contain only letters"
	case "alphanum":
		return "Must contain only letters and numbers"
	case "url":
		return "Must be a valid URL"
	case "uuid":
		return "Must be a valid UUID"
	case "datetime":
		return "Must be a valid date and time"
	case "date":
		return "Must be a valid date"
	default:
		// Return a generic message for unknown tags
		if param != "" {
			return fmt.Sprintf("Invalid value (constraint: %s, expected: %s)", tag, param)
		}
		return fmt.Sprintf("Invalid value (constraint: %s)", tag)
	}
}
