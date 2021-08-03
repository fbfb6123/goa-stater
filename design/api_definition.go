package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
	_ "goa.design/plugins/v3/zaplogger"
)

var _ = API("goa_starter", func() {
	Title("goa_starterulator Service")
	Description("Service for adding numbers, a Goa teaser")
	cors.Origin("*", func() {
		cors.Headers("X-Authorization", "X-Time", "X-Api-Version",
			"Content-Type", "Origin", "Authorization")
		cors.Methods("GET")
		cors.Expose("Content-Type", "Origin")
		cors.MaxAge(100)
		cors.Credentials()
	})
	Server("goa_starter", func() {
		Host("localhost", func() {
			URI("http://localhost:8085/api/v1")
			URI("grpc://localhost:8080")
		})
	})
})
