package design

import (
	. "goa.design/goa/v3/dsl"
)

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
})

var _ = Service("goa_starter-calc", func() {
	Description("The goa_starter-calc service performs operations on numbers.")

	Method("add", func() {
		Payload(func() {
			Field(1, "c", Int, "Left operand")
			Field(2, "d", Int, "Right operand")
			Required("c", "d")
		})

		Result(Int)

		HTTP(func() {
			GET("/goa_starter-calc/{c}/{d}")
		})

		GRPC(func() {
		})
	})
})

var _ = Service("term_limit", func() {
	Description("OW表示可能一覧取得")

	Method("get", func() {
		Payload(func() {
			Field(1, "id", Int, "term_limit のID")
			Field(2, "muid", String, "MUID")
			Field(3, "media_id", UInt64, "メディアID")
			Required("id")
			Required("muid")
			Required("media_id")
		})

		Result(TermLimitResponse)

		HTTP(func() {
			GET("/term_limit")
			Param("muid", String, "endUserのMUID")
			Param("media_id", UInt64, "メディアID")
		})

		GRPC(func() {
		})
	})
})