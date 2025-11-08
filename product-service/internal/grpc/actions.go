package grpc

import (
	"context"
	productsRPC "github.com/autumnterror/volha-proto/gen/products"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"productService/internal/utils/convert"
	"productService/internal/utils/format"
	"productService/internal/views"
)

// ---------- Product ----------

func (s *ServerAPI) CreateProduct(ctx context.Context, req *productsRPC.ProductId) (*emptypb.Empty, error) {
	const op = "productsRPC.ServerAPI.CreateProduct"
	log.Println(format.String(op, req))
	return handleCRUDResponse(ctx, op, func() error {
		return s.API.CreateProduct(convert.ToProductViewId(req))
	})
}

func (s *ServerAPI) UpdateProduct(ctx context.Context, req *productsRPC.ProductId) (*emptypb.Empty, error) {
	const op = "productsRPC.ServerAPI.UpdateProduct"
	log.Println(format.String(op, req))
	return handleCRUDResponse(ctx, op, func() error {
		return s.API.UpdateProduct(convert.ToProductViewId(req), req.GetId())
	})
}

func (s *ServerAPI) DeleteProduct(ctx context.Context, req *productsRPC.Id) (*emptypb.Empty, error) {
	const op = "productsRPC.ServerAPI.DeleteProduct"
	log.Println(format.String(op, req))
	return handleCRUDResponse(ctx, op, func() error {
		return s.API.DeleteProduct(req.GetId())
	})
}

func (s *ServerAPI) GetAllProducts(ctx context.Context, _ *emptypb.Empty) (*productsRPC.ProductList, error) {
	const op = "productsRPC.ServerAPI.GetAllProducts"
	log.Println(op)
	data, err := handleListResponse(ctx, op, s.API.GetAllProducts, convert.ToProductList)
	if err != nil {
		return nil, err
	}
	return data.(*productsRPC.ProductList), nil
}

// ---------- Brand ----------

func (s *ServerAPI) CreateBrand(ctx context.Context, req *productsRPC.Brand) (*emptypb.Empty, error) {
	const op = "productsRPC.CreateBrand"
	log.Println(format.String(op, req))
	return handleCRUDResponse(ctx, op, func() error {
		return s.API.CreateBrand(&views.Brand{Id: req.Id, Name: req.Name})
	})
}
func (s *ServerAPI) UpdateBrand(ctx context.Context, req *productsRPC.Brand) (*emptypb.Empty, error) {
	const op = "productsRPC.UpdateBrand"
	log.Println(format.String(op, req))
	return handleCRUDResponse(ctx, op, func() error {
		return s.API.UpdateBrand(&views.Brand{Id: req.Id, Name: req.Name}, req.Id)
	})
}
func (s *ServerAPI) DeleteBrand(ctx context.Context, req *productsRPC.Id) (*emptypb.Empty, error) {
	const op = "productsRPC.DeleteBrand"
	log.Println(format.String(op, req))
	return handleCRUDResponse(ctx, op, func() error {
		return s.API.DeleteBrand(req.Id)
	})
}
func (s *ServerAPI) GetAllBrands(ctx context.Context, _ *emptypb.Empty) (*productsRPC.BrandList, error) {
	const op = "productsRPC.GetAllBrands"
	log.Println(op)
	data, err := handleListResponse(ctx, op, s.API.GetAllBrands, convert.ToBrandList)
	if err != nil {
		return nil, err
	}
	return data.(*productsRPC.BrandList), nil
}

// ---------- Category ----------

