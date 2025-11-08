package psql

import (
	"fmt"
	"github.com/lib/pq"
	"log"
	"productService/internal/utils/format"
	"productService/internal/views"
)

func (d Driver) GetAllProductColorPhotos() ([]views.ProductColorPhotos, error) {
	const op = "PostgresDb.GetAllProductColorPhotos"
	var list []views.ProductColorPhotos

	rows, err := d.Driver.Query(`SELECT product_id, color_id, photos FROM product_color_photos`)
	if err != nil {
		return nil, format.Error(op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var pcp views.ProductColorPhotos
		if err := rows.Scan(&pcp.ProductId, &pcp.ColorId, pq.Array(&pcp.Photos)); err != nil {
			log.Println(format.Error(op, err))
			continue
		}
		list = append(list, pcp)
	}
	return list, nil
}

func (d Driver) CreateProductColorPhotos(pcp *views.ProductColorPhotos) error {
	const op = "PostgresDb.CreateProductColorPhotos"
	query := `INSERT INTO product_color_photos (product_id, color_id, photos) VALUES ($1, $2, $3)`
	_, err := d.Driver.Exec(query, pcp.ProductId, pcp.ColorId, pq.Array(pcp.Photos))
	return format.Error(op, err)
}

func (d Driver) UpdateProductColorPhotos(pcp *views.ProductColorPhotos) error {
	const op = "PostgresDb.UpdateProductColorPhotos"
	query := `UPDATE product_color_photos SET photos = $3 WHERE product_id = $1 AND color_id = $2`

	result, err := d.Driver.Exec(query, pcp.ProductId, pcp.ColorId, pq.Array(pcp.Photos))
	if err != nil {
		return format.Error(op, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return format.Error(op, fmt.Errorf("record with product_id=%s and color_id=%s not found", pcp.ProductId, pcp.ColorId))
	}
	return nil
}

func (d Driver) DeleteProductColorPhotos(productId, colorId string) error {
	const op = "PostgresDb.DeleteProductColorPhotos"
	_, err := d.Driver.Exec(`DELETE FROM product_color_photos WHERE product_id = $1 AND color_id = $2`, productId, colorId)
	return format.Error(op, err)
}

func (d Driver) GetPhotosByProductAndColor(productId, colorId string) ([]string, error) {
	const op = "PostgresDb.GetPhotosByProductAndColor"
	var photos []string

	err := d.Driver.QueryRow(
		`SELECT photos FROM product_color_photos WHERE product_id = $1 AND color_id = $2`,
		productId, colorId,
	).Scan(pq.Array(&photos))

	if err != nil {
		return nil, format.Error(op, err)
	}
	return photos, nil
}
