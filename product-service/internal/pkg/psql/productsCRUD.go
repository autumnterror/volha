package psql

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"productService/internal/utils/format"
	"productService/internal/views"
	"strings"

	"github.com/lib/pq"
)

type Repository interface {
	GetAllColors() ([]views.Color, error)
	CreateColor(c *views.Color) error
	UpdateColor(c *views.Color, id string) error
	DeleteColor(id string) error
	GetAllMaterials() ([]views.Material, error)
	CreateMaterial(m *views.Material) error
	UpdateMaterial(m *views.Material, id string) error
	DeleteMaterial(id string) error
	GetAllCountries() ([]views.Country, error)
	GetDictionariesByCategory(id string) (*views.Dictionaries, error)
	CreateCountry(c *views.Country) error
	UpdateCountry(c *views.Country, id string) error
	DeleteCountry(id string) error
	GetAllCategories() ([]views.Category, error)
	CreateCategory(c *views.Category) error
	UpdateCategory(c *views.Category, id string) error
	DeleteCategory(id string) error
	GetAllBrands() ([]views.Brand, error)
	CreateBrand(b *views.Brand) error
	UpdateBrand(b *views.Brand, id string) error
	DeleteBrand(id string) error
	GetAllProducts() ([]views.Product, error)
	GetProductById(id string) (*views.Product, error)
	CreateProduct(p *views.ProductId) error
	UpdateProduct(p *views.ProductId, id string) error
	DeleteProduct(id string) error
	FilterProducts(filter *views.ProductFilter) ([]views.Product, error)
	GetDictionaries() (*views.Dictionaries, error)
	SearchProducts(filter *views.ProductSearch) ([]views.Product, error)
	GetAllProductColorPhotos() ([]views.ProductColorPhotos, error)
	CreateProductColorPhotos(pcp *views.ProductColorPhotos) error
	UpdateProductColorPhotos(pcp *views.ProductColorPhotos) error
	DeleteProductColorPhotos(productId, colorId string) error
	GetPhotosByProductAndColor(productId, colorId string) ([]string, error)
}

type SqlRepo interface {
	Query(query string, args ...any) (*sql.Rows, error)
	Exec(query string, args ...any) (sql.Result, error)
	QueryRow(query string, args ...any) *sql.Row
}

type Driver struct {
	Driver SqlRepo
}

func fetchBrand(db SqlRepo, id string) views.Brand {
	var b views.Brand
	err := db.QueryRow("SELECT id, name FROM brands WHERE id = $1", id).Scan(&b.Id, &b.Name)
	if err != nil {
		log.Printf("fetchBrand error: %v", err)
		return views.Brand{}
	}
	return b
}

func fetchCategory(db SqlRepo, id string) views.Category {
	var c views.Category
	err := db.QueryRow("SELECT id, title, uri FROM categories WHERE id = $1", id).Scan(&c.Id, &c.Title, &c.Uri)
	if err != nil {
		log.Printf("fetchCategory error: %v", err)
		return views.Category{}
	}
	return c
}

func fetchCountry(db SqlRepo, id string) views.Country {
	var c views.Country
	err := db.QueryRow("SELECT id, title, friendly FROM countries WHERE id = $1", id).Scan(&c.Id, &c.Title, &c.Friendly)
	if err != nil {
		log.Printf("fetchCountry error: %v", err)
		return views.Country{}
	}
	return c
}

func fetchMaterialsByProductID(db SqlRepo, productID string) []views.Material {
	rows, err := db.Query(`
		SELECT m.id, m.title
		FROM product_materials pm
		JOIN materials m ON m.id = pm.material_id
		WHERE pm.product_id = $1
	`, productID)
	if err != nil {
		log.Printf("fetchMaterials error: %v", err)
		return nil
	}
	defer rows.Close()

	var out []views.Material
	for rows.Next() {
		var m views.Material
		if err := rows.Scan(&m.Id, &m.Title); err != nil {
			log.Printf("fetchMaterials scan error: %v", err)
			continue
		}
		out = append(out, m)
	}
	return out
}

