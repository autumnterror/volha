package grpc

import (
	"context"
	"log"
	"productService/internal/pkg/psql"
	"productService/internal/utils/convert"
	"productService/internal/utils/format"
	"productService/internal/views"

	productsRPC "github.com/autumnterror/volha-proto/gen/products"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ServerAPI struct {
	productsRPC.UnimplementedProductsServer
	API psql.Repository
}

func Register(
	server *grpc.Server,
	API psql.Repository,
) {
	productsRPC.RegisterProductsServer(server, &ServerAPI{
		API: API,
	})
}
func handleCRUDResponse(ctx context.Context, op string, action func() error) (*emptypb.Empty, error) {
	res := make(chan error, 1)
	go func() {
		res <- action()
	}()
	select {
	case <-ctx.Done():
		log.Println(format.Error(op, status.Error(codes.DeadlineExceeded, "Context dead")))
		return nil, format.Error(op, status.Error(codes.DeadlineExceeded, "Context dead"))
	case err := <-res:
		if err != nil {
			log.Println(format.Error(op, status.Error(codes.Internal, err.Error())))
			return nil, format.Error(op, status.Error(codes.Internal, err.Error()))
		}
		log.Println(format.String(op, "SUCCESS"))
		return &emptypb.Empty{}, nil
	}
}

type result struct {
	data any
	err  error
}

func handleListResponse[T any](ctx context.Context, op string, fetch func() ([]T, error), convert func([]T) any) (any, error) {
	res := make(chan result, 1)
	go func() {
		items, err := fetch()
		if err != nil {
			res <- result{err: format.Error(op, status.Error(codes.Internal, err.Error()))}
			return
		}
		res <- result{data: convert(items)}
	}()

	select {
	case <-ctx.Done():
		log.Println(format.Error(op, status.Error(codes.DeadlineExceeded, "Context dead")))
		return nil, format.Error(op, status.Error(codes.DeadlineExceeded, "Context dead"))
	case r := <-res:
		log.Println(format.String(op, "SUCCESS"))
		return r.data, r.err
	}
}

func (s *ServerAPI) GetAllProducts(ctx context.Context, req *productsRPC.GetAllProductsPagination) (*productsRPC.ProductList, error) {
	const op = "productsRPC.ServerAPI.GetAllProducts"
	log.Println(op)

	type result struct {
		data *productsRPC.ProductList
		err  error
	}
	res := make(chan result, 1)

	if req.GetStart() < 0 {
		return nil, status.Error(codes.InvalidArgument, "start < 0")
	}

	go func() {
		list, err := s.API.GetAllProducts(int(req.GetStart()), int(req.GetEnd()))
		if err != nil {
			log.Println(format.Error(op, err))
			res <- result{
				data: nil,
				err:  status.Error(codes.Internal, err.Error()),
			}
			return
		}

		res <- result{
			data: convert.ToProductList(list).(*productsRPC.ProductList),
			err:  nil,
		}
	}()

	select {
	case <-ctx.Done():
		return nil, status.Error(codes.DeadlineExceeded, "Context deadline exceeded")
	case r := <-res:
		log.Println(format.String(op, "SUCCESS"))
		return r.data, r.err
	}
}

func (s *ServerAPI) SearchProducts(ctx context.Context, req *productsRPC.ProductSearch) (*productsRPC.ProductList, error) {
	const op = "productsRPC.ServerAPI.SearchProducts"

	type result struct {
		data *productsRPC.ProductList
		err  error
	}
	res := make(chan result, 1)

	go func() {
		filter := convert.ToProductSearch(req)
		list, err := s.API.SearchProducts(filter)
		if err != nil {
			log.Println(format.Error(op, err))
			res <- result{
				data: nil,
				err:  status.Error(codes.Internal, err.Error()),
			}
			return
		}

		res <- result{
			data: convert.ToProductList(list).(*productsRPC.ProductList),
			err:  nil,
		}
	}()

	select {
	case <-ctx.Done():
		return nil, status.Error(codes.DeadlineExceeded, "Context deadline exceeded")
	case r := <-res:
		log.Println(format.String(op, "SUCCESS"))
		return r.data, r.err
	}
}

func (s *ServerAPI) FilterProducts(ctx context.Context, req *productsRPC.ProductFilter) (*productsRPC.ProductList, error) {
	const op = "productsRPC.ServerAPI.FilterProducts"
	type result struct {
		data *productsRPC.ProductList
		err  error
	}
	res := make(chan result, 1)

	go func() {
		filter := convert.ToProductFilterView(req)
		list, err := s.API.FilterProducts(filter.(*views.ProductFilter))
		if err != nil {
			log.Println(format.Error(op, err))
			res <- result{
				data: nil,
				err:  status.Error(codes.Internal, err.Error()),
			}
			return
		}

		res <- result{
			data: convert.ToProductList(list).(*productsRPC.ProductList),
			err:  nil,
		}
	}()

	select {
	case <-ctx.Done():
		return nil, status.Error(codes.DeadlineExceeded, "Context deadline exceeded")
	case r := <-res:
		log.Println(format.String(op, "SUCCESS"))
		return r.data, r.err
	}
}

func (s *ServerAPI) GetDictionaries(
	ctx context.Context,
	_ *emptypb.Empty,
) (*productsRPC.Dictionaries, error) {
	const op = "productsRPC.ServerAPI.GetDictionaries"

	type result struct {
		data *productsRPC.Dictionaries
		err  error
	}
	res := make(chan result, 1)

	go func() {
		dict, err := s.API.GetDictionaries()
		if err != nil {
			res <- result{err: format.Error(op, status.Error(codes.Internal, err.Error()))}
			return
		}

		prDict := &productsRPC.Dictionaries{
			Brands:     convert.ToBrandList(dict.Brands).(*productsRPC.BrandList),
			Categories: convert.ToCategoryList(dict.Categories).(*productsRPC.CategoryList),
			Countries:  convert.ToCountryList(dict.Countries).(*productsRPC.CountryList),
			Materials:  convert.ToMaterialList(dict.Materials).(*productsRPC.MaterialList),
			Colors:     convert.ToColorList(dict.Colors).(*productsRPC.ColorList),

			MinPrice:  int32(dict.MinPrice),
			MaxPrice:  int32(dict.MaxPrice),
			MinWidth:  int32(dict.MinWidth),
			MaxWidth:  int32(dict.MaxWidth),
			MinHeight: int32(dict.MinHeight),
			MaxHeight: int32(dict.MaxHeight),
			MinDepth:  int32(dict.MinDepth),
			MaxDepth:  int32(dict.MaxDepth),
		}

		res <- result{data: prDict}
	}()

	select {
	case <-ctx.Done():
		return nil, status.Error(codes.DeadlineExceeded, "Context deadline exceeded")
	case r := <-res:
		log.Println(format.String(op, "SUCCESS"))
		return r.data, r.err
	}
}

func (s *ServerAPI) GetDictionariesByCategory(
	ctx context.Context,
	req *productsRPC.Id,
) (*productsRPC.DictionariesByCategory, error) {
	const op = "productsRPC.ServerAPI.GetDictionariesByCategory"

	type result struct {
		data *productsRPC.DictionariesByCategory
		err  error
	}
	res := make(chan result, 1)

	go func() {
		dict, err := s.API.GetDictionariesByCategory(req.GetId())
		if err != nil {
			res <- result{err: format.Error(op, status.Error(codes.Internal, err.Error()))}
			return
		}

		prDict := &productsRPC.DictionariesByCategory{
			Brands:    convert.ToBrandList(dict.Brands).(*productsRPC.BrandList),
			Countries: convert.ToCountryList(dict.Countries).(*productsRPC.CountryList),
			Materials: convert.ToMaterialList(dict.Materials).(*productsRPC.MaterialList),
			Colors:    convert.ToColorList(dict.Colors).(*productsRPC.ColorList),

			MinPrice:  int32(dict.MinPrice),
			MaxPrice:  int32(dict.MaxPrice),
			MinWidth:  int32(dict.MinWidth),
			MaxWidth:  int32(dict.MaxWidth),
			MinHeight: int32(dict.MinHeight),
			MaxHeight: int32(dict.MaxHeight),
			MinDepth:  int32(dict.MinDepth),
			MaxDepth:  int32(dict.MaxDepth),
		}

		res <- result{data: prDict}
	}()

	select {
	case <-ctx.Done():
		return nil, status.Error(codes.DeadlineExceeded, "Context deadline exceeded")
	case r := <-res:
		log.Println(format.String(op, "SUCCESS"))
		return r.data, r.err
	}
}

func (s *ServerAPI) GetPhotosByProductAndColor(ctx context.Context, req *productsRPC.ProductColorPhotosId) (*productsRPC.PhotoList, error) {
	const op = "productsRPC.GetPhotosByProductAndColor"
	log.Println(op)

	type result struct {
		data *productsRPC.PhotoList
		err  error
	}
	res := make(chan result, 1)

	go func() {
		pc, err := s.API.GetPhotosByProductAndColor(req.ProductId, req.ColorId)
		if err != nil {
			res <- result{err: format.Error(op, status.Error(codes.Internal, err.Error()))}
			return
		}

		res <- result{data: &productsRPC.PhotoList{Photos: pc}}
	}()

	select {
	case <-ctx.Done():
		return nil, status.Error(codes.DeadlineExceeded, "Context deadline exceeded")
	case r := <-res:
		log.Println(format.String(op, "SUCCESS"))
		return r.data, r.err
	}
}

func (s *ServerAPI) GetProduct(ctx context.Context, req *productsRPC.Id) (*productsRPC.Product, error) {
	const op = "productsRPC.ServerAPI.GetAllProducts"
	log.Println(op)

	type result struct {
		data *productsRPC.Product
		err  error
	}
	res := make(chan result, 1)

	go func() {
		pc, err := s.API.GetProductById(req.GetId())
		if err != nil {
			res <- result{err: format.Error(op, status.Error(codes.Internal, err.Error()))}
			return
		}

		res <- result{data: convert.ToRPCProduct(pc)}
	}()

	select {
	case <-ctx.Done():
		return nil, status.Error(codes.DeadlineExceeded, "Context deadline exceeded")
	case r := <-res:
		log.Println(format.String(op, "SUCCESS"))
		return r.data, r.err
	}
}
