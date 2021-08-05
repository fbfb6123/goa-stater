package goastarterapi

import (
	"context"
	log "goa_starter/gen/log"
	termlimit "goa_starter/gen/term_limit"
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

// Add implements add.
func (s *termLimitsrvc) Add(ctx context.Context, p *termlimit.AddPayload) (res int, err error) {
	s.logger.Info("termLimit.add")
	return p.C - p.D, nil
}
