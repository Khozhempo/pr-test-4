package usecase

import (
	"context"
	"pr-test-4/books"
	"pr-test-4/books/models"
)

func (b BooksUseCase) GetBooks(ctx context.Context) ([]*models.Book, error) {
	return b.booksRepo.ReadBooks(ctx, 3)
}

func (b BooksUseCase) DeleteBook(ctx context.Context, bookId int) error {
	return b.booksRepo.RemoveBook(ctx, bookId)
}

type BooksUseCase struct {
	booksRepo books.Repository
}

func NewBooksUseCase(booksRepo books.Repository) *BooksUseCase {
	return &BooksUseCase{
		booksRepo: booksRepo,
	}
}
