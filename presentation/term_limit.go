package presentation

import (
	"context"
	log "goa_starter/gen/log"
	termlimit "goa_starter/gen/term_limit"
	"goa_starter/usecase"
)

// term_limit service example implementation.
// The example methods log the requests and return zero values.
type termLimitsrvc struct {
	logger *log.Logger
}

// NewTermLimit returns the term_limit service implementation.
func NewTermLimit(logger *log.Logger) termlimit.Service {
	return &termLimitsrvc{logger}
}

// Get implements get.
// func (s *termLimitsrvc) Get(ctx context.Context, p *termlimit.GetPayload) (res int, err error) {
// 	s.logger.Info("termLimit.get")
// 	return
// }

func (s *termLimitsrvc) Get(ctx context.Context, p *termlimit.GetPayload) (*termlimit.TermLimitResponse, error) {
	res := &termlimit.TermLimitResponse{}
	s.logger.Info("termlimit.get")

	id := uint64(p.ID)

	s.logger.Info("ID : " + string(rune(id)))

	u := usecase.NewTermLimitUseCase(s.logger)

	off, err := u.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	res.CreateFromTermLimit(off)

	return res, nil
}