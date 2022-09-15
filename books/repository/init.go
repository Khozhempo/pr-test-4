package repository

import (
	"sync"
	"time"
)

type Book struct {
	Title         string
	Author        string
	PublisherYear time.Time
}

type BooksRepository struct {
	Books map[int]Book
	sync.RWMutex
}

var storage BooksRepository

func init() {
	storage.Books = make(map[int]Book)
	fullFill()
}

func NewRepository() *BooksRepository {
	return &storage
}

func fullFill() {
	storage.Books[1] = Book{
		Title:         "Book1",
		Author:        "Author1",
		PublisherYear: time.Now(),
	}

	storage.Books[2] = Book{
		Title:         "Book2",
		Author:        "Author2",
		PublisherYear: time.Now().Add(time.Hour * 12),
	}

	storage.Books[3] = Book{
		Title:         "Book3",
		Author:        "Author3",
		PublisherYear: time.Now().Add(time.Hour * 12 * 2),
	}

	storage.Books[4] = Book{
		Title:         "Book4",
		Author:        "Author4",
		PublisherYear: time.Now().Add(time.Hour * 12 * 4),
	}

}
