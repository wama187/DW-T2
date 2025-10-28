package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"app/server/apps/user"
	"app/server/apps/book/v1"
	"app/server/apps/book/v2"
	"app/server/databases"
	"app/server/apps/auth"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2"

	goahttp "goa.design/goa/v3/http"
	usersapi "app/server/gen/users"
	userssrv "app/server/gen/http/users/server"
	booksapi "app/server/gen/books"
	booksrv "app/server/gen/http/books/server"
	booksapiv2 "app/server/gen/books_v2"
	booksrvv2 "app/server/gen/http/books_v2/server"
	authapi "app/server/gen/auth"
	authsrv "app/server/gen/http/auth/server"


	"github.com/rs/cors"
)

func main() {
	err := godotenv.Load("/app/.env")
	if err != nil {
		log.Fatal("Error cargando el archivo .env")
	}

	cfg := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT"),
		ClientSecret: os.Getenv("GOOGLE_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	
	port := "8080"
	jwtSecret := "a7BzjaA19AB187zmj99MlZaEMN"

	// Conexión a la base de datos
	postgreDB, err := databases.ConnectPostgres()
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}

	if err = databases.Migrate(); err != nil {
		log.Fatalf("Error al migrar la base de datos: %v", err)
	}

	mongoDB, err := databases.ConnectMongo()
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos Mongo: %v", err)
	}

	// Inicializar servicios
	userService := user.NewUserService(
		user.NewUserRepositoryPostgres(postgreDB),
		jwtSecret,
		24*time.Hour,
	)

	bookService := book_v1.NewService(
		book_v1.NewBookRepositoryPostgres(postgreDB),
		jwtSecret,
	)

	bookServiceV2 := book_v2.NewService(
		book_v2.NewBookRepositoryMongo(mongoDB, "db", "books"),
		jwtSecret,
	)

	authService := auth.NewAuth(
		user.NewUserRepositoryPostgres(postgreDB),
	 	jwtSecret, 
	 	24*time.Hour,
		cfg,
	)

	// Crear endpoints
	usersEndpoints := usersapi.NewEndpoints(userService)
	booksEndpoints := booksapi.NewEndpoints(bookService)
	booksEndpointsV2 := booksapiv2.NewEndpoints(bookServiceV2)
	authEndpoints := authapi.NewEndpoints(authService)

	// Crear router HTTP
	mux := goahttp.NewMuxer()

	// Crear servidores Goa
	usersServer := userssrv.New(usersEndpoints, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, nil, nil)
	booksServer := booksrv.New(booksEndpoints, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, nil, nil)
	booksServerV2 := booksrvv2.New(booksEndpointsV2, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, nil, nil)
	authServer := authsrv.New(authEndpoints, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, nil, nil)

	// Montar endpoints
	userssrv.Mount(mux, usersServer)
	booksrv.Mount(mux, booksServer)
	booksrvv2.Mount(mux, booksServerV2)
	authsrv.Mount(mux, authServer)
	log.Println("Endpoints montados correctamente")

	// Configurar CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(mux)

	// Configurar servidor HTTP
	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Ejecutar servidor
	go func() {
		log.Printf("Servidor iniciado en http://localhost:%s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error en el servidor HTTP: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Apagando servidor...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Error al apagar el servidor: %v", err)
	}

	log.Println("Servidor detenido correctamente")
}
