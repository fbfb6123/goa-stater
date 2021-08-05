package infra

import (
	"goa_starter/domain/entity"
	"goa_starter/domain/repository"
	"goa_starter/lib/sf_error"

	"gorm.io/gorm"
)

type TermLimitRepository struct {
	db *gorm.DB
}

func NewTermLimitRepository(db *gorm.DB) repository.TermLimitRepository {
	return &TermLimitRepository{db: db}
}

func (_ *TermLimitRepository) TableName() string {
	return "term_limit"
}

func (r *TermLimitRepository) FindByID(ID uint64, tx *gorm.DB) (*entity.TermLimit, *sf_error.SfError) {
	var db *gorm.DB
	var e entity.TermLimit

	if tx != nil {
		db = tx.Debug().Table(r.TableName()).Where("id = ?", ID).
			First(&e)
	} else {
		db = r.db.Debug().Table(r.TableName()).Where("id = ?", ID).
			First(&e)
	}

	if db.Error == gorm.ErrRecordNotFound {
		return nil, sf_error.NewSfError(sf_error.NotFoundError, db.Error.Error())
	} else if db.Error != nil {
		return nil, sf_error.NewSfError(sf_error.InternalError, db.Error.Error())
	}

	return &e, nil
}