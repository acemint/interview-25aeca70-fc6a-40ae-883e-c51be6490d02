package gobookcabin

import "gorm.io/gorm"

type GormAuthenticationService struct {
	db *gorm.DB
}

func NewGormAuthenticationService(db *gorm.DB) *GormAuthenticationService {
	return &GormAuthenticationService{
		db: db,
	}
}
