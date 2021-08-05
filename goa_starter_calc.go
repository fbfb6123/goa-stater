package goastarterapi

import (
	"context"
	goastartercalc "goa_starter/gen/goa_starter_calc"
	log "goa_starter/gen/log"
)

// goa_starter-calc service example implementation.
// The example methods log the requests and return zero values.
type goaStarterCalcsrvc struct {
	logger *log.Logger
}

// NewGoaStarterCalc returns the goa_starter-calc service implementation.
func NewGoaStarterCalc(logger *log.Logger) goastartercalc.Service {
	return &goaStarterCalcsrvc{logger}
}

// Add implements add.
func (s *goaStarterCalcsrvc) Add(ctx context.Context, p *goastartercalc.AddPayload) (res int, err error) {
	s.logger.Info("goaStarterCalc.add")
	return p.C - p.D, nil
}
