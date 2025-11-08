package convert

import (
	productsRPC "github.com/autumnterror/volha-proto/gen/products"
	"productService/internal/views"
)

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

func ToProductViewId(req *productsRPC.ProductId) *views.ProductId {
	return &views.ProductId{
		Id:          req.GetId(),
		Title:       req.GetTitle(),
		Article:     req.GetArticle(),
		Brand:       req.GetBrand(),
		Category:    req.GetCategory(),
		Country:     req.GetCountry(),
		Width:       int(req.GetWidth()),
		Height:      int(req.GetHeight()),
		Depth:       int(req.GetDepth()),
		Materials:   req.GetMaterials(),
		Colors:      req.GetColors(),
		Photos:      req.GetPhotos(),
		Seems:       req.GetSeems(),
		Price:       int(req.GetPrice()),
		Description: req.GetDescription(),
	}
}

func ToProductViewIdFromProduct(req *productsRPC.Product) *views.ProductId {
	return &views.ProductId{
		Id:          req.GetId(),
		Title:       req.GetTitle(),
		Article:     req.GetArticle(),
		Brand:       req.GetBrand().GetId(),
		Category:    req.GetCategory().GetId(),
		Country:     req.GetCountry().GetId(),
		Width:       int(req.GetWidth()),
		Height:      int(req.GetHeight()),
		Depth:       int(req.GetDepth()),
		Materials:   ToStringIds(req.GetMaterials()),
		Colors:      ToStringIds(req.GetColors()),
		Photos:      req.GetPhotos(),
		Seems:       ToViewsProductIdSlice(req.GetSeems()),
		Price:       int(req.GetPrice()),
		Description: req.GetDescription(),
	}
}

func ToViewsProductIdSlice(list []*productsRPC.Product) []string {
	var out []string
	for _, p := range list {
		out = append(out, p.Id)
	}
	return out
}

func ToViewsProductSlice(list []*productsRPC.Product) []views.Product {
	var out []views.Product
	for _, p := range list {
		out = append(out, *ToProductView(p))
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

func ToProductSearch(r *productsRPC.ProductSearch) *views.ProductSearch {
	return &views.ProductSearch{
		Id:      r.GetId(),
		Title:   r.GetTitle(),
		Article: r.GetArticle(),
	}
}

func ToStringIds[T interface{ GetId() string }](in []T) []string {
	var ids []string
	for _, e := range in {
		ids = append(ids, e.GetId())
	}
	return ids
}
