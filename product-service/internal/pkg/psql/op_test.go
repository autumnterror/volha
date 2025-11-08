package psql

import (
	"github.com/stretchr/testify/assert"
	"log"
	"productService/config"
	"productService/internal/views"
	"testing"
)

//func TestCRUDProductColorPhotos(t *testing.T) {
//	t.Parallel()
//	db, err := NewConnect(config.Test())
//	assert.NoError(t, err)
//
//	tx, err := db.Driver.Begin()
//	assert.NoError(t, err)
//
//	driver := Driver{Driver: tx}
//
//	t.Cleanup(func() {
//		assert.NoError(t, tx.Rollback())
//		assert.NoError(t, db.Disconnect())
//	})
//
//	// Подготовка связанных сущностей
//	brand := &views.Brand{Id: "brand2", Name: "TestBrand2"}
//	category := &views.Category{Id: "cat2", Title: "TestCategory2", Uri: "test-cat-2"}
//	country := &views.Country{Id: "country2", Title: "TestCountry2", Friendly: "Friendly"}
//	color := &views.Color{Id: "color2", Name: "Blue", Hex: "#0000FF"}
//
//	assert.NoError(t, driver.CreateBrand(brand))
//	assert.NoError(t, driver.CreateCategory(category))
//	assert.NoError(t, driver.CreateCountry(country))
//	assert.NoError(t, driver.CreateColor(color))
//
//	product := &views.ProductId{
//		Id:          "prod2",
//		Title:       "Test Product",
//		Article:     "ART-002",
//		Brand:       brand.Id,
//		Category:    category.Id,
//		Country:     country.Id,
//		Width:       40,
//		Height:      90,
//		Depth:       25,
//		Materials:   []string{},
//		Colors:      []string{color.Id},
//		Photos:      []string{"common1.jpg"},
//		Seems:       []string{},
//		Price:       1500,
//		Description: "Product with color photos",
//	}
//	assert.NoError(t, driver.CreateProduct(product))
//
//	// Создать запись
//	pcp := &views.ProductColorPhotos{
//		ProductId: product.Id,
//		ColorId:   color.Id,
//		Photos:    []string{"blue1.jpg", "blue2.jpg"},
//	}
//	assert.NoError(t, driver.CreateProductColorPhotos(pcp))
//
//	// Получить все
//	all, err := driver.GetAllProductColorPhotos()
//	assert.NoError(t, err)
//	assert.NotEmpty(t, all)
//	log.Println("all product_color_photos:", all)
//
//	// Получить по product_id + color_id
//	photos, err := driver.GetPhotosByProductAndColor(product.Id, color.Id)
//	assert.NoError(t, err)
//	assert.Equal(t, []string{"blue1.jpg", "blue2.jpg"}, photos)
//
//	// Обновить запись
//	pcp.Photos = []string{"blue3.jpg", "blue4.jpg"}
//	assert.NoError(t, driver.UpdateProductColorPhotos(pcp))
//
//	photos, err = driver.GetPhotosByProductAndColor(product.Id, color.Id)
//	assert.NoError(t, err)
//	assert.Equal(t, []string{"blue3.jpg", "blue4.jpg"}, photos)
//
//	// Удалить запись
//	assert.NoError(t, driver.DeleteProductColorPhotos(product.Id, color.Id))
//
//	_, err = driver.GetPhotosByProductAndColor(product.Id, color.Id)
//	assert.Error(t, err) // должна быть ошибка, записи больше нет
//}

