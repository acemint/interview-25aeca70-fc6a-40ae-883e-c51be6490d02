package gobookcabin

import (
	"github.com/go-playground/validator/v10"
)

var (
	AppConfigurationInstance *AppConfiguration
	ValidatorInstance        *validator.Validate
)

// InitializeGlobals provide utils to be used throughout the application
func InitializeGlobals() {
	AppConfigurationInstance = &AppConfiguration{}
	err := loadConfiguration(AppConfigurationInstance)
	if err != nil {
		panic(err.Error())
	}

	ValidatorInstance = newValidator()
}
