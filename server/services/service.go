package services

import (
	"context"
	"github.com/Sh1wOo/client-server-solid-golang-grpc/database"
	"github.com/Sh1wOo/client-server-solid-golang-grpc/library"
)

type LibraryService struct {
	library.UnimplementedLibraryServiceServer
	db *database.Database
}

func NewLibraryService(db *database.Database) *LibraryService {
	return &LibraryService{db: db}
}

func (s *LibraryService) CreateBook(ctx context.Context, request *library.BookRequest) (*library.BookResponse, error) {
	return s.db.CreateBook(ctx, request)
}

func (s *LibraryService) GetAllBooks(ctx context.Context, request *library.AllBooksRequest) (*library.AllBooksResponse, error) {
	return s.db.GetAllBooks(ctx, request)
}
