package psql

import (
	"fmt"
	"log"
	"productService/internal/utils/format"
	"productService/internal/views"
)

func (d Driver) GetAllCountries() ([]views.Country, error) {
	const op = "PostgresDb.GetAllCountries"
	var list []views.Country

	rows, err := d.Driver.Query(`SELECT id, title, friendly FROM countries`)
	if err != nil {
		return nil, format.Error(op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var c views.Country
		if err := rows.Scan(&c.Id, &c.Title, &c.Friendly); err != nil {
			log.Println(format.Error(op, err))
			continue
		}
		list = append(list, c)
	}
	return list, nil
}

func (d Driver) CreateCountry(c *views.Country) error {
	const op = "PostgresDb.CreateCountry"
	query := `INSERT INTO countries (id, title, friendly) VALUES ($1, $2, $3)`
	_, err := d.Driver.Exec(query, c.Id, c.Title, c.Friendly)
	return format.Error(op, err)
}

func (d Driver) UpdateCountry(c *views.Country, id string) error {
	const op = "PostgresDb.UpdateCountry"
	query := `UPDATE countries SET title = $2, friendly = $3 WHERE id = $1`

	result, err := d.Driver.Exec(query, id, c.Title, c.Friendly)
	if err != nil {
		return format.Error(op, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return format.Error(op, fmt.Errorf("country with id %s not found", id))
	}
	return nil
}

func (d Driver) DeleteCountry(id string) error {
	const op = "PostgresDb.DeleteCountry"
	_, err := d.Driver.Exec(`DELETE FROM countries WHERE id = $1`, id)
	return format.Error(op, err)
}
