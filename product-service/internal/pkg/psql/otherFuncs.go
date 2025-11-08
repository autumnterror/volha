package psql

import (
	"log"
	"productService/internal/utils/format"
	"productService/internal/views"
	"strconv"
	"strings"
)

func (d Driver) GetDictionariesByCategory(id string) (*views.Dictionaries, error) {
	const op = "PostgresDb.GetDictionaries"

	query := `
		WITH dicts AS (
			SELECT 'brand' as type, id, name, '' as extra1, '' as extra2 FROM brands
			UNION ALL
			SELECT 'category', id, title, uri, '' FROM categories
			UNION ALL
			SELECT 'country', id, title, friendly, '' FROM countries
			UNION ALL
			SELECT 'material', id, title, '', '' FROM materials
			UNION ALL
			SELECT 'color', id, name, hex, '' FROM colors
		),
		stats AS (
			SELECT 
				MIN(price)::text AS min_price,
				MAX(price)::text AS max_price,
				MIN(width)::text AS min_width,
				MAX(width)::text AS max_width,
				MIN(height)::text AS min_height,
				MAX(height)::text AS max_height,
				MIN(depth)::text AS min_depth,
				MAX(depth)::text AS max_depth
			FROM products
			WHERE category_id = $1
		)
		SELECT * FROM dicts
		UNION ALL
		SELECT 'stats', '', min_price, max_price, min_width || ',' || max_width || ',' || min_height || ',' || max_height || ',' || min_depth || ',' || max_depth FROM stats;
	`

	rows, err := d.Driver.Query(query, id)
	if err != nil {
		return nil, format.Error(op, err)
	}
	defer rows.Close()

	result := &views.Dictionaries{}
	for rows.Next() {
		var typ, id, field1, field2, field3 string
		if err := rows.Scan(&typ, &id, &field1, &field2, &field3); err != nil {
			log.Println(format.Error(op, err))
			continue
		}

		switch typ {
		case "brand":
			result.Brands = append(result.Brands, views.Brand{Id: id, Name: field1})
		case "category":
			result.Categories = append(result.Categories, views.Category{Id: id, Title: field1, Uri: field2})
		case "country":
			result.Countries = append(result.Countries, views.Country{Id: id, Title: field1, Friendly: field2})
		case "material":
			result.Materials = append(result.Materials, views.Material{Id: id, Title: field1})
		case "color":
			result.Colors = append(result.Colors, views.Color{Id: id, Name: field1, Hex: field2})
		case "stats":
			// field1 = min_price, field2 = max_price, field3 = "minW,maxW,minH,maxH,minD,maxD"
			if minPrice, err := strconv.Atoi(field1); err == nil {
				result.MinPrice = minPrice
			}
			if maxPrice, err := strconv.Atoi(field2); err == nil {
				result.MaxPrice = maxPrice
			}
			parts := strings.Split(field3, ",")
			if len(parts) == 6 {
				result.MinWidth, _ = strconv.Atoi(parts[0])
				result.MaxWidth, _ = strconv.Atoi(parts[1])
				result.MinHeight, _ = strconv.Atoi(parts[2])
				result.MaxHeight, _ = strconv.Atoi(parts[3])
				result.MinDepth, _ = strconv.Atoi(parts[4])
				result.MaxDepth, _ = strconv.Atoi(parts[5])
			}
		}
	}

	return result, nil
}

func (d Driver) GetDictionaries() (*views.Dictionaries, error) {
	const op = "PostgresDb.GetDictionaries"

	query := `
		WITH dicts AS (
			SELECT 'brand' as type, id, name, '' as extra1, '' as extra2 FROM brands
			UNION ALL
			SELECT 'category', id, title, uri, '' FROM categories
			UNION ALL
			SELECT 'country', id, title, friendly, '' FROM countries
			UNION ALL
			SELECT 'material', id, title, '', '' FROM materials
			UNION ALL
			SELECT 'color', id, name, hex, '' FROM colors
		),
		stats AS (
			SELECT 
				MIN(price)::text AS min_price,
				MAX(price)::text AS max_price,
				MIN(width)::text AS min_width,
				MAX(width)::text AS max_width,
				MIN(height)::text AS min_height,
				MAX(height)::text AS max_height,
				MIN(depth)::text AS min_depth,
				MAX(depth)::text AS max_depth
			FROM products
		)
		SELECT * FROM dicts
		UNION ALL
		SELECT 'stats', '', min_price, max_price, min_width || ',' || max_width || ',' || min_height || ',' || max_height || ',' || min_depth || ',' || max_depth FROM stats;
	`

	rows, err := d.Driver.Query(query)
	if err != nil {
		return nil, format.Error(op, err)
	}
	defer rows.Close()

	result := &views.Dictionaries{}
	for rows.Next() {
		var typ, id, field1, field2, field3 string
		if err := rows.Scan(&typ, &id, &field1, &field2, &field3); err != nil {
			log.Println(format.Error(op, err))
			continue
		}

		switch typ {
		case "brand":
			result.Brands = append(result.Brands, views.Brand{Id: id, Name: field1})
		case "category":
			result.Categories = append(result.Categories, views.Category{Id: id, Title: field1, Uri: field2})
		case "country":
			result.Countries = append(result.Countries, views.Country{Id: id, Title: field1, Friendly: field2})
		case "material":
			result.Materials = append(result.Materials, views.Material{Id: id, Title: field1})
		case "color":
			result.Colors = append(result.Colors, views.Color{Id: id, Name: field1, Hex: field2})
		case "stats":
			// field1 = min_price, field2 = max_price, field3 = "minW,maxW,minH,maxH,minD,maxD"
			if minPrice, err := strconv.Atoi(field1); err == nil {
				result.MinPrice = minPrice
			}
			if maxPrice, err := strconv.Atoi(field2); err == nil {
				result.MaxPrice = maxPrice
			}
			parts := strings.Split(field3, ",")
			if len(parts) == 6 {
				result.MinWidth, _ = strconv.Atoi(parts[0])
				result.MaxWidth, _ = strconv.Atoi(parts[1])
				result.MinHeight, _ = strconv.Atoi(parts[2])
				result.MaxHeight, _ = strconv.Atoi(parts[3])
				result.MinDepth, _ = strconv.Atoi(parts[4])
				result.MaxDepth, _ = strconv.Atoi(parts[5])
			}
		}
	}

	return result, nil
}
