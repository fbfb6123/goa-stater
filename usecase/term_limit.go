package usecase

import (
	"context"
	"goa_starter/di"
	"goa_starter/domain/model"
	"goa_starter/domain/repository"
	log "goa_starter/gen/log"
	"goa_starter/lib/sf_error"
)

type (
	TermLimitUseCase struct {
		logger          *log.Logger
		termlimitRepository repository.TermLimitRepository
	}
)

func NewTermLimitUseCase(logger *log.Logger) *TermLimitUseCase {
	return &TermLimitUseCase{
		logger:          logger,
		termlimitRepository: di.InjectTermLimitRepository(),
	}
}

func (u *TermLimitUseCase) Get(ctx context.Context, id uint64) (*model.TermLimit, *sf_error.SfError) {
	var res model.TermLimit

	off, err := u.termlimitRepository.FindByID(id, nil)
	if err != nil {
		return nil, err
	}

	res = *model.ToModelTermLimit(off)

	return &res, nil
}