func fetchColorsByProductID(db SqlRepo, productID string) []views.Color {
	rows, err := db.Query(`
		SELECT c.id, c.name, c.hex
		FROM product_colors pc
		JOIN colors c ON c.id = pc.color_id
		WHERE pc.product_id = $1
	`, productID)
	if err != nil {
		log.Printf("fetchColors error: %v", err)
		return nil
	}
	defer rows.Close()

	var out []views.Color
	for rows.Next() {
		var c views.Color
		if err := rows.Scan(&c.Id, &c.Name, &c.Hex); err != nil {
			log.Printf("fetchColors scan error: %v", err)
			continue
		}
		out = append(out, c)
	}
	return out
}

func fetchSimilarProducts(db SqlRepo, productID string) []views.Product {
	const query = `
		SELECT p.id, p.title, p.article, p.brand_id, p.category_id, p.country_id,
			   p.width, p.height, p.depth, p.photos, p.price, p.description
		FROM product_seems ps
		JOIN products p ON p.id = ps.similar_product_id
		WHERE ps.product_id = $1
	`

	rows, err := db.Query(query, productID)
	if err != nil {
		log.Printf("fetchSimilarProducts error: %v", err)
		return nil
	}
	defer rows.Close()

	type rawProduct struct {
		product    views.Product
		brandID    string
		categoryID string
		countryID  string
	}

	var rawList []rawProduct
	for rows.Next() {
		var rp rawProduct
		if err := rows.Scan(&rp.product.Id, &rp.product.Title, &rp.product.Article,
			&rp.brandID, &rp.categoryID, &rp.countryID,
			&rp.product.Width, &rp.product.Height, &rp.product.Depth,
			pq.Array(&rp.product.Photos), &rp.product.Price, &rp.product.Description); err != nil {
			log.Printf("fetchSimilarProducts scan error: %v", err)
			continue
		}
		rawList = append(rawList, rp)
	}

	var result []views.Product
	for _, r := range rawList {
		r.product.Brand = fetchBrand(db, r.brandID)
		r.product.Category = fetchCategory(db, r.categoryID)
		r.product.Country = fetchCountry(db, r.countryID)
		r.product.Materials = fetchMaterialsByProductID(db, r.product.Id)
		r.product.Colors = fetchColorsByProductID(db, r.product.Id)
		r.product.Seems = nil // ⚠️ избегаем рекурсии
		result = append(result, r.product)
	}

	return result
}

func (d Driver) SearchProducts(filter *views.ProductSearch) ([]views.Product, error) {
	const op = "PostgresDb.SearchProducts"

	var (
		query string
		args  []interface{}
	)

	switch {
	case filter.Id != "":
		query = `SELECT id, title, article, brand_id, category_id, country_id,
				 width, height, depth, photos, price, description
				 FROM products WHERE id = $1`
		args = append(args, filter.Id)
	case filter.Article != "":
		query = `SELECT id, title, article, brand_id, category_id, country_id,
				 width, height, depth, photos, price, description
				 FROM products WHERE article = $1`
		args = append(args, filter.Article)
	case filter.Title != "":
		query = `SELECT id, title, article, brand_id, category_id, country_id,
				 width, height, depth, photos, price, description
				 FROM products WHERE title ILIKE $1`
		args = append(args, "%"+filter.Title+"%")
	default:
		return nil, format.Error(op, errors.New("no search parameter provided"))
	}

	rows, err := d.Driver.Query(query, args...)
	if err != nil {
		return nil, format.Error(op, err)
	}
	defer rows.Close()

	var products []views.Product

	for rows.Next() {
		var brandID, categoryID, countryID string
		var p views.Product
		if err := rows.Scan(&p.Id, &p.Title, &p.Article, &brandID, &categoryID, &countryID,
			&p.Width, &p.Height, &p.Depth, pq.Array(&p.Photos), &p.Price, &p.Description); err != nil {
			log.Println(format.Error(op, err))
			continue
		}

		p.Brand = fetchBrand(d.Driver, brandID)
		p.Category = fetchCategory(d.Driver, categoryID)
		p.Country = fetchCountry(d.Driver, countryID)
		p.Materials = fetchMaterialsByProductID(d.Driver, p.Id)
		p.Colors = fetchColorsByProductID(d.Driver, p.Id)
		p.Seems = fetchSimilarProducts(d.Driver, p.Id)

		products = append(products, p)
	}

	return products, nil
}

