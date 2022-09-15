package books

import (
	"context"
	"pr-test-4/books/models"
)

type UseCase interface {
	GetBooks(ctx context.Context) ([]*models.Book, error)
	DeleteBook(ctx context.Context, bookId int) error
}
