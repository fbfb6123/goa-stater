package goastarterapi

import (
	"context"
	goastarter "goa_starter/gen/goa_starter"
	log "goa_starter/gen/log"
)

// goa_starter service example implementation.
// The example methods log the requests and return zero values.
type goaStartersrvc struct {
	logger *log.Logger
}

// NewGoaStarter returns the goa_starter service implementation.
func NewGoaStarter(logger *log.Logger) goastarter.Service {
	return &goaStartersrvc{logger}
}

// Add implements add.
func (s *goaStartersrvc) Add(ctx context.Context, p *goastarter.AddPayload) (res int, err error) {
	s.logger.Info("goaStarter.add")
	return p.A + p.B, nil
}

// func (s *goaStartersrvc) Get(ctx context.Context, p *goastarter.AddPayload) (res int, err error) {
// 	s.logger.Info("goaStarter.add")
// 	return p.A + p.B, nil
// }

// func (s *offerAttributessrvc) Get(ctx context.Context, p *offerattributes.GetPayload) (*offerattributes.OfferResponse, error) {
// 	res := &offerattributes.OfferResponse{}
// 	s.logger.Info("offerAttributes.get")

// 	id := uint64(p.ID)

// 	s.logger.Info("offer ID : " + string(rune(id)))

// 	u := usecase.NewOfferUseCase(s.logger)

// 	off, err := u.Get(ctx, id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	res.CreateFromOffer(off)

// 	return res, nil
// }