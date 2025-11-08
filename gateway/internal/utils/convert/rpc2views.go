package convert

import (
	"gateway/internal/views"
	productsRPC "github.com/autumnterror/volha-proto/gen/products"
)

func ToProductViewList(r *productsRPC.ProductList) []views.Product {
	var list []views.Product
	for _, p := range r.Products {
		list = append(list, *ToProductView(p))
	}
	return list
}

func ToColorViewList(c *productsRPC.ColorList) []views.Color {
	var list []views.Color
	for _, x := range c.Colors {
		list = append(list, views.Color{
			Id:   x.Id,
			Name: x.Name,
			Hex:  x.Hex,
		})
	}
	return list
}

func ToMaterialViewList(m *productsRPC.MaterialList) []views.Material {
	var list []views.Material
	for _, x := range m.Materials {
		list = append(list, views.Material{
			Id:    x.Id,
			Title: x.Title,
		})
	}
	return list
}

func ToCountryViewList(c *productsRPC.CountryList) []views.Country {
	var list []views.Country
	for _, x := range c.Countries {
		list = append(list, views.Country{
			Id:       x.Id,
			Title:    x.Title,
			Friendly: x.Friendly,
		})
	}
	return list
}

func ToCategoryViewList(in *productsRPC.CategoryList) []views.Category {
	var list []views.Category
	for _, c := range in.Categories {
		list = append(list, views.Category{
			Id:    c.Id,
			Title: c.Title,
			Uri:   c.Uri,
			Img:   c.GetImg(),
		})
	}
	return list
}

func ToDictionariesView(in *productsRPC.Dictionaries) *views.Dictionaries {
	return &views.Dictionaries{
		Brands:     ToBrandList(in.Brands),
		Categories: ToCategoryViewList(in.Categories),
		Countries:  ToCountryViewList(in.Countries),
		Materials:  ToMaterialViewList(in.Materials),
		Colors:     ToColorViewList(in.Colors),
		MinPrice:   int(in.MinPrice),
		MaxPrice:   int(in.MaxPrice),
		MinWidth:   int(in.MinWidth),
		MaxWidth:   int(in.MaxWidth),
		MinHeight:  int(in.MinHeight),
		MaxHeight:  int(in.MaxHeight),
		MinDepth:   int(in.MinDepth),
		MaxDepth:   int(in.MaxDepth),
	}
}
func ToDictionariesViewFromByCategory(in *productsRPC.DictionariesByCategory) *views.Dictionaries {
	return &views.Dictionaries{
		Brands:     ToBrandList(in.Brands),
		Categories: []views.Category{},
		Countries:  ToCountryViewList(in.Countries),
		Materials:  ToMaterialViewList(in.Materials),
		Colors:     ToColorViewList(in.Colors),
		MinPrice:   int(in.MinPrice),
		MaxPrice:   int(in.MaxPrice),
		MinWidth:   int(in.MinWidth),
		MaxWidth:   int(in.MaxWidth),
		MinHeight:  int(in.MinHeight),
		MaxHeight:  int(in.MaxHeight),
		MinDepth:   int(in.MinDepth),
		MaxDepth:   int(in.MaxDepth),
	}
}

func ToProductView(req *productsRPC.Product) *views.Product {
	return &views.Product{
		Id:      req.GetId(),
		Title:   req.GetTitle(),
		Article: req.GetArticle(),
		Brand: views.Brand{
			Id:   req.GetBrand().GetId(),
			Name: req.GetBrand().GetName(),
		},
		Category: views.Category{
			Id:    req.GetCategory().GetId(),
			Title: req.GetCategory().GetTitle(),
			Uri:   req.GetCategory().GetUri(),
		},
		Country: views.Country{
			Id:       req.GetCountry().GetId(),
			Title:    req.GetCountry().GetTitle(),
			Friendly: req.GetCountry().GetFriendly(),
		},
		Width:       int(req.GetWidth()),
		Height:      int(req.GetHeight()),
		Depth:       int(req.GetDepth()),
		Materials:   ToMaterialView(req.GetMaterials()),
		Colors:      ToColorView(req.GetColors()),
		Photos:      req.GetPhotos(),
		Seems:       ToViewsProductSlice(req.GetSeems()),
		Price:       int(req.GetPrice()),
		Description: req.GetDescription(),
	}
}
func ToViewsProductSlice(in []*productsRPC.Product) []views.Product {
	var out []views.Product
	for _, p := range in {
		out = append(out, *ToProductView(p))
	}
	return out
}

func ToColorView(in []*productsRPC.Color) []views.Color {
	var out []views.Color
	for _, c := range in {
		out = append(out, views.Color{
			Id:   c.GetId(),
			Name: c.GetName(),
			Hex:  c.GetHex(),
		})
	}
	return out
}

func ToMaterialView(in []*productsRPC.Material) []views.Material {
	var out []views.Material
	for _, m := range in {
		out = append(out, views.Material{
			Id:    m.GetId(),
			Title: m.GetTitle(),
		})
	}
	return out
}

func ToBrandList(in *productsRPC.BrandList) []views.Brand {
	var list []views.Brand
	for _, b := range in.Brands {
		list = append(list, views.Brand{
			Id:   b.Id,
			Name: b.Name,
		})
	}
	return list
}

func ToColorPhotoList(in *productsRPC.ProductColorPhotosList) []views.ProductColorPhotos {
	var list []views.ProductColorPhotos
	for _, b := range in.Items {
		list = append(list, views.ProductColorPhotos{
			ProductId: b.ProductId,
			ColorId:   b.ColorId,
			Photos:    b.Photos,
		})
	}
	return list
}
