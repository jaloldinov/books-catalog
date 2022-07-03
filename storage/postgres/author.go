package postgres

import (
	"github.com/jaloldinov/books-catalog/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type authorRepo struct {
	db *sqlx.DB
}

func (r *authorRepo) CreateAuthor(entity models.Author) (string, error) {
	var resp string

	query := `INSERT INTO author (
		id,
		firstname,
		lastname,
		created_at,
		updated_at
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5
	) RETURNING id; `

	row := r.db.QueryRow(query, entity.ID, entity.Firstname, entity.Lastname, entity.CreatedAt, entity.UpdatedAt)

	if err := row.Scan(&resp); err != nil {
		return "", err
	}

	return resp, nil
}

func (r *authorRepo) GetAuthor(id string) (models.Author, error) {
	var resp models.Author

	query := `
		SELECT
			id,
			firstname,
			lastname,
			created_at,
			updated_at
		FROM author
		WHERE id = $1;
	`

	row := r.db.QueryRow(query, id)
	if err := row.Scan(
		&resp.ID,
		&resp.Firstname,
		&resp.Lastname,
		&resp.CreatedAt,
		&resp.UpdatedAt,
	); err != nil {
		return resp, err
	}

	return resp, nil
}

func (r *authorRepo) GetAllAuthors(queryParam models.ApplicationQueryParamModel) ([]models.Author, error) {
	var resp []models.Author = []models.Author{}

	params := make(map[string]interface{})

	query := `SELECT
		id,
		firstname,
		lastname,
		created_at,
		updated_at
	FROM
		author`
	filter := " WHERE 1=1"
	offset := " OFFSET 0"
	limit := " LIMIT 10"

	if len(queryParam.Search) > 0 {
		params["search"] = queryParam.Search
		filter += " AND (firstname ILIKE '%' || :search || '%') OR (lastname ILIKE '%' || :search || '%')"
	}

	if queryParam.Offset > 0 {
		params["offset"] = queryParam.Offset
		offset = " OFFSET :offset"
	}

	if queryParam.Limit > 0 {
		params["limit"] = queryParam.Limit
		limit = " LIMIT :limit"
	}

	countQuery := "SELECT count(1) FROM author" + filter
	row, err := r.db.NamedQuery(countQuery, params)
	if err != nil {
		return resp, err
	}
	defer row.Close()

	q := query + filter + offset + limit
	rows, err := r.db.NamedQuery(q, params)
	if err != nil {
		return resp, err
	}
	defer rows.Close()
	for rows.Next() {
		var author models.Author
		err = rows.Scan(
			&author.ID,
			&author.Firstname,
			&author.Lastname,
			&author.CreatedAt,
			&author.UpdatedAt,
		)

		if err != nil {
			return resp, err
		}
		resp = append(resp, author)
	}

	return resp, nil
}

func (r *authorRepo) UpdateAuthor(entity models.UpdateAuthor, id string) (int64, error) {
	params := make(map[string]interface{})
	params["id"] = id

	query := `UPDATE author SET `

	if len(entity.Firstname) > 0 {
		params["firstname"] = entity.Firstname
		query += `firstname = :firstname,`
	}

	if len(entity.Lastname) > 0 {
		params["lastname"] = entity.Lastname
		query += `lastname = :lastname,`
	}

	query += `updated_at = now() WHERE id =:id`

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

func (r *authorRepo) DeleteAuthor(id string) (int64, error) {
	query := `DELETE FROM author WHERE id = $1`

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
