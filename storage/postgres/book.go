package postgres

import (
	"errors"

	"github.com/jaloldinov/books-catalog/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type bookRepo struct {
	db *sqlx.DB
}

func (r *bookRepo) CreateBook(details models.Book) (string, error) {
	var resp string
	var countCategory int
	var countAuthor int

	q1 := `SELECT count(1) FROM book_category WHERE id=$1;`
	row1 := r.db.QueryRow(q1, details.CategoryID)
	if err := row1.Scan(&countCategory); err != nil {
		return resp, err
	}

	if countCategory < 1 {
		return resp, errors.New("there is no category_name with the given id")
	}

	q2 := `SELECT count(1) FROM author WHERE id=$1;`
	row2 := r.db.QueryRow(q2, details.AuthorID)
	if err := row2.Scan(&countAuthor); err != nil {
		return resp, err
	}
	if countAuthor < 1 {
		return resp, errors.New("there is no author with the given id")
	}

	query := `INSERT INTO book (
		id,
		book_name,
		category_id,
		author_id,
		created_at,
		updated_at
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6
	) RETURNING id;`

	row := r.db.QueryRow(query,
		details.ID,
		details.BookName,
		details.CategoryID,
		details.AuthorID,
		details.CreatedAt,
		details.UpdatedAt,
	)

	if err := row.Scan(&resp); err != nil {
		return "", err
	}

	return resp, nil
}

func (r *bookRepo) GetBook(id string) (models.Book, error) {
	var resp models.Book
	query := `
				SELECT
					id,
					book_name,
					author_id,
					category_id,
					created_at,
					updated_at
				FROM
					book
				WHERE id=$1;
			`

	row := r.db.QueryRow(query, id)
	if err := row.Scan(
		&resp.ID,
		&resp.BookName,
		&resp.AuthorID,
		&resp.CategoryID,
		&resp.CreatedAt,
		&resp.UpdatedAt,
	); err != nil {
		return resp, err
	}

	return resp, nil
}

func (r *bookRepo) GetAllBooks(queryParam models.ApplicationQueryParamModel) ([]models.Book, error) {
	var resp []models.Book = []models.Book{}

	params := make(map[string]interface{})

	query := `SELECT
		id,
		category_id,
		author_id,
		book_name,
		created_at,
		updated_at
	FROM
		book`
	filter := " WHERE 1=1"
	offset := " OFFSET 0"
	limit := " LIMIT 10"

	if len(queryParam.Search) > 0 {
		params["search"] = queryParam.Search
		filter += " AND (name ILIKE '%' || :search || '%')"
	}

	if queryParam.Offset > 0 {
		params["offset"] = queryParam.Offset
		offset = " OFFSET :offset"
	}

	if queryParam.Limit > 0 {
		params["limit"] = queryParam.Limit
		limit = " LIMIT :limit"
	}

	countQuery := "SELECT count(1) FROM book" + filter
	row, err := r.db.NamedQuery(countQuery, params)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	q := query + filter + offset + limit
	rows, err := r.db.NamedQuery(q, params)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var book models.Book
		err = rows.Scan(
			&book.ID,
			&book.CategoryID,
			&book.AuthorID,
			&book.BookName,
			&book.CreatedAt,
			&book.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
		resp = append(resp, book)
	}

	return resp, nil
}

func (r *bookRepo) UpdateBook(entity models.UpdateBook, id string) (int64, error) {
	params := make(map[string]interface{})
	params["id"] = id

	query := `UPDATE book SET `

	if len(entity.CategoryID) > 0 {
		params["category_id"] = entity.CategoryID
		query += `category_id = :category_id,`
	}

	if len(entity.AuthorID) > 0 {
		params["author_id"] = entity.AuthorID
		query += `author_id = :author_id,`
	}

	if len(entity.BookName) > 0 {
		params["book_name"] = entity.BookName
		query += `book_name = :book_name,`
	}

	query += `updated_at =  now() WHERE id =:id`

	result, err := r.db.NamedExec(query, params)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, err
}

func (r *bookRepo) DeleteBook(id string) (int64, error) {
	query := `DELETE FROM book WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, err
}
