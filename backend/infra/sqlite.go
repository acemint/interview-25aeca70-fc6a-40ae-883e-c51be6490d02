package infra

import (
	"fmt"
	"gobookcabin/gobookcabin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitializeGorm initializes the DB with specific configuration
func InitializeGorm() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s", gobookcabin.AppConfigurationInstance.DBString)
	return gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}
