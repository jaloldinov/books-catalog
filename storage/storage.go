package storage

import "github.com/jaloldinov/books-catalog/models"

type StorageI interface {
	CloseDB() error
	BookCategoryRepo() BookCategoryI
	BookRepo() BookI
	AuthorRepo() AuthorI
}

type BookCategoryI interface {
	GetBookCategory(id string) (models.BookCategory, error)
	GetAllBookCategories(queryParam models.ApplicationQueryParamModel) ([]models.BookCategory, error)
	CreateBookCategory(details models.BookCategory) (string, error)
	UpdateBookCategory(details *models.UpdateBookCategory, id string) (int64, error)
	DeleteBookCategory(id string) (int64, error)
}

type BookI interface {
	GetBook(id string) (models.Book, error)
	GetAllBooks(queryParam models.ApplicationQueryParamModel) ([]models.Book, error)
	CreateBook(details models.Book) (string, error)
	UpdateBook(details models.UpdateBook, id string) (int64, error)
	DeleteBook(id string) (int64, error)
}

type AuthorI interface {
	GetAuthor(id string) (models.Author, error)
	GetAllAuthors(queryParam models.ApplicationQueryParamModel) ([]models.Author, error)
	CreateAuthor(details models.Author) (string, error)
	UpdateAuthor(details models.UpdateAuthor, id string) (int64, error)
	DeleteAuthor(id string) (int64, error)
}