//func TestFilterProducts(t *testing.T) {
//	t.Parallel()
//	db, err := NewConnect(config.Test())
//	assert.NoError(t, err)
//
//	tx, err := db.Driver.Begin()
//	assert.NoError(t, err)
//
//	driver := Driver{Driver: tx}
//
//	t.Cleanup(func() {
//		assert.NoError(t, tx.Rollback())
//		assert.NoError(t, db.Disconnect())
//	})
//
//	// Подготовка справочников
//	brands := []views.Brand{
//		{Id: "brandA", Name: "Brand A"},
//		{Id: "brandB", Name: "Brand B"},
//	}
//	categories := []views.Category{
//		{Id: "catA", Title: "Cat A", Uri: "cat-a"},
//	}
//	countries := []views.Country{
//		{Id: "countryX", Title: "Country X", Friendly: "X"},
//	}
//	materials := []views.Material{
//		{Id: "mat1", Title: "Wood"},
//		{Id: "mat2", Title: "Metal"},
//	}
//	colors := []views.Color{
//		{Id: "color1", Name: "Black", Hex: "#000000"},
//		{Id: "color2", Name: "White", Hex: "#FFFFFF"},
//	}
//
//	for _, b := range brands {
//		assert.NoError(t, driver.CreateBrand(&b))
//	}
//	for _, c := range categories {
//		assert.NoError(t, driver.CreateCategory(&c))
//	}
//	for _, c := range countries {
//		assert.NoError(t, driver.CreateCountry(&c))
//	}
//	for _, m := range materials {
//		assert.NoError(t, driver.CreateMaterial(&m))
//	}
//	for _, c := range colors {
//		assert.NoError(t, driver.CreateColor(&c))
//	}
//
//	// Продукты
//	products := []views.ProductId{
//		{
//			Id:          "p1",
//			Title:       "Wooden Table",
//			Article:     "ART001",
//			Brand:       "brandA",
//			Category:    "catA",
//			Country:     "countryX",
//			Width:       100,
//			Height:      75,
//			Depth:       60,
//			Materials:   []string{"mat1"},
//			Colors:      []string{"color1"},
//			Photos:      []string{"wood1.jpg"},
//			Price:       500,
//			Description: "Made of wood",
//		},
//		{
//			Id:          "p2",
//			Title:       "Metal Chair",
//			Article:     "ART002",
//			Brand:       "brandB",
//			Category:    "catA",
//			Country:     "countryX",
//			Width:       45,
//			Height:      90,
//			Depth:       50,
//			Materials:   []string{"mat2"},
//			Colors:      []string{"color2"},
//			Photos:      []string{"metal1.jpg"},
//			Price:       200,
//			Description: "Made of metal",
//		},
//		{
//			Id:          "p3",
//			Title:       "Hybrid Bench",
//			Article:     "ART003",
//			Brand:       "brandA",
//			Category:    "catA",
//			Country:     "countryX",
//			Width:       120,
//			Height:      85,
//			Depth:       60,
//			Materials:   []string{"mat1", "mat2"},
//			Colors:      []string{"color1", "color2"},
//			Photos:      []string{"bench.jpg"},
//			Price:       800,
//			Description: "Made of wood and metal",
//		},
//	}
//
//	for _, p := range products {
//		assert.NoError(t, driver.CreateProduct(&p))
//	}
//	extractIds := func(products []views.Product) []string {
//		var ids []string
//		for _, p := range products {
//			ids = append(ids, p.Id)
//		}
//		return ids
//	}
//
//	t.Run("Filter by Category", func(t *testing.T) {
//		filter := &views.ProductFilter{Category: []string{"catA"}}
//		t.Log(filter)
//		found, err := driver.FilterProducts(filter)
//		t.Log(found)
//		assert.NoError(t, err)
//		assert.GreaterOrEqual(t, len(found), 3) // в тестовых точно 3
//	})
//
//	t.Run("Filter by Non-existent Brand", func(t *testing.T) {
//		filter := &views.ProductFilter{Brand: []string{"brandX"}}
//		t.Log(filter)
//		found, err := driver.FilterProducts(filter)
//		t.Log(found)
//		assert.NoError(t, err)
//		assert.Len(t, found, 0)
//	})
//
//	t.Run("Filter by Width Range", func(t *testing.T) {
//		filter := &views.ProductFilter{MinWidth: 100, MaxWidth: 120}
//		t.Log(filter)
//		found, err := driver.FilterProducts(filter)
//		t.Log(found)
//		assert.NoError(t, err)
//		assert.Subset(t, extractIds(found), []string{"p1", "p3"})
//	})
//
//	t.Run("Filter by Height Minimum", func(t *testing.T) {
//		filter := &views.ProductFilter{MinHeight: 85}
//		t.Log(filter)
//		found, err := driver.FilterProducts(filter)
//		t.Log(found)
//		assert.NoError(t, err)
//		assert.Subset(t, extractIds(found), []string{"p2", "p3"})
//	})
//
//	t.Run("Filter by Depth Max", func(t *testing.T) {
//		filter := &views.ProductFilter{MaxDepth: 50}
//		t.Log(filter)
//		found, err := driver.FilterProducts(filter)
//		t.Log(found)
//		assert.NoError(t, err)
//		assert.Subset(t, extractIds(found), []string{"p2"})
//	})
//
//	t.Run("Filter by Material and Brand", func(t *testing.T) {
//		filter := &views.ProductFilter{
//			Materials: []string{"mat1"},
//			Brand:     []string{"brandA"},
//		}
//		t.Log(filter)
//		found, err := driver.FilterProducts(filter)
//		t.Log(found)
//		assert.NoError(t, err)
//		assert.ElementsMatch(t, []string{"p1", "p3"}, extractIds(found))
//	})
//
//	t.Run("Filter by Multiple Colors", func(t *testing.T) {
//		filter := &views.ProductFilter{Colors: []string{"color1", "color2"}}
//		t.Log(filter)
//		found, err := driver.FilterProducts(filter)
//		t.Log(found)
//		assert.NoError(t, err)
//		assert.ElementsMatch(t, []string{"p1", "p2", "p3"}, extractIds(found))
//	})
//
//	t.Run("Filter with Limit Only", func(t *testing.T) {
//		filter := &views.ProductFilter{Limit: 1}
//		t.Log(filter)
//		found, err := driver.FilterProducts(filter)
//		t.Log(found)
//		assert.NoError(t, err)
//		assert.Len(t, found, 1)
//	})
//
//	t.Run("Filter with Offset Only", func(t *testing.T) {
//		filter := &views.ProductFilter{Offset: 1}
//		t.Log(filter)
//		found, err := driver.FilterProducts(filter)
//		t.Log(found)
//		assert.NoError(t, err)
//		assert.GreaterOrEqual(t, len(found), 2)
//	})
//
//	t.Run("Filter with All Fields", func(t *testing.T) {
//		filter := &views.ProductFilter{
//			Brand:     []string{"brandA"},
//			Category:  []string{"catA"},
//			Country:   []string{"countryX"},
//			MinWidth:  100,
//			MaxWidth:  120,
//			Materials: []string{"mat1"},
//			Colors:    []string{"color1"},
//			MinPrice:  400,
//			MaxPrice:  900,
//		}
//		t.Log(filter)
//		found, err := driver.FilterProducts(filter)
//		t.Log(found)
//		assert.NoError(t, err)
//		assert.ElementsMatch(t, []string{"p1", "p3"}, extractIds(found))
//	})
//
//	t.Run("Filter Invalid SortBy (ignored)", func(t *testing.T) {
//		filter := &views.ProductFilter{SortBy: "invalid_field"}
//		t.Log(filter)
//		found, err := driver.FilterProducts(filter)
//		t.Log(found)
//		assert.NoError(t, err)
//		assert.GreaterOrEqual(t, len(found), 3)
//	})
//
//	t.Run("Filter with Empty Filter (all products)", func(t *testing.T) {
//		filter := &views.ProductFilter{}
//		t.Log(filter)
//		found, err := driver.FilterProducts(filter)
//		t.Log(found)
//		assert.NoError(t, err)
//		assert.GreaterOrEqual(t, len(found), 3) // хотя бы тестовые 3
//	})
//
//	t.Run("Filter with Non-existent Brand", func(t *testing.T) {
//		filter := &views.ProductFilter{Brand: []string{"nonexistent_brand"}}
//		t.Log(filter)
//		found, err := driver.FilterProducts(filter)
//		t.Log(found)
//		assert.NoError(t, err)
//		assert.Len(t, found, 0)
//	})
//
//	t.Run("Filter with Multiple Conditions", func(t *testing.T) {
//		filter := &views.ProductFilter{
//			Brand:    []string{"brandA"},
//			Colors:   []string{"color1"},
//			MinPrice: 400,
//			MaxPrice: 900,
//		}
//		t.Log(filter)
//		found, err := driver.FilterProducts(filter)
//		t.Log(found)
//		assert.NoError(t, err)
//		assert.Subset(t, extractIds(found), []string{"p1", "p3"})
//	})
//
//	t.Run("Sort by Title ASC", func(t *testing.T) {
//		filter := &views.ProductFilter{
//			SortBy:    "title",
//			SortOrder: "asc",
//		}
//		t.Log(filter)
//		found, err := driver.FilterProducts(filter)
//		t.Log(found)
//		assert.NoError(t, err)
//		assert.True(t, len(found) >= 2)
//		for i := 1; i < len(found); i++ {
//			assert.LessOrEqual(t, found[i-1].Title, found[i].Title)
//		}
//	})
//
//	t.Run("Sort by Height DESC", func(t *testing.T) {
//		filter := &views.ProductFilter{
//			SortBy:    "height",
//			SortOrder: "desc",
//		}
//		t.Log(filter)
//		found, err := driver.FilterProducts(filter)
//		t.Log(found)
//		assert.NoError(t, err)
//		assert.True(t, len(found) >= 2)
//		for i := 1; i < len(found); i++ {
//			assert.GreaterOrEqual(t, found[i-1].Height, found[i].Height)
//		}
//	})
//
//	t.Run("Pagination - Limit Only", func(t *testing.T) {
//		filter := &views.ProductFilter{
//			SortBy: "price", SortOrder: "asc", Limit: 1,
//		}
//		t.Log(filter)
//		found, err := driver.FilterProducts(filter)
//		t.Log(found)
//		assert.NoError(t, err)
//		assert.Len(t, found, 1)
//	})
//
//	t.Run("Pagination - Offset Only", func(t *testing.T) {
//		filter := &views.ProductFilter{
//			SortBy: "price", SortOrder: "asc", Offset: 1,
//		}
//		t.Log(filter)
//		found, err := driver.FilterProducts(filter)
//		t.Log(found)
//		assert.NoError(t, err)
//		assert.True(t, len(found) >= 1)
//	})
//
//	t.Run("Filter by MaxWidth Only", func(t *testing.T) {
//		filter := &views.ProductFilter{MaxWidth: 50}
//		t.Log(filter)
//		found, err := driver.FilterProducts(filter)
//		t.Log(found)
//		assert.NoError(t, err)
//		for _, p := range found {
//			assert.LessOrEqual(t, p.Width, 50)
//		}
//	})
//
//	t.Run("Filter by Multiple Materials", func(t *testing.T) {
//		filter := &views.ProductFilter{Materials: []string{"mat1", "mat2"}}
//		t.Log(filter)
//		found, err := driver.FilterProducts(filter)
//		t.Log(found)
//		assert.NoError(t, err)
//		assert.Subset(t, extractIds(found), []string{"p1", "p2", "p3"})
//	})
//
//}

