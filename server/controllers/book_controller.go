package controllers

import (
	"encoding/json"
	"github.com/Sh1wOo/client-server-solid-golang-grpc/library"
	"github.com/Sh1wOo/client-server-solid-golang-grpc/services"
	"net/http"
)

type BookController struct {
	libraryService *services.LibraryService
}

func NewBookController(libraryService *services.LibraryService) *BookController {
	return &BookController{libraryService: libraryService}
}

func (c *BookController) GetBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	booksResponse, err := c.libraryService.GetAllBooks(r.Context(), &library.AllBooksRequest{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(booksResponse.Book)
}

func (c *BookController) AddBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var book library.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bookRequest := &library.BookRequest{
		Book: &book,
	}

	bookResponse, err := c.libraryService.CreateBook(r.Context(), bookRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bookResponse.Book)
}
