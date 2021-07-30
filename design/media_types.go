package design

import (
	. "goa.design/goa/v3/dsl"
)

// レスポンスデータの定義
// MediaTypeに名前をつけます
var IntegerMedia = ResultType("application/vnd.integer+json", func() {
	// 説明
	Description("example")
	// どのような値があるか（複数定義出来る）
	Attributes(func() {
		// idはInteger型
		Attribute("id", UInt64, "id", func() {
			// 返すレスポンスの例
			Example(1)
		})
		// レスポンスに必須な要素（基本は全て必須にした方が楽）
		Required("id")
	})
	// 返すレスポンスのフォーマット（別記事で紹介予定）
	View("default", func() {
		Attribute("id")
	})
})