func (d Driver) GetProductById(id string) (*views.Product, error) {
	const op = "PostgresDb.getProductById"

	const query = `
		SELECT id, title, article, brand_id, category_id, country_id,
		       width, height, depth, photos, price, description
		FROM products
		WHERE id = $1
	`

	var (
		brandID, categoryID, countryID string
		p                              views.Product
	)

	err := d.Driver.QueryRow(query, id).Scan(
		&p.Id,
		&p.Title,
		&p.Article,
		&brandID,
		&categoryID,
		&countryID,
		&p.Width,
		&p.Height,
		&p.Depth,
		pq.Array(&p.Photos),
		&p.Price,
		&p.Description,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("product not find")
		}
		return nil, format.Error(op, err)
	}

	p.Brand = fetchBrand(d.Driver, brandID)
	p.Category = fetchCategory(d.Driver, categoryID)
	p.Country = fetchCountry(d.Driver, countryID)
	p.Materials = fetchMaterialsByProductID(d.Driver, p.Id)
	p.Colors = fetchColorsByProductID(d.Driver, p.Id)
	p.Seems = fetchSimilarProducts(d.Driver, p.Id)

	return &p, nil
}

func (d Driver) GetAllProducts() ([]views.Product, error) {
	const op = "PostgresDb.GetAllProducts"

	rows, err := d.Driver.Query(`
		SELECT id, title, article, brand_id, category_id, country_id,
			   width, height, depth, photos, price, description
		FROM products
	`)
	if err != nil {
		return nil, format.Error(op, err)
	}
	defer rows.Close()

	type rawProduct struct {
		product    views.Product
		brandID    string
		categoryID string
		countryID  string
	}

	var rawList []rawProduct

	for rows.Next() {
		var rp rawProduct
		if err := rows.Scan(&rp.product.Id, &rp.product.Title, &rp.product.Article,
			&rp.brandID, &rp.categoryID, &rp.countryID,
			&rp.product.Width, &rp.product.Height, &rp.product.Depth,
			pq.Array(&rp.product.Photos), &rp.product.Price, &rp.product.Description); err != nil {
			log.Println(format.Error(op, err))
			continue
		}
		rawList = append(rawList, rp)
	}

	var result []views.Product
	for _, r := range rawList {
		r.product.Brand = fetchBrand(d.Driver, r.brandID)
		r.product.Category = fetchCategory(d.Driver, r.categoryID)
		r.product.Country = fetchCountry(d.Driver, r.countryID)
		r.product.Materials = fetchMaterialsByProductID(d.Driver, r.product.Id)
		r.product.Colors = fetchColorsByProductID(d.Driver, r.product.Id)
		r.product.Seems = fetchSimilarProducts(d.Driver, r.product.Id)
		result = append(result, r.product)
	}

	return result, nil
}

