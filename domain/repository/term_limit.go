package repository

import (
	"goa_starter/domain/entity"
	"goa_starter/lib/sf_error"

	"gorm.io/gorm"
)

type TermLimitRepository interface {
	// Find(*entity.TermLimit, *sf_error.SfError)
	FindByID(ID uint64, tx *gorm.DB) (*entity.TermLimit, *sf_error.SfError)
}
