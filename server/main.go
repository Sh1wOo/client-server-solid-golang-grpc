package main

import (
	"github.com/Sh1wOo/client-server-solid-golang-grpc/controllers"
	"github.com/Sh1wOo/client-server-solid-golang-grpc/database"
	"github.com/Sh1wOo/client-server-solid-golang-grpc/library"
	"github.com/Sh1wOo/client-server-solid-golang-grpc/services"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

func main() {
	// --------------- Connect to DataBase ---------------
	dsn := "host=localhost user=admin password=123456 dbname=library port=5432 sslmode=disable"

	db, err := database.NewDatabase(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// --------------- Server Register ---------------
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	libraryService := services.NewLibraryService(db)
	library.RegisterLibraryServiceServer(grpcServer, libraryService)

	// --------------- HTTP Server ---------------
	// cors
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		Debug:          true,
	})

	httpServer := &http.Server{
		Addr:    ":8090",
		Handler: corsMiddleware.Handler(http.DefaultServeMux),
	}

	bookController := controllers.NewBookController(libraryService)
	http.HandleFunc("/books", bookController.GetBooks)
	http.HandleFunc("/books/add", bookController.AddBook)

	// Start gRPC server
	go func() {
		log.Printf("gRPC server listening at %v", lis.Addr())
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	}()

	// Start HTTP server
	log.Printf("HTTP server listening at %s", httpServer.Addr)
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("failed to serve HTTP: %v", err)
	}
}
