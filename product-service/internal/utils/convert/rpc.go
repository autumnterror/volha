package convert

import (
	productsRPC "github.com/autumnterror/volha-proto/gen/products"
	"productService/internal/views"
)

func ToProductList(list []views.Product) any {
	resp := &productsRPC.ProductList{}

	for _, p := range list {
		// Преобразуем вложенные поля
		brand := &productsRPC.Brand{
			Id:   p.Brand.Id,
			Name: p.Brand.Name,
		}
		category := &productsRPC.Category{
			Id:    p.Category.Id,
			Title: p.Category.Title,
			Uri:   p.Category.Uri,
		}
		country := &productsRPC.Country{
			Id:       p.Country.Id,
			Title:    p.Country.Title,
			Friendly: p.Country.Friendly,
		}

		var materials []*productsRPC.Material
		for _, m := range p.Materials {
			materials = append(materials, &productsRPC.Material{
				Id:    m.Id,
				Title: m.Title,
			})
		}

		var colors []*productsRPC.Color
		for _, c := range p.Colors {
			colors = append(colors, &productsRPC.Color{
				Id:   c.Id,
				Name: c.Name,
				Hex:  c.Hex,
			})
		}

		resp.Products = append(resp.Products, &productsRPC.Product{
			Id:          p.Id,
			Title:       p.Title,
			Article:     p.Article,
			Brand:       brand,
			Country:     country,
			Category:    category,
			Width:       int32(p.Width),
			Height:      int32(p.Height),
			Depth:       int32(p.Depth),
			Materials:   materials,
			Colors:      colors,
			Photos:      p.Photos,
			Seems:       ToRPCProductSlice(p.Seems),
			Price:       int32(p.Price),
			Description: p.Description,
		})
	}

	return resp
}

func ToRPCProduct(p *views.Product) *productsRPC.Product {
	return &productsRPC.Product{
		Id:      p.Id,
		Title:   p.Title,
		Article: p.Article,
		Brand: &productsRPC.Brand{
			Id:   p.Brand.Id,
			Name: p.Brand.Name,
		},
		Category: &productsRPC.Category{
			Id:    p.Category.Id,
			Title: p.Category.Title,
			Uri:   p.Category.Uri,
		},
		Country: &productsRPC.Country{
			Id:       p.Country.Id,
			Title:    p.Country.Title,
			Friendly: p.Country.Friendly,
		},
		Width:       int32(p.Width),
		Height:      int32(p.Height),
		Depth:       int32(p.Depth),
		Materials:   ToRPCMaterialList(p.Materials),
		Colors:      ToRPCColorList(p.Colors),
		Photos:      p.Photos,
		Seems:       ToRPCProductSlice(p.Seems),
		Price:       int32(p.Price),
		Description: p.Description,
	}
}

func ToRPCProductSlice(list []views.Product) []*productsRPC.Product {
	var out []*productsRPC.Product
	for _, p := range list {
		out = append(out, ToRPCProduct(&p))
	}
	return out
}

func ToProductFilterView(p *productsRPC.ProductFilter) any {
	return &views.ProductFilter{
		Brand:     p.Brand,
		Country:   p.Country,
		Category:  p.Category,
		Materials: p.Materials,
		Colors:    p.Colors,
		MinWidth:  int(p.MinWidth),
		MaxWidth:  int(p.MaxWidth),
		MinHeight: int(p.MinHeight),
		MaxHeight: int(p.MaxHeight),
		MinDepth:  int(p.MinDepth),
		MaxDepth:  int(p.MaxDepth),
		MinPrice:  int(p.MinPrice),
		MaxPrice:  int(p.MaxPrice),
		SortBy:    p.SortBy,
		SortOrder: p.SortOrder,
		Offset:    int(p.Offset),
		Limit:     int(p.Limit),
	}
}

func ToRPCMaterialList(list []views.Material) []*productsRPC.Material {
	var out []*productsRPC.Material
	for _, m := range list {
		out = append(out, &productsRPC.Material{
			Id:    m.Id,
			Title: m.Title,
		})
	}
	return out
}

func ToRPCColorList(list []views.Color) []*productsRPC.Color {
	var out []*productsRPC.Color
	for _, c := range list {
		out = append(out, &productsRPC.Color{
			Id:   c.Id,
			Name: c.Name,
			Hex:  c.Hex,
		})
	}
	return out
}

func ToBrandList(bv []views.Brand) any {
	var bl []*productsRPC.Brand

	for _, b := range bv {
		bl = append(bl, &productsRPC.Brand{
			Id:   b.Id,
			Name: b.Name,
		})
	}
	return &productsRPC.BrandList{Brands: bl}
}

func ToCategoryList(bv []views.Category) any {
	var bl []*productsRPC.Category

	for _, b := range bv {
		bl = append(bl, &productsRPC.Category{
			Id:    b.Id,
			Title: b.Title,
			Uri:   b.Uri,
			Img:   b.Img,
		})
	}
	return &productsRPC.CategoryList{Categories: bl}
}

func ToCountryList(bv []views.Country) any {
	var bl []*productsRPC.Country

	for _, b := range bv {
		bl = append(bl, &productsRPC.Country{
			Id:       b.Id,
			Title:    b.Title,
			Friendly: b.Friendly,
		})
	}
	return &productsRPC.CountryList{Countries: bl}
}

func ToMaterialList(bv []views.Material) any {
	var bl []*productsRPC.Material

	for _, b := range bv {
		bl = append(bl, &productsRPC.Material{
			Id:    b.Id,
			Title: b.Title,
		})
	}
	return &productsRPC.MaterialList{Materials: bl}
}

func ToColorList(bv []views.Color) any {
	var bl []*productsRPC.Color

	for _, b := range bv {
		bl = append(bl, &productsRPC.Color{
			Id:   b.Id,
			Name: b.Name,
			Hex:  b.Hex,
		})
	}
	return &productsRPC.ColorList{Colors: bl}
}

func ToProductColorList(bv []views.ProductColorPhotos) any {
	var bl []*productsRPC.ProductColorPhotos

	for _, b := range bv {
		bl = append(bl, &productsRPC.ProductColorPhotos{
			ProductId: b.ProductId,
			ColorId:   b.ColorId,
			Photos:    b.Photos,
		})
	}
	return &productsRPC.ProductColorPhotosList{Items: bl}
}