//func TestClean(t *testing.T) {
//	db, err := NewConnect(config.Test())
//	assert.NoError(t, err)
//
//	driver := Driver{Driver: db.Driver}
//	defer db.Disconnect()
//
//	cl, err := driver.GetAllColors()
//	assert.NoError(t, err)
//	for _, i := range cl {
//		assert.NoError(t, driver.DeleteColor(i.Id))
//	}
//
//	b, err := driver.GetAllBrands()
//	assert.NoError(t, err)
//	for _, i := range b {
//		assert.NoError(t, driver.DeleteBrand(i.Id))
//	}
//
//	ct, err := driver.GetAllCategories()
//	assert.NoError(t, err)
//	for _, i := range ct {
//		assert.NoError(t, driver.DeleteCategory(i.Id))
//	}
//
//	cn, err := driver.GetAllCountries()
//	assert.NoError(t, err)
//	for _, i := range cn {
//		assert.NoError(t, driver.DeleteCountry(i.Id))
//	}
//
//	m, err := driver.GetAllMaterials()
//	assert.NoError(t, err)
//	for _, i := range m {
//		assert.NoError(t, driver.DeleteMaterial(i.Id))
//	}
//	p, err := driver.GetAllProducts()
//	assert.NoError(t, err)
//	for _, i := range p {
//		assert.NoError(t, driver.DeleteProduct(i.Id))
//	}
//}

