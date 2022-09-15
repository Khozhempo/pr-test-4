package books

import (
	"context"
	"pr-test-4/books/models"
)

type Repository interface {
	ReadBooks(ctx context.Context, sum int) ([]*models.Book, error)
	RemoveBook(ctx context.Context, bookId int) error
}
