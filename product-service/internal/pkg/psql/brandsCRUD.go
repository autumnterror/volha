package psql

import (
	"fmt"
	"log"
	"productService/internal/utils/format"
	"productService/internal/views"
)

func (d Driver) GetAllBrands() ([]views.Brand, error) {
	const op = "PostgresDb.GetAllBrands"

	var list []views.Brand
	rows, err := d.Driver.Query(`SELECT id, name FROM brands`)
	if err != nil {
		return nil, format.Error(op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var b views.Brand
		if err := rows.Scan(&b.Id, &b.Name); err != nil {
			log.Println(format.Error(op, err))
			continue
		}
		list = append(list, b)
	}

	return list, nil
}

func (d Driver) CreateBrand(b *views.Brand) error {
	const op = "PostgresDb.CreateBrand"

	query := `INSERT INTO brands (id, name) VALUES ($1, $2)`
	_, err := d.Driver.Exec(query, b.Id, b.Name)
	if err != nil {
		return format.Error(op, err)
	}

	return nil
}

func (d Driver) UpdateBrand(b *views.Brand, id string) error {
	const op = "PostgresDb.UpdateBrand"

	query := `UPDATE brands SET name = $2 WHERE id = $1`
	result, err := d.Driver.Exec(query, id, b.Name)
	if err != nil {
		return format.Error(op, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return format.Error(op, err)
	}
	if rowsAffected == 0 {
		return format.Error(op, fmt.Errorf("brand with id %s not found", id))
	}

	return nil
}

func (d Driver) DeleteBrand(id string) error {
	const op = "PostgresDb.DeleteBrand"

	query := `DELETE FROM brands WHERE id = $1`
	_, err := d.Driver.Exec(query, id)
	if err != nil {
		return format.Error(op, err)
	}
	return nil
}