//func TestGetDictionaries(t *testing.T) {
//	db, err := NewConnect(config.Test())
//	assert.NoError(t, err)
//
//	tx, err := db.Driver.Begin()
//	assert.NoError(t, err)
//
//	driver := Driver{Driver: tx}
//	defer db.Disconnect()
//	t.Cleanup(func() { _ = tx.Rollback() })
//	brand := &views.Brand{Id: "brand_test", Name: "Test Brand"}
//	cat := &views.Category{Id: "cat_test", Title: "Test Cat", Uri: "test-uri"}
//	cou := &views.Country{Id: "country_test", Title: "Test Country", Friendly: "FriendlyName"}
//	m := &views.Material{Id: "mat_test", Title: "Test Material"}
//	col := &views.Color{Id: "color_test", Name: "TestColor", Hex: "#123456"}
//
//	err = driver.CreateBrand(brand)
//	assert.NoError(t, err)
//	err = driver.CreateCategory(cat)
//	assert.NoError(t, err)
//	err = driver.CreateCountry(cou)
//	assert.NoError(t, err)
//	err = driver.CreateMaterial(m)
//	assert.NoError(t, err)
//	err = driver.CreateColor(col)
//	assert.NoError(t, err)
//
//	d, err := driver.GetDictionaries()
//	assert.NoError(t, err)
//	log.Println(d)
//}

