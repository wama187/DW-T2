package design

import . "goa.design/goa/v3/dsl"

var BookV2 = Type("BookV2", func() {
	Description("Representa un libro")
	Attribute("_id", String, "ID único del libro")
	Attribute("title", String, "Título del libro")
	Attribute("author", String, "Autor del libro")
	Attribute("owner_id", UInt, "ID del usuario que posee el libro")
	Required("_id", "title", "author", "owner_id")
})


var BookPayloadV2 = Type("BookPayloadV2", func() {
	Description("Datos necesarios para un libro")
	Token("jwt")
	Attribute("_id", String, "ID único del libro")
	Attribute("title", String, "Título del libro")
	Attribute("author", String, "Autor del libro")
	Required("title", "author")
})

var UpdateBookPayloadV2 = Type("UpdateBookPayloadV2", func() {
	Description("Datos necesarios para actualizar un libro")
	Token("jwt")
	Attribute("_id", String, "ID del libro a actualizar")
	Attribute("title", String, "Título del libro")
	Attribute("author", String, "Autor del libro")
	Required("title", "author")
})


var _ = API("Books V2", func() {
	Title("Manejo de libros")
	Description("Manejo de libros utilizando MongoDB")
	Version("2.0")

	Server("booksV2", func() {
		Description("Servidor de manejo de libros")
		Host("localhost", func() {
			URI("http://localhost:8080")
		})
	})
})


var _ = Service("books_v2", func() {
	Description("Servicio para gestión de libros")
	Security(JWTAuth)

	Method("create", func() {
		Description("Crear un nuevo libro")
		Payload(BookPayloadV2)
		Result(Empty)
		HTTP(func() {
			POST("v2/books")
			Response(StatusOK)
			Response(StatusUnauthorized)
			Response(StatusInternalServerError)
		})
	})

	Method("update", func() {
		Description("Actualizar un libro existente")
		Payload(UpdateBookPayloadV2)
		Result(Empty)
		HTTP(func() {
			PUT("v2/books")
			Response(StatusOK)
			Response(StatusUnauthorized)
			Response(StatusNotFound)
			Response(StatusInternalServerError)
		})
	})

	Method("list", func() {
		Description("Listar todos los libros")
		Payload(func() {
			Token("jwt")
		})
		Result(ArrayOf(BookV2))
		HTTP(func() {
			GET("v2/books")
			Response(StatusOK)
			Response(StatusUnauthorized)
			Response(StatusInternalServerError)
		})

	})

	Method("delete", func() {
		Description("Eliminar un libro")
		Payload(func() {
			Attribute("_id", String, "ID del libro a eliminar")
			Token("jwt")
			Required("_id")
		})

		Result(Empty)
		HTTP(func() {
			DELETE("v2/books/{_id}")
			Response(StatusOK)
			Response(StatusUnauthorized)
			Response(StatusInternalServerError)
			Response(StatusNotFound)
		})
	})

})
