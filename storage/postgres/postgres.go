package postgres

import (
	"log"

	"github.com/jaloldinov/books-catalog/storage"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type postgres struct {
	db               *sqlx.DB
	authorRepo       *authorRepo
	bookCategoryRepo *bookCategoryRepo
	bookRepo         *bookRepo
}

func NewPostgres(str string) storage.StorageI {
	db, err := sqlx.Connect("postgres", str)
	if err != nil {
		log.Fatalf("Could not connect to db: %e", err)
		return nil
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Could not connect to db error ping: %e", err)
		return nil
	}

	return &postgres{
		db:               db,
		authorRepo:       &authorRepo{db},
		bookRepo:         &bookRepo{db},
		bookCategoryRepo: &bookCategoryRepo{db},
	}
}

func (pg *postgres) CloseDB() error {
	return pg.db.Close()
}

func (pg *postgres) AuthorRepo() storage.AuthorI {
	return pg.authorRepo
}

func (pg *postgres) BookCategoryRepo() storage.BookCategoryI {
	return pg.bookCategoryRepo
}

func (pg *postgres) BookRepo() storage.BookI {
	return pg.bookRepo
}
