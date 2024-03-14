package database

import (
	"context"
	"github.com/Sh1wOo/client-server-solid-golang-grpc/library"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
)

type Book struct {
	gorm.Model
	Title  string
	Author string
	Year   int
}

type Database struct {
	db *gorm.DB
}

func NewDatabase(dsn string) (*Database, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// --------------- Auto Migrate ---------------
	err = db.AutoMigrate(&Book{})
	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

func (d *Database) Close() {
	sqlDB, err := d.db.DB()
	if err != nil {
		// Обработка ошибки
		return
	}
	sqlDB.Close()
}

func (d *Database) CreateBook(ctx context.Context, request *library.BookRequest) (*library.BookResponse, error) {
	book := Book{
		Title:  request.Book.Title,
		Author: request.Book.Author,
		Year:   int(request.Book.Year),
	}

	result := d.db.Create(&book)
	if result.Error != nil {
		return nil, result.Error
	}

	return &library.BookResponse{
		Book: &library.Book{
			Id:     strconv.Itoa(int(book.ID)),
			Title:  book.Title,
			Author: book.Author,
			Year:   int32(book.Year),
		},
	}, nil
}

func (d *Database) GetAllBooks(ctx context.Context, request *library.AllBooksRequest) (*library.AllBooksResponse, error) {
	var books []Book

	result := d.db.Find(&books)

	if result.Error != nil {
		return nil, result.Error
	}

	var booksProto []*library.Book

	for _, book := range books {
		booksProto = append(booksProto, &library.Book{
			Id:     strconv.Itoa(int(book.ID)),
			Title:  book.Title,
			Author: book.Author,
			Year:   int32(book.Year),
		})
	}

	return &library.AllBooksResponse{Book: booksProto}, nil
}
