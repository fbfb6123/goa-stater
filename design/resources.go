package design

import (
	. "goa.design/goa/v3/dsl"
)

// /actionsの定義をする
var _ = Resource("actions", func() {
	// actionsリソースのベースパス
	BasePath("/actions")
	/*
	   リソースに対してどういった操作を行うか定義する
	   add リソースを追加する
	   list    リソースをリストで取得する
	   delete  リソースを削除する
	   上記のような感じで定義すればおｋです。
	*/
	Action("ping", func() {
		// アクションの説明
		Description("サーバーとの導通確認")
		Routing(
			// エンドポイント -> GET http://localhost/api/v1/actions/pingになる
			GET("/ping"),
		)
		// 返したいレスポンス
		// 200 OK + MessageTypeで定義しているMediaType
		Response(OK, MessageType)
		// 400 BadRequest + ErrorMediaというデフォルトで容易されているMediaType
		// 足りないパラメーターなどがあれば自動的に返される
		Response(BadRequest, ErrorMedia)
	})
	Action("hello", func() {
		Description("挨拶する")
		Routing(
			GET("/hello"),
		)
		// リクエストで付加出来るパラメーター
		Params(func() {
			// nameという名前でパラメーターをStringでなげれる
			Param("name", String, "名前", func() {
				// もし、空だった場合は空文字を格納する
				Default("")
			})
			// 必ず設定されるべきパラメーター（デフォルト値があるので、存在しなければ空になる）
			Required("name")
		})
		Response(OK, MessageType)
		Response(BadRequest, ErrorMedia)
	})
	Action("ID", func() {
		Description("複数アクション（:id）")
		Routing(
			// エンドポイントにリソースを指定出来る
			// GET http://localhost:8080/api/v1/actions/1になる
			GET("/:id"),
		)
		Params(func() {
			// :IDはIntegert型でなければならない。
			Param("id", Integer, "id")
			// Requiredはリソースを含めたエンドポイントになるので、定義しなくても良い
			//Required("id")
		})
		Response(OK, IntegerType)
		// 指定したリソースが無ければNotFoundを返す可能生がある
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})

// Swaggerをローカルで実行するめの定義
var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET")
	})
	Files("/swagger.json", "swagger/swagger.json")
	Files("/swagger/*filepath", "public/swagger/")
})