// CreateProduct product
func (d Driver) CreateProduct(p *views.ProductId) error {
	const op = "PostgresDb.CreateProduct"

	_, err := d.Driver.Exec(`
		INSERT INTO products (
			id, title, article, brand_id, category_id, country_id, 
			width, height, depth, photos, price, description
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
	`, p.Id, p.Title, p.Article, p.Brand, p.Category, p.Country,
		p.Width, p.Height, p.Depth, pq.Array(p.Photos), p.Price, p.Description)
	if err != nil {
		return format.Error(op, err)
	}

	for _, m := range p.Materials {
		_, err := d.Driver.Exec(`INSERT INTO product_materials (product_id, material_id) VALUES ($1, $2)`, p.Id, m)
		if err != nil {
			return format.Error(op, err)
		}
	}

	for _, c := range p.Colors {
		_, err := d.Driver.Exec(`INSERT INTO product_colors (product_id, color_id) VALUES ($1, $2)`, p.Id, c)
		if err != nil {
			return format.Error(op, err)
		}
	}
	if len(p.Seems) != 0 {
		for _, s := range p.Seems {
			_, err := d.Driver.Exec(`INSERT INTO product_seems (product_id, similar_product_id) VALUES ($1, $2)`, p.Id, s)
			if err != nil {
				return format.Error(op, err)
			}
		}
	}

	return nil
}

// UpdateProduct product
func (d Driver) UpdateProduct(p *views.ProductId, id string) error {
	const op = "PostgresDb.UpdateProduct"

	_, err := d.Driver.Exec(`
		UPDATE products SET 
			title = $2, article = $3, brand_id = $4, category_id = $5,
			country_id = $6, width = $7, height = $8, depth = $9,
			photos = $10, price = $11, description = $12
		WHERE id = $1
	`, id, p.Title, p.Article, p.Brand, p.Category, p.Country,
		p.Width, p.Height, p.Depth, pq.Array(p.Photos), p.Price, p.Description)
	if err != nil {
		return format.Error(op, err)
	}

	// Сначала удаляем старые связи
	_, err = d.Driver.Exec(`DELETE FROM product_materials WHERE product_id = $1`, id)
	if err != nil {
		return format.Error(op, err)
	}
	_, err = d.Driver.Exec(`DELETE FROM product_colors WHERE product_id = $1`, id)
	if err != nil {
		return format.Error(op, err)
	}
	_, err = d.Driver.Exec(`DELETE FROM product_seems WHERE product_id = $1`, id)
	if err != nil {
		return format.Error(op, err)
	}
	// Добавляем новые
	for _, m := range p.Materials {
		_, err := d.Driver.Exec(`INSERT INTO product_materials (product_id, material_id) VALUES ($1, $2)`, id, m)
		if err != nil {
			return format.Error(op, err)
		}
	}

	for _, c := range p.Colors {
		_, err := d.Driver.Exec(`INSERT INTO product_colors (product_id, color_id) VALUES ($1, $2)`, id, c)
		if err != nil {
			return format.Error(op, err)
		}
	}

	for _, s := range p.Seems {
		_, err := d.Driver.Exec(`INSERT INTO product_seems (product_id, similar_product_id) VALUES ($1, $2)`, id, s)
		if err != nil {
			return format.Error(op, err)
		}
	}

	return nil
}

// DeleteProduct product
func (d Driver) DeleteProduct(id string) error {
	const op = "PostgresDb.DeleteProduct"
	_, err := d.Driver.Exec(`DELETE FROM products WHERE id = $1`, id)
	return format.Error(op, err)
}

