package psql

import (
	"fmt"
	"log"
	"productService/internal/utils/format"
	"productService/internal/views"
)

func (d Driver) GetAllMaterials() ([]views.Material, error) {
	const op = "PostgresDb.GetAllMaterials"
	var list []views.Material

	rows, err := d.Driver.Query(`SELECT id, title FROM materials`)
	if err != nil {
		return nil, format.Error(op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var m views.Material
		if err := rows.Scan(&m.Id, &m.Title); err != nil {
			log.Println(format.Error(op, err))
			continue
		}
		list = append(list, m)
	}
	return list, nil
}

func (d Driver) CreateMaterial(m *views.Material) error {
	const op = "PostgresDb.CreateMaterial"
	query := `INSERT INTO materials (id, title) VALUES ($1, $2)`
	_, err := d.Driver.Exec(query, m.Id, m.Title)
	return format.Error(op, err)
}

func (d Driver) UpdateMaterial(m *views.Material, id string) error {
	const op = "PostgresDb.UpdateMaterial"
	query := `UPDATE materials SET title = $2 WHERE id = $1`

	result, err := d.Driver.Exec(query, id, m.Title)
	if err != nil {
		return format.Error(op, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return format.Error(op, fmt.Errorf("material with id %s not found", id))
	}
	return nil
}

func (d Driver) DeleteMaterial(id string) error {
	const op = "PostgresDb.DeleteMaterial"
	_, err := d.Driver.Exec(`DELETE FROM materials WHERE id = $1`, id)
	return format.Error(op, err)
}