//func TestCRUDSimpleTables(t *testing.T) {
//	t.Parallel()
//	t.Run("brands", func(t *testing.T) {
//		t.Parallel()
//
//		db, err := NewConnect(config.Test())
//		assert.NoError(t, err)
//
//		tx, err := db.Driver.Begin()
//		assert.NoError(t, err)
//
//		driver := Driver{Driver: tx}
//		defer db.Disconnect()
//		t.Cleanup(func() { _ = tx.Rollback() })
//
//		brand := &views.Brand{Id: "brand_test", Name: "Test Brand"}
//		updated := &views.Brand{Id: "brand_test", Name: "Updated Brand"}
//
//		assert.NoError(t, driver.CreateBrand(brand))
//
//		all, err := driver.GetAllBrands()
//		assert.NoError(t, err)
//		log.Println("brands after create:", all)
//
//		assert.NoError(t, driver.UpdateBrand(updated, brand.Id))
//
//		all, err = driver.GetAllBrands()
//		assert.NoError(t, err)
//		log.Println("brands after update:", all)
//
//		assert.NoError(t, driver.DeleteBrand(brand.Id))
//	})
//
//	t.Run("categories", func(t *testing.T) {
//		t.Parallel()
//
//		db, err := NewConnect(config.Test())
//		assert.NoError(t, err)
//
//		tx, err := db.Driver.Begin()
//		assert.NoError(t, err)
//
//		driver := Driver{Driver: tx}
//		defer db.Disconnect()
//		t.Cleanup(func() { _ = tx.Rollback() })
//
//		c := &views.Category{Id: "cat_test", Title: "Test Cat", Uri: "test-uri"}
//		updated := &views.Category{Id: "cat_test", Title: "Updated", Uri: "upd-uri"}
//
//		assert.NoError(t, driver.CreateCategory(c))
//
//		all, err := driver.GetAllCategories()
//		assert.NoError(t, err)
//		log.Println("categories after create:", all)
//
//		assert.NoError(t, driver.UpdateCategory(updated, c.Id))
//
//		all, err = driver.GetAllCategories()
//		assert.NoError(t, err)
//		log.Println("categories after update:", all)
//
//		assert.NoError(t, driver.DeleteCategory(c.Id))
//	})
//
//	t.Run("countries", func(t *testing.T) {
//		t.Parallel()
//
//		db, err := NewConnect(config.Test())
//		assert.NoError(t, err)
//
//		tx, err := db.Driver.Begin()
//		assert.NoError(t, err)
//
//		driver := Driver{Driver: tx}
//		defer db.Disconnect()
//		t.Cleanup(func() { _ = tx.Rollback() })
//
//		c := &views.Country{Id: "country_test", Title: "Test Country", Friendly: "FriendlyName"}
//		updated := &views.Country{Id: "country_test", Title: "Updated", Friendly: "UpdatedFriendly"}
//
//		assert.NoError(t, driver.CreateCountry(c))
//
//		all, err := driver.GetAllCountries()
//		assert.NoError(t, err)
//		log.Println("countries after create:", all)
//
//		assert.NoError(t, driver.UpdateCountry(updated, c.Id))
//
//		all, err = driver.GetAllCountries()
//		assert.NoError(t, err)
//		log.Println("countries after update:", all)
//
//		assert.NoError(t, driver.DeleteCountry(c.Id))
//	})
//
//	t.Run("materials", func(t *testing.T) {
//		t.Parallel()
//
//		db, err := NewConnect(config.Test())
//		assert.NoError(t, err)
//
//		tx, err := db.Driver.Begin()
//		assert.NoError(t, err)
//
//		driver := Driver{Driver: tx}
//		defer db.Disconnect()
//		t.Cleanup(func() { _ = tx.Rollback() })
//
//		m := &views.Material{Id: "mat_test", Title: "Test Material"}
//		updated := &views.Material{Id: "mat_test", Title: "Updated Material"}
//
//		assert.NoError(t, driver.CreateMaterial(m))
//
//		all, err := driver.GetAllMaterials()
//		assert.NoError(t, err)
//		log.Println("materials after create:", all)
//
//		assert.NoError(t, driver.UpdateMaterial(updated, m.Id))
//
//		all, err = driver.GetAllMaterials()
//		assert.NoError(t, err)
//		log.Println("materials after update:", all)
//
//		assert.NoError(t, driver.DeleteMaterial(m.Id))
//	})
//
//	t.Run("colors", func(t *testing.T) {
//		t.Parallel()
//
//		db, err := NewConnect(config.Test())
//		assert.NoError(t, err)
//
//		tx, err := db.Driver.Begin()
//		assert.NoError(t, err)
//
//		driver := Driver{Driver: tx}
//		defer db.Disconnect()
//		t.Cleanup(func() { _ = tx.Rollback() })
//
//		c := &views.Color{Id: "color_test", Name: "TestColor", Hex: "#123456"}
//		updated := &views.Color{Id: "color_test", Name: "UpdatedColor", Hex: "#654321"}
//
//		assert.NoError(t, driver.CreateColor(c))
//
//		all, err := driver.GetAllColors()
//		assert.NoError(t, err)
//		log.Println("colors after create:", all)
//
//		assert.NoError(t, driver.UpdateColor(updated, c.Id))
//
//		all, err = driver.GetAllColors()
//		assert.NoError(t, err)
//		log.Println("colors after update:", all)
//
//		assert.NoError(t, driver.DeleteColor(c.Id))
//	})
//}

