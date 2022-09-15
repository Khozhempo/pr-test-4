package repository

import (
	"context"
	"errors"
	"pr-test-4/books/models"
)

func (r BooksRepository) ReadBooks(ctx context.Context, sum int) ([]*models.Book, error) {
	var readFromStorage = make([]*models.Book, 0, sum)
	var counter int

	storage.RLock()
	defer storage.RUnlock()

	if len(storage.Books) == 0 {
		return nil, errors.New("no books in base")
	}

	for key, value := range storage.Books {
		readFromStorage = append(
			readFromStorage,
			&models.Book{Id: key, Title: value.Title, Author: value.Author, PublisherYear: value.PublisherYear},
		)

		counter++
		if counter == sum {
			break
		}
	}
	return readFromStorage, nil
}

func (r BooksRepository) RemoveBook(ctx context.Context, bookId int) error {

	if _, ok := storage.Books[bookId]; ok {
		storage.Lock()
		delele(storage.Books, bookId)
		storage.Unlock()
		return nil
	}

	return errors.New("no books found with requested id")
}

func delele(books map[int]Book, id int) {
	delete(books, id)
}
