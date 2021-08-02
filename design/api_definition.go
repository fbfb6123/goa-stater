package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/zaplogger"
)

var _ = API("goa_starter", func() {
	Title("goa_starterulator Service")
	Description("Service for adding numbers, a Goa teaser")
	Server("goa_starter", func() {
		Host("localhost", func() {
			URI("http://localhost:8085/api/v1")
			URI("grpc://localhost:8080")
		})
	})
})