func (d Driver) FilterProducts(filter *views.ProductFilter) ([]views.Product, error) {
	const op = "PostgresDb.FilterProducts"

	var conditions []string
	var args []interface{}
	argPos := 1

	buildFilter := func(field string, values []string) {
		if len(values) == 0 {
			return
		}
		placeholders := make([]string, len(values))
		for i, v := range values {
			placeholders[i] = fmt.Sprintf("$%d", argPos)
			args = append(args, v)
			argPos++
		}
		conditions = append(conditions, fmt.Sprintf("%s IN (%s)", field, strings.Join(placeholders, ",")))
	}

	buildFilter("brand_id", filter.Brand)
	buildFilter("category_id", filter.Category)
	buildFilter("country_id", filter.Country)

	// Числовые фильтры
	numericFilters := []struct {
		field string
		value int
		op    string
	}{
		{"width", filter.MinWidth, ">="},
		{"width", filter.MaxWidth, "<="},
		{"height", filter.MinHeight, ">="},
		{"height", filter.MaxHeight, "<="},
		{"depth", filter.MinDepth, ">="},
		{"depth", filter.MaxDepth, "<="},
		{"price", filter.MinPrice, ">="},
		{"price", filter.MaxPrice, "<="},
	}

	for _, f := range numericFilters {
		if f.value > 0 {
			conditions = append(conditions, fmt.Sprintf("%s %s $%d", f.field, f.op, argPos))
			args = append(args, f.value)
			argPos++
		}
	}

	if len(filter.Materials) > 0 {
		placeholders := make([]string, len(filter.Materials))
		for i, v := range filter.Materials {
			placeholders[i] = fmt.Sprintf("$%d", argPos)
			args = append(args, v)
			argPos++
		}
		conditions = append(conditions, fmt.Sprintf(`
			EXISTS (
				SELECT 1 FROM product_materials pm
				WHERE pm.product_id = products.id AND pm.material_id IN (%s)
			)`, strings.Join(placeholders, ",")))
	}

	if len(filter.Colors) > 0 {
		placeholders := make([]string, len(filter.Colors))
		for i, v := range filter.Colors {
			placeholders[i] = fmt.Sprintf("$%d", argPos)
			args = append(args, v)
			argPos++
		}
		conditions = append(conditions, fmt.Sprintf(`
			EXISTS (
				SELECT 1 FROM product_colors pc
				WHERE pc.product_id = products.id AND pc.color_id IN (%s)
			)`, strings.Join(placeholders, ",")))
	}

	query := `
		SELECT id, title, article, brand_id, category_id, country_id,
			   width, height, depth, photos, price, description
		FROM products
	`
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	// Сортировка
	if filter.SortBy != "" {
		validSort := map[string]bool{"price": true, "width": true, "height": true, "depth": true, "title": true}
		if validSort[filter.SortBy] {
			order := "ASC"
			if strings.ToLower(filter.SortOrder) == "desc" {
				order = "DESC"
			}
			query += fmt.Sprintf(" ORDER BY %s %s", filter.SortBy, order)
		}
	}

	// Пагинация
	if filter.Limit > 0 {
		query += fmt.Sprintf(" LIMIT $%d", argPos)
		args = append(args, filter.Limit)
		argPos++
	}
	if filter.Offset > 0 {
		query += fmt.Sprintf(" OFFSET $%d", argPos)
		args = append(args, filter.Offset)
	}

	// Выполнение запроса
	rows, err := d.Driver.Query(query, args...)
	if err != nil {
		return nil, format.Error(op, err)
	}
	defer rows.Close()

	type rawProduct struct {
		product    views.Product
		brandID    string
		categoryID string
		countryID  string
	}

	var rawList []rawProduct

	for rows.Next() {
		var rp rawProduct
		if err := rows.Scan(&rp.product.Id, &rp.product.Title, &rp.product.Article,
			&rp.brandID, &rp.categoryID, &rp.countryID,
			&rp.product.Width, &rp.product.Height, &rp.product.Depth,
			pq.Array(&rp.product.Photos), &rp.product.Price, &rp.product.Description); err != nil {
			log.Println(format.Error(op, err))
			continue
		}
		rawList = append(rawList, rp)
	}

	var result []views.Product
	for _, r := range rawList {
		r.product.Brand = fetchBrand(d.Driver, r.brandID)
		r.product.Category = fetchCategory(d.Driver, r.categoryID)
		r.product.Country = fetchCountry(d.Driver, r.countryID)
		r.product.Materials = fetchMaterialsByProductID(d.Driver, r.product.Id)
		r.product.Colors = fetchColorsByProductID(d.Driver, r.product.Id)
		r.product.Seems = nil
		result = append(result, r.product)
	}

	return result, nil
}