func TestCRUDProducts(t *testing.T) {
	t.Parallel()
	db, err := NewConnect(config.Test())
	assert.NoError(t, err)

	tx, err := db.Driver.Begin()
	assert.NoError(t, err)

	//tx := db.Driver

	driver := Driver{Driver: tx}

	t.Cleanup(func() {
		assert.NoError(t, tx.Rollback())
		assert.NoError(t, db.Disconnect())
	})

	// Подготовка связанных сущностей
	brand := &views.Brand{Id: "brand1", Name: "TestBrand"}
	category := &views.Category{Id: "cat1", Title: "TestCategory", Uri: "test-cat"}
	country := &views.Country{Id: "country1", Title: "TestCountry", Friendly: "Friendly"}
	material := &views.Material{Id: "mat1", Title: "Leather"}
	color := &views.Color{Id: "color1", Name: "Red", Hex: "#FF0000"}

	assert.NoError(t, driver.CreateBrand(brand))
	assert.NoError(t, driver.CreateCategory(category))
	assert.NoError(t, driver.CreateCountry(country))
	assert.NoError(t, driver.CreateMaterial(material))
	assert.NoError(t, driver.CreateColor(color))
	cl, err := driver.GetAllColors()
	assert.NoError(t, err)
	log.Println(cl)

	b, err := driver.GetAllBrands()
	assert.NoError(t, err)
	log.Println(b)

	ct, err := driver.GetAllCategories()
	assert.NoError(t, err)
	log.Println(ct)

	cn, err := driver.GetAllCountries()
	assert.NoError(t, err)
	log.Println(cn)

	m, err := driver.GetAllMaterials()
	assert.NoError(t, err)
	log.Println(m)

	product := &views.ProductId{
		Id:          "prod1",
		Title:       "Test ProductId",
		Article:     "ART-001",
		Brand:       brand.Id,
		Category:    category.Id,
		Country:     country.Id,
		Width:       50,
		Height:      100,
		Depth:       30,
		Materials:   []string{material.Id},
		Colors:      []string{color.Id},
		Photos:      []string{"img1.jpg", "img2.jpg"},
		Seems:       []string{},
		Price:       999,
		Description: "Test description",
	}

	assert.NoError(t, driver.CreateProduct(product))

	products, err := driver.GetAllProducts()
	assert.NoError(t, err)
	assert.NotEmpty(t, products)
	log.Println("products after create:", products, "\n\n", "")

	pr, err := driver.GetProductById("prod1")
	assert.NoError(t, err)
	assert.NotEmpty(t, pr)
	log.Println("Get product by id: ", pr)

	product.Title = "Updated ProductId"
	product.Width = 80
	assert.NoError(t, driver.UpdateProduct(product, product.Id))

	products, err = driver.GetAllProducts()
	assert.NoError(t, err)
	log.Println("products after update:", products, "\n\n", "")

	assert.NoError(t, driver.DeleteProduct(product.Id))

	products, err = driver.GetAllProducts()
	assert.NoError(t, err)
	log.Println("products after delete:", products, "\n\n", "")
}

