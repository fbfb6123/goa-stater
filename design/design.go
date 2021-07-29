package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("goa_starter", func() {
	Title("goa_starter Service")
	Description("Service for adding numbers, a Goa teaser")
	Server("goa_starter", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
			URI("grpc://localhost:8080")
		})
	})
})

var _ = Service("goa_starter", func() {
	Description("The goa_starter service performs operations on numbers.")

	Method("add", func() {
		Payload(func() {
			Field(1, "a", Int, "Left operand")
			Field(2, "b", Int, "Right operand")
			Required("a", "b")
		})

		Result(Int)

		HTTP(func() {
			GET("/add/{a}/{b}")
		})

		GRPC(func() {
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})
