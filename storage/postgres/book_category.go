package postgres

import (
	"github.com/jaloldinov/books-catalog/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type bookCategoryRepo struct {
	db *sqlx.DB
}

func (r *bookCategoryRepo) CreateBookCategory(entity models.BookCategory) (string, error) {
	var resp string

	query := `INSERT INTO book_category (
		id,
		category_name,
		created_at,
		updated_at
	) VALUES (
		$1,
		$2,
		$3,
		$4
	) RETURNING id;`

	row := r.db.QueryRow(query, entity.ID, entity.CategoryName, entity.CreatedAt, entity.UpdatedAt)

	if err := row.Scan(&resp); err != nil {
		return "", err
	}

	return resp, nil
}

func (r *bookCategoryRepo) GetBookCategory(id string) (models.BookCategory, error) {
	var resp models.BookCategory
	query := `
				SELECT
					id,
					category_name,
					created_at,
					updated_at
				FROM
					book_category
				WHERE id=$1;
			`

	row := r.db.QueryRow(query, id)
	if err := row.Scan(
		&resp.ID,
		&resp.CategoryName,
		&resp.CreatedAt,
		&resp.UpdatedAt,
	); err != nil {
		return resp, err
	}

	return resp, nil
}

func (r *bookCategoryRepo) GetAllBookCategories(queryParam models.ApplicationQueryParamModel) ([]models.BookCategory, error) {
	var resp []models.BookCategory = []models.BookCategory{}

	params := make(map[string]interface{})

	query := `SELECT
		id,
		category_name,
		created_at,
		updated_at
	FROM
		book_category`
	filter := " WHERE 1=1"
	offset := " OFFSET 0"
	limit := " LIMIT 10"

	if len(queryParam.Search) > 0 {
		params["search"] = queryParam.Search
		filter += " AND (category_name ILIKE '%' || :search || '%')"
	}

	if queryParam.Offset > 0 {
		params["offset"] = queryParam.Offset
		offset = " OFFSET :offset"
	}

	if queryParam.Limit > 0 {
		params["limit"] = queryParam.Limit
		limit = " LIMIT :limit"
	}

	countQuery := "SELECT count(1) FROM book_category" + filter
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
		var category models.BookCategory
		err = rows.Scan(
			&category.ID,
			&category.CategoryName,
			&category.CreatedAt,
			&category.UpdatedAt,
		)

		if err != nil {
			return resp, err
		}
		resp = append(resp, category)
	}

	return resp, nil
}

func (r *bookCategoryRepo) UpdateBookCategory(entity *models.UpdateBookCategory, id string) (int64, error) {
	params := make(map[string]interface{})
	params["id"] = id

	query := `UPDATE book_category SET `

	if len(entity.CategoryName) > 0 {
		params["category_name"] = entity.CategoryName
		query += `category_name = :category_name,`
	}

	query += `updated_at = now() WHERE id =:id`

	respult, err := r.db.NamedExec(query, params)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := respult.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, err
}

func (r *bookCategoryRepo) DeleteBookCategory(id string) (int64, error) {
	query := `DELETE FROM book_category WHERE id = $1`

	respult, err := r.db.Exec(query, id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := respult.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, err
}
