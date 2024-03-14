package main

import (
	"github.com/Sh1wOo/client-server-solid-golang-grpc/database"
	"github.com/Sh1wOo/client-server-solid-golang-grpc/library"
	"github.com/Sh1wOo/client-server-solid-golang-grpc/services"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// --------------- Connect to DataBase ---------------
	dsn := "host=localhost user=admin password=123456 dbname=library port=5432 sslmode=disable"

	db, err := database.NewDatabase(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// --------------- Server Register ---------------
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	serverRegister := grpc.NewServer()
	libraryService := services.NewLibraryService(db)
	library.RegisterLibraryServiceServer(serverRegister, libraryService)

	log.Printf("server listening at %v", lis.Addr())
	if err := serverRegister.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
