package handlers

import (
	"context"
	"gateway/config"
	"gateway/internal/grpc/products"
	"gateway/internal/pkg/redis"
	"gateway/internal/utils/format"
	"gateway/internal/views"
	"log"
	"net/http"
	"time"

	"github.com/rs/xid"

	"github.com/labstack/echo/v4"
)

type Apis struct {
	apiProduct *products.Client
	rds        *redis.Client
	cfg        *config.Config
}

func New(
	apiProduct *products.Client,
	rds *redis.Client,
	cfg *config.Config,
) *Apis {
	return &Apis{
		rds:        rds,
		apiProduct: apiProduct,
		cfg:        cfg,
	}
}

func classifyQuery(q string) *views.ProductSearch {
	if _, err := xid.FromBytes([]byte(q)); err == nil {
		return &views.ProductSearch{Id: q}
	}

	if len(q) == 8 && isDigitsOnly(q) {
		return &views.ProductSearch{Article: q}
	}

	return &views.ProductSearch{Title: q}
}

func isDigitsOnly(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

// SearchProducts godoc
// @Summary Поиск продуктов
// @Description Ищет продукт по ID (UUID), article (8 цифр) или названию (частичное совпадение)
// @Tags product
// @Produce json
// @Param query query string true "Поисковый запрос"
// @Success 200 {object} []views.Product
// @Failure 400 {object} views.SWGErrorResponse
// @Failure 502 {object} views.SWGErrorResponse
// @Router /api/product/search [get]
func (a *Apis) SearchProducts(c echo.Context) error {
	const op = "handlers.SearchProducts"

	query := c.QueryParam("query")
	if query == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "empty query"})
	}

	filter := classifyQuery(query)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	list, err := a.apiProduct.SearchProducts(ctx, filter)
	if err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "could not search products"})
	}

	if len(list) == 0 {
		list = []views.Product{}
	}

	return c.JSON(http.StatusOK, list)
}

// CreateProduct godoc
// @Summary Создать продукт
// @Description Добавляет новый продукт
// @Tags product
// @Accept json
// @Produce json
// @Param product body views.ProductId true "Новый продукт"
// @Success 200 {object} views.SWGIdResponse "Продукт успешно создан"
// @Failure 400 {object} views.SWGErrorResponse "Неверный формат данных"
// @Failure 500 {object} views.SWGErrorResponse "Ошибка на сервере"
// @Failure 502 {object} views.SWGErrorResponse "Ошибка взаимодействия с сервисом"
// @Router /api/product/create [post]
func (a *Apis) CreateProduct(c echo.Context) error {
	const op = "handlers.CreateProduct"

	//delete cache dict
	if err := a.rds.CleanDictionaries(); err != nil {
		log.Println(format.Error(op, err))
	}

	var p views.ProductId
	if err := c.Bind(&p); err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad JSON"})
	}
	if len(p.Article) != 8 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "article must be 8 digits"})
	}

	p.Id = xid.New().String()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := a.apiProduct.CreateProduct(ctx, &p); err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "could not create product"})
	}

	return c.JSON(http.StatusOK, map[string]string{"id": p.Id})
}

// UpdateProduct godoc
// @Summary Обновить продукт
// @Description Обновляет данные существующего продукта
// @Tags product
// @Accept json
// @Produce json
// @Param id query string true "ID продукта"
// @Param product body views.ProductId true "Обновлённые данные продукта"
// @Success 200 {object} views.SWGSuccessResponse "Продукт успешно обновлён"
// @Failure 400 {object} views.SWGErrorResponse "Неверный ID или данные"
// @Failure 500 {object} views.SWGErrorResponse "Ошибка на сервере"
// @Failure 502 {object} views.SWGErrorResponse "Ошибка взаимодействия с сервисом"
// @Router /api/product/update [put]
func (a *Apis) UpdateProduct(c echo.Context) error {
	const op = "handlers.UpdateProduct"

	//delete cache dict
	if err := a.rds.CleanDictionaries(); err != nil {
		log.Println(format.Error(op, err))
	}

	id := c.QueryParam("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "missing id"})
	}

	var p views.ProductId
	if err := c.Bind(&p); err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad JSON"})
	}
	p.Id = id

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := a.apiProduct.UpdateProduct(ctx, &p); err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "could not update product"})
	}

	return c.JSON(http.StatusOK, map[string]string{"answer": "product updated successfully"})
}

