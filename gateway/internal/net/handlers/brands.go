package handlers

import (
	"context"
	"gateway/internal/utils/format"
	"gateway/internal/views"
	"github.com/rs/xid"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// GetAllBrands godoc
// @Summary Получить все бренды
// @Description Возвращает список всех брендов
// @Tags brands
// @Produce json
// @Success 200 {object} views.SWGBrandListResponse
// @Failure 500 {object} views.SWGErrorResponse
// @Failure 502 {object} views.SWGErrorResponse
// @Router /api/brand/getall [get]
func (a *Apis) GetAllBrands(c echo.Context) error {
	const op = "handlers.GetAllBrands"

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	list, err := a.apiProduct.GetAllBrands(ctx)
	if err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "failed to get brands"})
	}

	if len(list) == 0 {
		list = []views.Brand{}
	}

	return c.JSON(http.StatusOK, list)
}

// CreateBrand godoc
// @Summary Создать новый бренд
// @Description Добавляет новый бренд
// @Tags brands
// @Accept json
// @Produce json
// @Param brand body views.Brand true "Данные нового бренда"
// @Success 200 {object} views.SWGSuccessResponse
// @Failure 400 {object} views.SWGErrorResponse
// @Failure 500 {object} views.SWGErrorResponse
// @Failure 502 {object} views.SWGErrorResponse
// @Router /api/brand/create [post]
func (a *Apis) CreateBrand(c echo.Context) error {
	const op = "handlers.CreateBrand"

	//delete cache dict
	if err := a.rds.CleanDictionaries(); err != nil {
		log.Println(format.Error(op, err))
	}

	var brand views.Brand
	if err := c.Bind(&brand); err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid JSON"})
	}
	brand.Id = xid.New().String()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := a.apiProduct.CreateBrand(ctx, &brand)
	if err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "failed to create brand"})
	}

	return c.JSON(http.StatusOK, map[string]string{"answer": "brand created"})
}

// UpdateBrand godoc
// @Summary Обновить бренд
// @Description Обновляет информацию о бренде по ID
// @Tags brands
// @Accept json
// @Produce json
// @Param id query string true "ID бренда"
// @Param brand body views.Brand true "Данные бренда"
// @Success 200 {object} views.SWGSuccessResponse
// @Failure 400 {object} views.SWGErrorResponse
// @Failure 500 {object} views.SWGErrorResponse
// @Failure 502 {object} views.SWGErrorResponse
// @Router /api/brand/update [put]
func (a *Apis) UpdateBrand(c echo.Context) error {
	const op = "handlers.UpdateBrand"

	//delete cache dict
	if err := a.rds.CleanDictionaries(); err != nil {
		log.Println(format.Error(op, err))
	}

	id := c.QueryParam("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "missing id"})
	}

	var brand views.Brand
	if err := c.Bind(&brand); err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid JSON"})
	}
	brand.Id = id

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := a.apiProduct.UpdateBrand(ctx, &brand)
	if err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "failed to update brand"})
	}

	return c.JSON(http.StatusOK, map[string]string{"answer": "brand updated"})
}

// DeleteBrand godoc
// @Summary Удалить бренд
// @Description Удаляет бренд по ID
// @Tags brands
// @Produce json
// @Param id query string true "ID бренда"
// @Success 200 {object} views.SWGSuccessResponse
// @Failure 400 {object} views.SWGErrorResponse
// @Failure 500 {object} views.SWGErrorResponse
// @Failure 502 {object} views.SWGErrorResponse
// @Router /api/brand/delete [delete]
func (a *Apis) DeleteBrand(c echo.Context) error {
	const op = "handlers.DeleteBrand"

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

	err := a.apiProduct.DeleteBrand(ctx, id)
	if err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "failed to delete brand"})
	}

	return c.JSON(http.StatusOK, map[string]string{"answer": "brand deleted"})
}
