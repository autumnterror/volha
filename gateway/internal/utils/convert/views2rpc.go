package convert

import (
	"gateway/internal/views"
	productsRPC "github.com/autumnterror/volha-proto/gen/products"
)

func ToBrandRPC(b *views.Brand) *productsRPC.Brand {
	return &productsRPC.Brand{Id: b.Id, Name: b.Name}
}

func ToCategoryRPC(c *views.Category) *productsRPC.Category {
	return &productsRPC.Category{Id: c.Id, Title: c.Title, Uri: c.Uri, Img: c.Img}
}

func ToCountryRPC(c *views.Country) *productsRPC.Country {
	return &productsRPC.Country{Id: c.Id, Title: c.Title, Friendly: c.Friendly}
}

func ToColorRPC(c *views.Color) *productsRPC.Color {
	return &productsRPC.Color{Id: c.Id, Name: c.Name, Hex: c.Hex}
}

func ToMaterialRPC(m *views.Material) *productsRPC.Material {
	return &productsRPC.Material{Id: m.Id, Title: m.Title}
}

func ToCategoryList(cl *productsRPC.CategoryList) []views.Category {
	var res []views.Category
	for _, c := range cl.Categories {
		res = append(res, views.Category{
			Id:    c.GetId(),
			Title: c.GetTitle(),
			Uri:   c.GetUri(),
			Img:   c.GetImg(),
		})
	}
	return res
}

func ToProductFilterRPC(v *views.ProductFilter) *productsRPC.ProductFilter {
	return &productsRPC.ProductFilter{
		Brand:     v.Brand,
		Country:   v.Country,
		Category:  v.Category,
		Materials: v.Materials,
		Colors:    v.Colors,
		MinWidth:  int32(v.MinWidth),
		MaxWidth:  int32(v.MaxWidth),
		MinHeight: int32(v.MinHeight),
		MaxHeight: int32(v.MaxHeight),
		MinDepth:  int32(v.MinDepth),
		MaxDepth:  int32(v.MaxDepth),
		MinPrice:  int32(v.MinPrice),
		MaxPrice:  int32(v.MaxPrice),
		SortBy:    v.SortBy,
		SortOrder: v.SortOrder,
		Offset:    int32(v.Offset),
		Limit:     int32(v.Limit),
	}
}

func ToProductIdRPC(p *views.ProductId) *productsRPC.ProductId {
	return &productsRPC.ProductId{
		Id:          p.Id,
		Title:       p.Title,
		Article:     p.Article,
		Brand:       p.Brand,
		Category:    p.Category,
		Country:     p.Country,
		Width:       int32(p.Width),
		Height:      int32(p.Height),
		Depth:       int32(p.Depth),
		Materials:   p.Materials,
		Colors:      p.Colors,
		Photos:      p.Photos,
		Seems:       p.Seems,
		Price:       int32(p.Price),
		Description: p.Description,
	}
}

func ToProductIdRPCFromProduct(p *views.Product) *productsRPC.ProductId {
	return &productsRPC.ProductId{
		Id:          p.Id,
		Title:       p.Title,
		Article:     p.Article,
		Brand:       p.Brand.Id,
		Category:    p.Category.Id,
		Country:     p.Country.Id,
		Width:       int32(p.Width),
		Height:      int32(p.Height),
		Depth:       int32(p.Depth),
		Materials:   getIdMaterial(p.Materials),
		Colors:      getIdColor(p.Colors),
		Photos:      p.Photos,
		Seems:       getIdSeem(p.Seems),
		Price:       int32(p.Price),
		Description: p.Description,
	}
}

func getIdMaterial(list []views.Material) []string {
	var res []string
	for _, i := range list {
		res = append(res, i.Id)
	}
	return res
}
func getIdColor(list []views.Color) []string {
	var res []string
	for _, i := range list {
		res = append(res, i.Id)
	}
	return res
}
func getIdSeem(list []views.Product) []string {
	var res []string
	for _, i := range list {
		res = append(res, i.Id)
	}
	return res
}

func ToColorPhotoRPC(in *views.ProductColorPhotos) *productsRPC.ProductColorPhotos {
	return &productsRPC.ProductColorPhotos{
		ProductId: in.ProductId,
		ColorId:   in.ColorId,
		Photos:    in.Photos,
	}
}
func ToColorPhotoIdRPC(in *views.ProductColorPhotosId) *productsRPC.ProductColorPhotosId {
	return &productsRPC.ProductColorPhotosId{
		ProductId: in.ProductId,
		ColorId:   in.ColorId,
	}
}