//
//func TestProductFilters(t *testing.T) {
//	t.Parallel()
//	db, err := NewConnect(config.Test())
//	assert.NoError(t, err)
//
//	tx, err := db.Driver.Begin()
//	assert.NoError(t, err)
//	//tx := db.Driver
//
//	driver := Driver{Driver: tx}
//	defer db.Disconnect()
//	t.Cleanup(func() { _ = tx.Rollback() })
//
//	// Подготовка справочников
//	brand := &views.Brand{Id: "brand1", Name: "Brand"}
//	category := &views.Category{Id: "cat1", Title: "Cat", Uri: "cat"}
//	country := &views.Country{Id: "country1", Title: "Country", Friendly: "Nice"}
//	material := &views.Material{Id: "mat1", Title: "Mat"}
//	color := &views.Color{Id: "color1", Name: "Green", Hex: "#00FF00"}
//
//	assert.NoError(t, driver.CreateBrand(brand))
//	assert.NoError(t, driver.CreateCategory(category))
//	assert.NoError(t, driver.CreateCountry(country))
//	assert.NoError(t, driver.CreateMaterial(material))
//	assert.NoError(t, driver.CreateColor(color))
//
//	// Создаём продукт
//	product := &views.ProductId{
//		Id:          "prod1",
//		Title:       "FilterProduct",
//		Article:     "F-001",
//		Brand:       brand.Id,
//		Category:    category.Id,
//		Country:     country.Id,
//		Width:       60,
//		Height:      110,
//		Depth:       40,
//		Materials:   []string{material.Id},
//		Colors:      []string{color.Id},
//		Photos:      []string{"x.jpg"},
//		Seems:       []string{},
//		Price:       888,
//		Description: "Filtered",
//	}
//	assert.NoError(t, driver.CreateProduct(product))
//
//	tests := []struct {
//		name   string
//		filter views.ProductFilter
//	}{
//		{"by brand", views.ProductFilter{Brand: []string{brand.Id}}},
//		{"by category", views.ProductFilter{Category: []string{category.Id}}},
//		{"by country", views.ProductFilter{Country: []string{country.Id}}},
//		{"by material", views.ProductFilter{Materials: []string{material.Id}}},
//		{"by color", views.ProductFilter{Colors: []string{color.Id}}},
//		{"by size", views.ProductFilter{MinWidth: 50, MaxWidth: 70, MinHeight: 100, MaxHeight: 120, MinDepth: 30, MaxDepth: 50}},
//		{"by price", views.ProductFilter{MinPrice: 500, MaxPrice: 1000}},
//		{"with sort asc", views.ProductFilter{SortBy: "price", SortOrder: "asc"}},
//		{"with sort desc", views.ProductFilter{SortBy: "width", SortOrder: "desc"}},
//		{"with limit and offset", views.ProductFilter{Limit: 1, Offset: 0}},
//		{"complex", views.ProductFilter{
//			Brand:     []string{brand.Id},
//			Category:  []string{category.Id},
//			Country:   []string{country.Id},
//			Materials: []string{material.Id},
//			Colors:    []string{color.Id},
//			MinWidth:  50, MaxWidth: 70,
//			MinHeight: 100, MaxHeight: 120,
//			MinDepth: 30, MaxDepth: 50,
//			MinPrice: 800, MaxPrice: 1000,
//			SortBy: "title", SortOrder: "asc",
//			Limit: 10,
//		}},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			result, err := driver.FilterProducts(&tt.filter)
//			assert.NoError(t, err)
//			assert.NotEmpty(t, result)
//			log.Printf("Filter '%s' returned %d product(s)\n", tt.name, len(result))
//		})
//	}
//}
