package psql

import (
	"fmt"
	"log"
	"productService/internal/utils/format"
	"productService/internal/views"
)

func (d Driver) GetAllCategories() ([]views.Category, error) {
	const op = "PostgresDb.GetAllCategories"
	var list []views.Category

	rows, err := d.Driver.Query(`SELECT id, title, uri, img FROM categories`)
	if err != nil {
		return nil, format.Error(op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var c views.Category
		if err := rows.Scan(&c.Id, &c.Title, &c.Uri, &c.Img); err != nil {
			log.Println(format.Error(op, err))
			continue
		}
		list = append(list, c)
	}
	return list, nil
}

func (d Driver) CreateCategory(c *views.Category) error {
	const op = "PostgresDb.CreateCategory"
	query := `INSERT INTO categories (id, title, uri, img) VALUES ($1, $2, $3, $4)`
	_, err := d.Driver.Exec(query, c.Id, c.Title, c.Uri, c.Img)
	return format.Error(op, err)
}

func (d Driver) UpdateCategory(c *views.Category, id string) error {
	const op = "PostgresDb.UpdateCategory"
	query := `UPDATE categories SET title = $2, uri = $3, img = $4 WHERE id = $1`

	result, err := d.Driver.Exec(query, id, c.Title, c.Uri, c.Img)
	if err != nil {
		return format.Error(op, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return format.Error(op, fmt.Errorf("category with id %s not found", id))
	}
	return nil
}

func (d Driver) DeleteCategory(id string) error {
	const op = "PostgresDb.DeleteCategory"
	_, err := d.Driver.Exec(`DELETE FROM categories WHERE id = $1`, id)
	return format.Error(op, err)
}
