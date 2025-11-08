package products

import (
	"context"
	"errors"
	"gateway/internal/utils/convert"
	"gateway/internal/utils/format"
	"gateway/internal/views"
	"log"
	"os"
	"path/filepath"
	"strings"

	productsRPC "github.com/autumnterror/volha-proto/gen/products"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (c *Client) SearchProducts(ctx context.Context, filter *views.ProductSearch) ([]views.Product, error) {
	const op = "grpc.client.SearchProducts"

	res, err := c.api.SearchProducts(ctx, &productsRPC.ProductSearch{
		Id:      filter.Id,
		Title:   filter.Title,
		Article: filter.Article,
	})
	if err != nil {
		return nil, format.Error(op, err)
	}

	return convert.ToProductViewList(res), nil
}

func (c *Client) GetProduct(ctx context.Context, id string) (*views.Product, error) {
	const op = "grpc.client.GetAllProducts"

	resp, err := c.api.GetProduct(ctx, &productsRPC.Id{Id: id})
	if err != nil {
		return nil, format.Error(op, err)
	}
	return convert.ToProductView(resp), nil
}

func (c *Client) CreateProduct(ctx context.Context, p *views.ProductId) error {
	const op = "grpc.client.CreateProduct"

	if _, err := c.api.CreateProduct(ctx, convert.ToProductIdRPC(p)); err != nil {
		return format.Error(op, err)
	}
	return nil
}

func (c *Client) UpdateProduct(ctx context.Context, p *views.ProductId) error {
	const op = "grpc.client.UpdateProduct"

	if _, err := c.api.UpdateProduct(ctx, convert.ToProductIdRPC(p)); err != nil {
		return format.Error(op, err)
	}
	return nil
}

func (c *Client) DeleteProduct(ctx context.Context, id string) error {
	const op = "grpc.client.DeleteProduct"

	exitCleanPhoto := make(chan bool, 1)
	go func() {
		resp, err := c.api.GetProduct(ctx, &productsRPC.Id{Id: id})
		if err != nil {
			exitCleanPhoto <- true
			return
		}
		for _, path := range resp.Photos {
			if err := deleteImage(path); err != nil {
				log.Println(format.Error(op, err))
			}
		}
		exitCleanPhoto <- true
	}()

	if _, err := c.api.DeleteProduct(ctx, &productsRPC.Id{Id: id}); err != nil {
		return format.Error(op, err)
	}
	<-exitCleanPhoto
	return nil
}

func deleteImage(relPath string) error {

	path := filepath.Join("images", relPath)

	cleanPath := strings.TrimPrefix(path, "./")
	if !strings.HasPrefix(cleanPath, "images/") && !strings.HasPrefix(cleanPath, "images\\") {
		return errors.New("invalid image path " + cleanPath)
	}

	fullPath := filepath.Join(".", cleanPath)
	fullPath = filepath.Clean(fullPath)

	if !strings.HasPrefix(fullPath, filepath.Clean("./images")) {
		return errors.New("unsafe path")
	}

	if err := os.Remove(fullPath); err != nil {
		return err
	}
	return nil
}

func (c *Client) GetAllProducts(ctx context.Context) ([]views.Product, error) {
	const op = "grpc.client.GetAllProducts"

	resp, err := c.api.GetAllProducts(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, format.Error(op, err)
	}
	return convert.ToProductViewList(resp), nil
}

func (c *Client) FilterProducts(ctx context.Context, f *views.ProductFilter) ([]views.Product, error) {
	const op = "grpc.client.FilterProducts"

	resp, err := c.api.FilterProducts(ctx, convert.ToProductFilterRPC(f))
	if err != nil {
		return nil, format.Error(op, err)
	}
	return convert.ToProductViewList(resp), nil
}

// BRAND

func (c *Client) CreateBrand(ctx context.Context, b *views.Brand) error {
	const op = "grpc.client.CreateBrand"
	_, err := c.api.CreateBrand(ctx, convert.ToBrandRPC(b))
	return format.Error(op, err)
}

func (c *Client) UpdateBrand(ctx context.Context, b *views.Brand) error {
	const op = "grpc.client.UpdateBrand"
	_, err := c.api.UpdateBrand(ctx, convert.ToBrandRPC(b))
	return format.Error(op, err)
}

func (c *Client) DeleteBrand(ctx context.Context, id string) error {
	const op = "grpc.client.DeleteBrand"
	_, err := c.api.DeleteBrand(ctx, &productsRPC.Id{Id: id})
	return format.Error(op, err)
}

func (c *Client) GetAllBrands(ctx context.Context) ([]views.Brand, error) {
	const op = "grpc.client.GetAllBrands"
	list, err := c.api.GetAllBrands(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, format.Error(op, err)
	}
	return convert.ToBrandList(list), nil
}

// CATEGORY

func (c *Client) CreateCategory(ctx context.Context, ct *views.Category) error {
	const op = "grpc.client.CreateCategory"
	_, err := c.api.CreateCategory(ctx, convert.ToCategoryRPC(ct))
	return format.Error(op, err)
}

func (c *Client) UpdateCategory(ctx context.Context, ct *views.Category) error {
	const op = "grpc.client.UpdateCategory"
	_, err := c.api.UpdateCategory(ctx, convert.ToCategoryRPC(ct))
	return format.Error(op, err)
}

func (c *Client) DeleteCategory(ctx context.Context, id string) error {
	const op = "grpc.client.DeleteCategory"
	_, err := c.api.DeleteCategory(ctx, &productsRPC.Id{Id: id})
	return format.Error(op, err)
}

func (c *Client) GetAllCategories(ctx context.Context) ([]views.Category, error) {
	const op = "grpc.client.GetAllCategories"
	list, err := c.api.GetAllCategories(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, format.Error(op, err)
	}
	return convert.ToCategoryList(list), nil
}

// COUNTRY

func (c *Client) CreateCountry(ctx context.Context, cn *views.Country) error {
	const op = "grpc.client.CreateCountry"
	_, err := c.api.CreateCountry(ctx, convert.ToCountryRPC(cn))
	return format.Error(op, err)
}

func (c *Client) UpdateCountry(ctx context.Context, cn *views.Country) error {
	const op = "grpc.client.UpdateCountry"
	_, err := c.api.UpdateCountry(ctx, convert.ToCountryRPC(cn))
	return format.Error(op, err)
}

func (c *Client) DeleteCountry(ctx context.Context, id string) error {
	const op = "grpc.client.DeleteCountry"
	_, err := c.api.DeleteCountry(ctx, &productsRPC.Id{Id: id})
	return format.Error(op, err)
}

func (c *Client) GetAllCountries(ctx context.Context) ([]views.Country, error) {
	const op = "grpc.client.GetAllCountries"
	list, err := c.api.GetAllCountries(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, format.Error(op, err)
	}
	return convert.ToCountryViewList(list), nil
}

// COLOR

func (c *Client) CreateColor(ctx context.Context, color *views.Color) error {
	const op = "grpc.client.CreateColor"
	_, err := c.api.CreateColor(ctx, convert.ToColorRPC(color))
	return format.Error(op, err)
}

func (c *Client) UpdateColor(ctx context.Context, color *views.Color) error {
	const op = "grpc.client.UpdateColor"
	_, err := c.api.UpdateColor(ctx, convert.ToColorRPC(color))
	return format.Error(op, err)
}

func (c *Client) DeleteColor(ctx context.Context, id string) error {
	const op = "grpc.client.DeleteColor"
	_, err := c.api.DeleteColor(ctx, &productsRPC.Id{Id: id})
	return format.Error(op, err)
}

func (c *Client) GetAllColors(ctx context.Context) ([]views.Color, error) {
	const op = "grpc.client.GetAllColors"
	list, err := c.api.GetAllColors(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, format.Error(op, err)
	}
	return convert.ToColorViewList(list), nil
}

// MATERIAL

func (c *Client) CreateMaterial(ctx context.Context, mat *views.Material) error {
	const op = "grpc.client.CreateMaterial"
	_, err := c.api.CreateMaterial(ctx, convert.ToMaterialRPC(mat))
	return format.Error(op, err)
}

func (c *Client) UpdateMaterial(ctx context.Context, mat *views.Material) error {
	const op = "grpc.client.UpdateMaterial"
	_, err := c.api.UpdateMaterial(ctx, convert.ToMaterialRPC(mat))
	return format.Error(op, err)
}

func (c *Client) DeleteMaterial(ctx context.Context, id string) error {
	const op = "grpc.client.DeleteMaterial"
	_, err := c.api.DeleteMaterial(ctx, &productsRPC.Id{Id: id})
	return format.Error(op, err)
}

func (c *Client) GetAllMaterials(ctx context.Context) ([]views.Material, error) {
	const op = "grpc.client.GetAllMaterials"
	list, err := c.api.GetAllMaterials(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, format.Error(op, err)
	}
	return convert.ToMaterialViewList(list), nil
}

// GET ALL DICTIONARIES

func (c *Client) GetAllDictionaries(ctx context.Context) (*views.Dictionaries, error) {
	const op = "grpc.client.GetAllDictionaries"
	resp, err := c.api.GetDictionaries(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, format.Error(op, err)
	}
	return convert.ToDictionariesView(resp), nil
}

func (c *Client) GetAllDictionariesByCategory(ctx context.Context, id string) (*views.Dictionaries, error) {
	const op = "grpc.client.GetAllDictionaries"
	resp, err := c.api.GetDictionariesByCategory(ctx, &productsRPC.Id{Id: id})
	if err != nil {
		return nil, format.Error(op, err)
	}
	return convert.ToDictionariesViewFromByCategory(resp), nil
}

//Color photos

func (c *Client) CreateProductColorPhotos(ctx context.Context, pc *views.ProductColorPhotos) error {
	const op = "grpc.client.CreateProductColorPhotos"
	_, err := c.api.CreateProductColorPhotos(ctx, convert.ToColorPhotoRPC(pc))
	if err != nil {
		return format.Error(op, err)
	}
	return nil
}

func (c *Client) UpdateProductColorPhotos(ctx context.Context, pc *views.ProductColorPhotos) error {
	const op = "grpc.client.CreateProductColorPhotos"
	_, err := c.api.UpdateProductColorPhotos(ctx, convert.ToColorPhotoRPC(pc))
	if err != nil {
		return format.Error(op, err)
	}
	return nil
}

func (c *Client) DeleteProductColorPhotos(ctx context.Context, productId, colorId string) error {
	const op = "grpc.client.DeleteProductColorPhotos"
	_, err := c.api.DeleteProductColorPhotos(ctx, &productsRPC.ProductColorPhotosId{
		ProductId: productId,
		ColorId:   colorId,
	})
	if err != nil {
		return format.Error(op, err)
	}
	return nil
}

func (c *Client) GetAllProductColorPhotos(ctx context.Context) ([]views.ProductColorPhotos, error) {
	const op = "grpc.client.GetAllProductColorPhotos"
	resp, err := c.api.GetAllProductColorPhotos(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, format.Error(op, err)
	}
	return convert.ToColorPhotoList(resp), nil
}

func (c *Client) GetPhotosByProductAndColor(ctx context.Context, productId, colorId string) ([]string, error) {
	const op = "grpc.client.GetPhotosByProductAndColor"
	resp, err := c.api.GetPhotosByProductAndColor(ctx, &productsRPC.ProductColorPhotosId{
		ProductId: productId,
		ColorId:   colorId,
	})
	if err != nil {
		return nil, format.Error(op, err)
	}
	return resp.GetPhotos(), nil
}
