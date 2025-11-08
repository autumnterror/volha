package psql

import (
	"fmt"
	"log"
	"productService/internal/utils/format"
	"productService/internal/views"
)

func (d Driver) GetAllColors() ([]views.Color, error) {
	const op = "PostgresDb.GetAllColors"
	var list []views.Color

	rows, err := d.Driver.Query(`SELECT id, name, hex FROM colors`)
	if err != nil {
		return nil, format.Error(op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var c views.Color
		if err := rows.Scan(&c.Id, &c.Name, &c.Hex); err != nil {
			log.Println(format.Error(op, err))
			continue
		}
		list = append(list, c)
	}
	return list, nil
}

func (d Driver) CreateColor(c *views.Color) error {
	const op = "PostgresDb.CreateColor"
	query := `INSERT INTO colors (id, name, hex) VALUES ($1, $2, $3)`
	_, err := d.Driver.Exec(query, c.Id, c.Name, c.Hex)
	return format.Error(op, err)
}

func (d Driver) UpdateColor(c *views.Color, id string) error {
	const op = "PostgresDb.UpdateColor"
	query := `UPDATE colors SET name = $2, hex = $3 WHERE id = $1`

	result, err := d.Driver.Exec(query, id, c.Name, c.Hex)
	if err != nil {
		return format.Error(op, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return format.Error(op, fmt.Errorf("color with id %s not found", id))
	}
	return nil
}

func (d Driver) DeleteColor(id string) error {
	const op = "PostgresDb.DeleteColor"
	_, err := d.Driver.Exec(`DELETE FROM colors WHERE id = $1`, id)
	return format.Error(op, err)
}
