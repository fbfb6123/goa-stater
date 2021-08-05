package model

import (
	"goa_starter/domain/entity"
	// "time"
)

type TermLimit struct {
	ID               uint64
	OfferID                          uint64
	Status                           uint64
	Amount             string
	PublishStartAt   string
}


func ToModelTermLimit(a *entity.TermLimit) *TermLimit {
	view := &TermLimit{}
	view.ID = a.ID
	view.Amount = a.Amount
	return view
}