// DeleteProduct godoc
// @Summary Удалить продукт
// @Description Удаляет продукт по ID
// @Tags product
// @Produce json
// @Param id query string true "ID продукта"
// @Success 200 {object} views.SWGSuccessResponse "Продукт успешно удалён"
// @Failure 400 {object} views.SWGErrorResponse "Неверный ID"
// @Failure 500 {object} views.SWGErrorResponse "Ошибка на сервере"
// @Failure 502 {object} views.SWGErrorResponse "Ошибка взаимодействия с сервисом"
// @Router /api/product/delete [delete]
func (a *Apis) DeleteProduct(c echo.Context) error {
	const op = "handlers.DeleteProduct"

	//delete cache dict
	if err := a.rds.CleanDictionaries(); err != nil {
		log.Println(format.Error(op, err))
	}

	id := c.QueryParam("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "missing id"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := a.apiProduct.DeleteProduct(ctx, id); err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "could not delete product"})
	}

	return c.JSON(http.StatusOK, map[string]string{"answer": "product deleted successfully"})
}

// GetAllProducts godoc
// @Summary Получить все продукты
// @Description Возвращает список всех продуктов
// @Tags product
// @Produce json
// @Success 200 {object} views.SWGProductListResponse "Успешный запрос"
// @Failure 500 {object} views.SWGErrorResponse "Ошибка на сервере"
// @Failure 502 {object} views.SWGErrorResponse "Ошибка взаимодействия с сервисом"
// @Router /api/product/getall [get]
func (a *Apis) GetAllProducts(c echo.Context) error {
	const op = "handlers.GetAllProducts"

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	list, err := a.apiProduct.GetAllProducts(ctx)
	if err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "could not fetch products"})
	}

	if len(list) == 0 {
		return c.JSON(http.StatusOK, []views.Product{})
	}

	return c.JSON(http.StatusOK, list)
}

// GetProduct godoc
// @Summary Получить продукт по id
// @Description Возвращает продукт
// @Tags product
// @Produce json
// @Param id query string true "ID продукта"
// @Success 200 {object} views.Product "Успешный запрос"
// @Failure 500 {object} views.SWGErrorResponse "Ошибка на сервере"
// @Failure 502 {object} views.SWGErrorResponse "Ошибка взаимодействия с сервисом"
// @Router /api/product/get [get]
func (a *Apis) GetProduct(c echo.Context) error {
	const op = "handlers.GetProduct"

	id := c.QueryParam("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "missing id"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	pr, err := a.apiProduct.GetProduct(ctx, id)
	if err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "could not fetch products"})
	}

	return c.JSON(http.StatusOK, pr)
}

// FilterProducts godoc
// @Summary Получить продукты по фильтру
// @Description Возвращает список продуктов, соответствующих заданным критериям фильтрации
// @Tags product
// @Accept json
// @Produce json
// @Param filter body views.ProductFilter true "Параметры фильтрации"
// @Success 200 {object} views.SWGProductListResponse "Успешный запрос"
// @Failure 400 {object} views.SWGErrorResponse "Неверный формат запроса"
// @Failure 500 {object} views.SWGErrorResponse "Ошибка на сервере"
// @Failure 502 {object} views.SWGErrorResponse "Ошибка взаимодействия с сервисом"
// @Router /api/product/filter [post]
func (a *Apis) FilterProducts(c echo.Context) error {
	const op = "handlers.FilterProducts"

	var f views.ProductFilter
	if err := c.Bind(&f); err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad JSON"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	list, err := a.apiProduct.FilterProducts(ctx, &f)
	if err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "could not filter products"})
	}

	if len(list) == 0 {
		return c.JSON(http.StatusOK, []views.Product{})
	}

	return c.JSON(http.StatusOK, list)
}
