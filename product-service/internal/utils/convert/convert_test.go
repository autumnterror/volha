package convert

import (
	productsRPC "github.com/autumnterror/volha-proto/gen/products"
	"productService/internal/views"
	"reflect"
	"testing"
)

func TestToRPCProductAndBack(t *testing.T) {
	original := &views.Product{
		Id:      "1",
		Title:   "Test",
		Article: "ART123",
		Brand:   views.Brand{Id: "b1", Name: "BrandX"},
		Category: views.Category{
			Id:    "cat1",
			Title: "CatTitle",
			Uri:   "uri123",
		},
		Country: views.Country{
			Id:       "c1",
			Title:    "CountryTitle",
			Friendly: "FriendlyName",
		},
		Width: 10, Height: 20, Depth: 30,
		Price:       999,
		Description: "Awesome product",
		Materials: []views.Material{
			{Id: "m1", Title: "Leather"},
		},
		Colors: []views.Color{
			{Id: "clr1", Name: "Red", Hex: "#FF0000"},
		},
		Photos: []string{"img1.jpg", "img2.jpg"},
		Seems: []views.Product{
			{Id: "2", Title: "Sim"},
		},
	}

	rpc := ToRPCProduct(original)
	converted := ToProductView(rpc)

	if !reflect.DeepEqual(original.Id, converted.Id) ||
		!reflect.DeepEqual(original.Brand.Id, converted.Brand.Id) ||
		len(converted.Seems) != len(original.Seems) {
		t.Errorf("Product conversion failed.\nOriginal: %+v\nConverted: %+v", original, converted)
	}
}

func TestToProductIdConversion(t *testing.T) {
	rpc := &productsRPC.Product{
		Id:      "3",
		Title:   "IDed",
		Article: "ART456",
		Brand:   &productsRPC.Brand{Id: "bid"},
		Category: &productsRPC.Category{
			Id: "cid",
		},
		Country: &productsRPC.Country{Id: "ctid"},
		Width:   10, Height: 10, Depth: 10, Price: 500,
		Materials: []*productsRPC.Material{
			{Id: "mat1"},
		},
		Colors: []*productsRPC.Color{
			{Id: "col1"},
		},
		Photos: []string{"ph1"},
		Seems: []*productsRPC.Product{
			{Id: "similar"},
		},
		Description: "Description",
	}

	converted := ToProductViewIdFromProduct(rpc)

	if converted.Brand != "bid" || len(converted.Materials) != 1 || len(converted.Seems) != 1 {
		t.Errorf("ToProductId conversion failed: %+v", converted)
	}
}

func TestToProductFilterViews(t *testing.T) {
	rpc := &productsRPC.ProductFilter{
		Brand:    []string{"b1", "b2"},
		Category: []string{"c1"},
		MinPrice: 100,
		MaxPrice: 999,
		SortBy:   "price",
		Limit:    20,
	}

	result := ToProductFilterView(rpc).(*views.ProductFilter)

	if result.MinPrice != 100 || result.Limit != 20 || len(result.Brand) != 2 {
		t.Errorf("Filter conversion incorrect: %+v", result)
	}
}

func TestToProductSearch(t *testing.T) {
	rpc := &productsRPC.ProductSearch{
		Id:      "id",
		Title:   "T",
		Article: "A",
	}
	res := ToProductSearch(rpc)

	if res.Id != "id" || res.Title != "T" || res.Article != "A" {
		t.Errorf("ProductSearch conversion failed: %+v", res)
	}
}

func TestMaterialColorConversion(t *testing.T) {
	materials := []*productsRPC.Material{{Id: "m1", Title: "Mat"}}
	colors := []*productsRPC.Color{{Id: "c1", Name: "Red", Hex: "#FF0000"}}

	m := ToMaterialView(materials)
	c := ToColorView(colors)

	if len(m) != 1 || m[0].Title != "Mat" {
		t.Errorf("Material conversion failed")
	}
	if len(c) != 1 || c[0].Hex != "#FF0000" {
		t.Errorf("Color conversion failed")
	}
}

func TestToViewsProductSlice(t *testing.T) {
	rpc := &productsRPC.Product{
		Id:       "test",
		Brand:    &productsRPC.Brand{Id: "b"},
		Category: &productsRPC.Category{Id: "c"},
		Country:  &productsRPC.Country{Id: "cc"},
	}
	list := ToViewsProductSlice([]*productsRPC.Product{rpc})
	if len(list) != 1 || list[0].Id != "test" {
		t.Errorf("ToViewsProductSlice failed: %+v", list)
	}
}

func TestToRPCProductSlice(t *testing.T) {
	viewsProd := views.Product{Id: "X"}
	rpc := ToRPCProductSlice([]views.Product{viewsProd})
	if len(rpc) != 1 || rpc[0].Id != "X" {
		t.Errorf("ToRPCProductSlice failed: %+v", rpc)
	}
}