func (s *ServerAPI) CreateCategory(ctx context.Context, req *productsRPC.Category) (*emptypb.Empty, error) {
	const op = "productsRPC.CreateCategory"
	log.Println(format.String(op, req))
	return handleCRUDResponse(ctx, op, func() error {
		return s.API.CreateCategory(&views.Category{Id: req.Id, Title: req.Title, Uri: req.Uri, Img: req.Img})
	})
}
func (s *ServerAPI) UpdateCategory(ctx context.Context, req *productsRPC.Category) (*emptypb.Empty, error) {
	const op = "productsRPC.UpdateCategory"
	log.Println(format.String(op, req))
	return handleCRUDResponse(ctx, op, func() error {
		return s.API.UpdateCategory(&views.Category{Id: req.Id, Title: req.Title, Uri: req.Uri, Img: req.Img}, req.Id)
	})
}
func (s *ServerAPI) DeleteCategory(ctx context.Context, req *productsRPC.Id) (*emptypb.Empty, error) {
	const op = "productsRPC.DeleteCategory"
	log.Println(format.String(op, req))
	return handleCRUDResponse(ctx, op, func() error {
		return s.API.DeleteCategory(req.Id)
	})
}
func (s *ServerAPI) GetAllCategories(ctx context.Context, _ *emptypb.Empty) (*productsRPC.CategoryList, error) {
	const op = "productsRPC.GetAllCategories"
	log.Println(op)
	data, err := handleListResponse(ctx, op, s.API.GetAllCategories, convert.ToCategoryList)
	if err != nil {
		return nil, err
	}
	return data.(*productsRPC.CategoryList), nil
}

// ---------- Country ----------

func (s *ServerAPI) CreateCountry(ctx context.Context, req *productsRPC.Country) (*emptypb.Empty, error) {
	const op = "productsRPC.CreateCountry"
	log.Println(format.String(op, req))
	return handleCRUDResponse(ctx, op, func() error {
		return s.API.CreateCountry(&views.Country{Id: req.Id, Title: req.Title, Friendly: req.Friendly})
	})
}
func (s *ServerAPI) UpdateCountry(ctx context.Context, req *productsRPC.Country) (*emptypb.Empty, error) {
	const op = "productsRPC.UpdateCountry"
	log.Println(format.String(op, req))
	return handleCRUDResponse(ctx, op, func() error {
		return s.API.UpdateCountry(&views.Country{Id: req.Id, Title: req.Title, Friendly: req.Friendly}, req.Id)
	})
}
func (s *ServerAPI) DeleteCountry(ctx context.Context, req *productsRPC.Id) (*emptypb.Empty, error) {
	const op = "productsRPC.DeleteCountry"
	log.Println(format.String(op, req))
	return handleCRUDResponse(ctx, op, func() error {
		return s.API.DeleteCountry(req.Id)
	})
}
func (s *ServerAPI) GetAllCountries(ctx context.Context, _ *emptypb.Empty) (*productsRPC.CountryList, error) {
	const op = "productsRPC.GetAllCountries"
	log.Println(op)
	data, err := handleListResponse(ctx, op, s.API.GetAllCountries, convert.ToCountryList)
	if err != nil {
		return nil, err
	}
	return data.(*productsRPC.CountryList), nil
}

// ---------- Material ----------

func (s *ServerAPI) CreateMaterial(ctx context.Context, req *productsRPC.Material) (*emptypb.Empty, error) {
	const op = "productsRPC.CreateMaterial"
	log.Println(format.String(op, req))
	return handleCRUDResponse(ctx, op, func() error {
		return s.API.CreateMaterial(&views.Material{Id: req.Id, Title: req.Title})
	})
}
func (s *ServerAPI) UpdateMaterial(ctx context.Context, req *productsRPC.Material) (*emptypb.Empty, error) {
	const op = "productsRPC.UpdateMaterial"
	log.Println(format.String(op, req))
	return handleCRUDResponse(ctx, op, func() error {
		return s.API.UpdateMaterial(&views.Material{Id: req.Id, Title: req.Title}, req.Id)
	})
}
func (s *ServerAPI) DeleteMaterial(ctx context.Context, req *productsRPC.Id) (*emptypb.Empty, error) {
	const op = "productsRPC.DeleteMaterial"
	log.Println(format.String(op, req))
	return handleCRUDResponse(ctx, op, func() error {
		return s.API.DeleteMaterial(req.Id)
	})
}
func (s *ServerAPI) GetAllMaterials(ctx context.Context, _ *emptypb.Empty) (*productsRPC.MaterialList, error) {
	const op = "productsRPC.GetAllMaterials"
	log.Println(op)
	data, err := handleListResponse(ctx, op, s.API.GetAllMaterials, convert.ToMaterialList)
	if err != nil {
		return nil, err
	}
	return data.(*productsRPC.MaterialList), nil
}

