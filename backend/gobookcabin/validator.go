package gobookcabin

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

// newValidator returns a singleton of Validator if not yet initialized
func newValidator() *validator.Validate {
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	return validate
}

// shouldBindJsonAndValidate returns a gin response in JSON format, containing the error during requestBody validation
func shouldBindJsonAndValidate(c *gin.Context, requestBody any) error {
	if c == nil {
		return Errorf(ErrCodeInternal, "context should not be nil")
	}
	if requestBody == nil {
		return Errorf(ErrCodeInvalid, "request body should not be nil")
	}
	if err := c.ShouldBindJSON(requestBody); err != nil {
		return Errorf(ErrCodeInvalid, "unable to process JSON")
	}
	return shouldValidate(requestBody)
}

func shouldValidate(request any) error {
	err := ValidatorInstance.Struct(request)
	if err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			var fields []string
			for _, errField := range validationErrors {
				fields = append(fields, errField.Field())
			}
			return Errorf(ErrCodeInvalid, "validation failed on fields: %s", strings.Join(fields, ", "))
		}
	}
	return nil
}

type ValidationError struct {
	Code   string
	Fields map[string]string
}

func (e *ValidationError) Error() string {
	return "failed to validate request body"
}
