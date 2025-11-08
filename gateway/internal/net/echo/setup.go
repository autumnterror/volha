package echo

import (
	"errors"
	"fmt"
	"gateway/config"
	"gateway/internal/grpc/products"
	"gateway/internal/net/handlers"
	"gateway/internal/net/mw"
	"gateway/internal/pkg/redis"
	"gateway/internal/utils/format"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Echo struct {
	e   *echo.Echo
	cfg *config.Config
}

const (
	ImagesDir      = "./images"
	MaxUploadBytes = 10 << 20 // 10 MB
)

func New(rds *redis.Client, a *products.Client, cfg *config.Config) *Echo {
	e := echo.New()

	h := handlers.New(a, rds, cfg)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Use(middleware.Logger(), middleware.Recover())
	e.Use(middleware.BodyLimit(fmt.Sprintf("%d", MaxUploadBytes)))
	e.Use(middleware.CORS())
	e.Static("/images", ImagesDir)

	userApi := e.Group("/api", mw.CheckId())
	{
		p := userApi.Group("/product")
		{
			p.GET("/search", h.SearchProducts)
			p.POST("/filter", h.FilterProducts)
			p.GET("/getall", h.GetAllProducts)
			p.GET("/get", h.GetProduct)
		}

		b := userApi.Group("/brand")
		{
			b.GET("/getall", h.GetAllBrands)
		}

		a := userApi.Group("/auth")
		{
			a.GET("/check", h.CheckPw)
		}

		c := userApi.Group("/category")
		{
			c.GET("/getall", h.GetAllCategories)
		}

		co := userApi.Group("/color")
		{
			co.GET("/getall", h.GetAllColors)
		}

		m := userApi.Group("/material")
		{
			m.GET("/getall", h.GetAllMaterials)
		}

		ct := userApi.Group("/country")
		{
			ct.GET("/getall", h.GetAllCountries)
		}

		dict := userApi.Group("/dictionaries")
		{
			dict.GET("/getall", h.GetAllDictionaries)
			dict.GET("/getall/category", h.GetAllDictionariesByCategory)
		}

		cp := userApi.Group("/productcolorphotos")
		{
			cp.GET("/getall", h.GetAllProductColorPhotos)
			cp.POST("/getphotos", h.GetPhotosByProductAndColor)
		}
	}

	adminApi := e.Group("/api", mw.CheckId()) //, mw.AdminAuth(cfg)) //TODO ON IF COOKIE
	{
		f := adminApi.Group("/files")
		{
			f.POST("/upload", h.UploadFile)
			f.DELETE("/delete", h.DeleteFile)
		}
		p := adminApi.Group("/product")
		{
			p.POST("/create", h.CreateProduct)
			p.PUT("/update", h.UpdateProduct)
			p.DELETE("/delete", h.DeleteProduct)
		}

		b := adminApi.Group("/brand")
		{
			b.POST("/create", h.CreateBrand)
			b.PUT("/update", h.UpdateBrand)
			b.DELETE("/delete", h.DeleteBrand)
		}

		c := adminApi.Group("/category")
		{
			c.POST("/create", h.CreateCategory)
			c.PUT("/update", h.UpdateCategory)
			c.DELETE("/delete", h.DeleteCategory)
		}

		co := adminApi.Group("/color")
		{
			co.POST("/create", h.CreateColor)
			co.PUT("/update", h.UpdateColor)
			co.DELETE("/delete", h.DeleteColor)
		}

		m := adminApi.Group("/material")
		{
			m.POST("/create", h.CreateMaterial)
			m.PUT("/update", h.UpdateMaterial)
			m.DELETE("/delete", h.DeleteMaterial)
		}

		ct := adminApi.Group("/country")
		{
			ct.POST("/create", h.CreateCountry)
			ct.PUT("/update", h.UpdateCountry)
			ct.DELETE("/delete", h.DeleteCountry)
		}
		cp := adminApi.Group("/productcolorphotos")
		{
			cp.POST("/create", h.CreateProductColorPhotos)
			cp.PUT("/update", h.UpdateProductColorPhotos)
			cp.DELETE("/delete", h.DeleteProductColorPhotos)
		}
	}

	return &Echo{
		e:   e,
		cfg: cfg,
	}
}

func (e *Echo) MustRun() {
	const op = "echo.Run"

	if err := e.e.Start(fmt.Sprintf(":%d", e.cfg.Port)); err != nil && !errors.Is(err, http.ErrServerClosed) {
		e.e.Logger.Fatal(format.Error(op, err))
	}
}

func (e *Echo) Stop() error {
	const op = "echo.Stop"

	if err := e.e.Close(); err != nil {
		return format.Error(op, err)
	}
	return nil
}