// ---------- Color ----------

func (s *ServerAPI) CreateColor(ctx context.Context, req *productsRPC.Color) (*emptypb.Empty, error) {
	const op = "productsRPC.CreateColor"
	log.Println(format.String(op, req))
	return handleCRUDResponse(ctx, op, func() error {
		return s.API.CreateColor(&views.Color{Id: req.Id, Name: req.Name, Hex: req.Hex})
	})
}
func (s *ServerAPI) UpdateColor(ctx context.Context, req *productsRPC.Color) (*emptypb.Empty, error) {
	const op = "productsRPC.UpdateColor"
	log.Println(format.String(op, req))
	return handleCRUDResponse(ctx, op, func() error {
		return s.API.UpdateColor(&views.Color{Id: req.Id, Name: req.Name, Hex: req.Hex}, req.Id)
	})
}
func (s *ServerAPI) DeleteColor(ctx context.Context, req *productsRPC.Id) (*emptypb.Empty, error) {
	const op = "productsRPC.DeleteColor"
	log.Println(format.String(op, req))
	return handleCRUDResponse(ctx, op, func() error {
		return s.API.DeleteColor(req.Id)
	})
}
func (s *ServerAPI) GetAllColors(ctx context.Context, _ *emptypb.Empty) (*productsRPC.ColorList, error) {
	const op = "productsRPC.GetAllColors"
	log.Println(op)
	data, err := handleListResponse(ctx, op, s.API.GetAllColors, convert.ToColorList)
	if err != nil {
		return nil, err
	}
	return data.(*productsRPC.ColorList), nil
}

// ---------- ColorPhotos ----------

func (s *ServerAPI) CreateProductColorPhotos(ctx context.Context, req *productsRPC.ProductColorPhotos) (*emptypb.Empty, error) {
	const op = "productsRPC.CreateProductColorPhotos"
	log.Println(format.String(op, req))
	return handleCRUDResponse(ctx, op, func() error {
		return s.API.CreateProductColorPhotos(&views.ProductColorPhotos{
			ProductId: req.ProductId,
			ColorId:   req.ColorId,
			Photos:    req.Photos,
		})
	})
}
func (s *ServerAPI) UpdateProductColorPhotos(ctx context.Context, req *productsRPC.ProductColorPhotos) (*emptypb.Empty, error) {
	const op = "productsRPC.UpdateProductColorPhotos"
	log.Println(format.String(op, req))
	return handleCRUDResponse(ctx, op, func() error {
		return s.API.UpdateProductColorPhotos(&views.ProductColorPhotos{
			ProductId: req.ProductId,
			ColorId:   req.ColorId,
			Photos:    req.Photos,
		})
	})
}
func (s *ServerAPI) DeleteProductColorPhotos(ctx context.Context, req *productsRPC.ProductColorPhotosId) (*emptypb.Empty, error) {
	const op = "productsRPC.DeleteProductColorPhotos"
	log.Println(format.String(op, req))
	return handleCRUDResponse(ctx, op, func() error {
		return s.API.DeleteProductColorPhotos(req.ProductId, req.ColorId)
	})
}

func (s *ServerAPI) GetAllProductColorPhotos(ctx context.Context, _ *emptypb.Empty) (*productsRPC.ProductColorPhotosList, error) {
	const op = "productsRPC.GetAllProductColorPhotos"
	log.Println(op)
	data, err := handleListResponse(ctx, op, s.API.GetAllProductColorPhotos, convert.ToProductColorList)
	if err != nil {
		return nil, err
	}
	return data.(*productsRPC.ProductColorPhotosList), nil
}
