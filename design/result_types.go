package design

import (
	"goa_starter/domain/model"

	. "goa.design/goa/v3/dsl"
)

var TermLimitResponse = ResultType("application/vnd.term_limit", func() {
	Description("表示可能上限")
	CreateFrom(model.TermLimit{})
	TypeName("TermLimitResponse")
	Attributes(func() {
		Field(1, "id", UInt64)
		Field(2, "offer_id", UInt64)
		// Field(3, "status", UInt64)
	})
	View("default", func() {
		Attribute("id")
		Attribute("offer_id")
		// Attribute("status")
	})
	Required("id", "offer_id")
})