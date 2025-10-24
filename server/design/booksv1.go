package design

import . "goa.design/goa/v3/dsl"

var Book = Type("BookV1", func() {
	Description("Representa un libro")
	Attribute("id", UInt, "ID único del libro")
	Attribute("title", String, "Título del libro")
	Attribute("author", String, "Autor del libro")
	Attribute("owner_id", UInt, "ID del usuario que posee el libro")
	Required("id", "title", "author", "owner_id")
})


var BookPayload = Type("BookPayloadV1", func() {
	Description("Datos necesarios para un libro")
	Token("jwt")
	Attribute("id", UInt, "ID único del libro")
	Attribute("title", String, "Título del libro")
	Attribute("author", String, "Autor del libro")
	Required("title", "author")
})

var UpdateBookPayload = Type("UpdateBookPayloadV1", func() {
	Description("Datos necesarios para actualizar un libro")
	Token("jwt")
	Attribute("id", UInt, "ID del libro a actualizar")
	Attribute("title", String, "Título del libro")
	Attribute("author", String, "Autor del libro")
	Required("title", "author")
})

var _ = API("Books V1", func() {
	Title("Manejo de libros")
	Description("Manejo de libros utilizando Postgre")
	Version("1.0")

	Server("booksV1", func() {
		Description("Servidor de manejo de libros")
		Host("localhost", func() {
			URI("http://localhost:8080")
		})
	})
})

var _ = Service("books", func() {
	Description("Servicio para gestión de libros")
	Security(JWTAuth)

	Method("create", func() {
		Description("Crear un nuevo libro")
		Payload(BookPayload)
		Result(Empty)
		HTTP(func() {
			POST("v1/books")
			Response(StatusOK)
			Response(StatusUnauthorized)
			Response(StatusInternalServerError)
		})
	})

	Method("update", func() {
		Description("Actualizar un libro existente")
		Payload(UpdateBookPayload)
		Result(Empty)
		HTTP(func() {
			PUT("v1/books")
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
		Result(ArrayOf(Book))
		HTTP(func() {
			GET("v1/books")
			Response(StatusOK)
			Response(StatusUnauthorized)
			Response(StatusInternalServerError)
		})

	})

	Method("delete", func() {
		Description("Eliminar un libro")
		Payload(func() {
			Attribute("id", UInt, "ID del libro a eliminar")
			Token("jwt")
			Required("id")
		})

		Result(Empty)
		HTTP(func() {
			DELETE("v1/books/{id}")
			Response(StatusOK)
			Response(StatusUnauthorized)
			Response(StatusInternalServerError)
			Response(StatusNotFound)
		})
	})

})
