package handlers

import (
	"context"
	"gateway/internal/utils/format"
	"gateway/internal/views"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

// CreateProductColorPhotos godoc
// @Summary Создать фотографии продукта для цвета
// @Description Добавляет массив фотографий для конкретного продукта и цвета
// @Tags product_color_photos
// @Accept json
// @Produce json
// @Param photos body views.ProductColorPhotos true "Новый набор фотографий"
// @Success 200 {object} views.SWGSuccessResponse "Фотографии успешно добавлены"
// @Failure 400 {object} views.SWGErrorResponse "Неверный формат данных"
// @Failure 500 {object} views.SWGErrorResponse "Ошибка на сервере"
// @Failure 502 {object} views.SWGErrorResponse "Ошибка взаимодействия с сервисом"
// @Router /api/productcolorphotos/create [post]
func (a *Apis) CreateProductColorPhotos(c echo.Context) error {
	const op = "handlers.CreateProductColorPhotos"

	var pcp views.ProductColorPhotos
	if err := c.Bind(&pcp); err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad JSON"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := a.apiProduct.CreateProductColorPhotos(ctx, &pcp); err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "could not create product color photos"})
	}

	return c.JSON(http.StatusOK, map[string]string{"answer": "product color photos created successfully"})
}

// UpdateProductColorPhotos godoc
// @Summary Обновить фотографии продукта для цвета
// @Description Обновляет массив фотографий для конкретного продукта и цвета
// @Tags product_color_photos
// @Accept json
// @Produce json
// @Param photos body views.ProductColorPhotos true "Обновлённые фотографии"
// @Success 200 {object} views.SWGSuccessResponse "Фотографии успешно обновлены"
// @Failure 400 {object} views.SWGErrorResponse "Неверные данные"
// @Failure 500 {object} views.SWGErrorResponse "Ошибка на сервере"
// @Failure 502 {object} views.SWGErrorResponse "Ошибка взаимодействия с сервисом"
// @Router /api/productcolorphotos/update [put]
func (a *Apis) UpdateProductColorPhotos(c echo.Context) error {
	const op = "handlers.UpdateProductColorPhotos"

	var pcp views.ProductColorPhotos
	if err := c.Bind(&pcp); err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad JSON"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := a.apiProduct.UpdateProductColorPhotos(ctx, &pcp); err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "could not update product color photos"})
	}

	return c.JSON(http.StatusOK, map[string]string{"answer": "product color photos updated successfully"})
}

// DeleteProductColorPhotos godoc
// @Summary Удалить фотографии продукта для цвета
// @Description Удаляет записи фотографий продукта для указанного цвета
// @Tags product_color_photos
// @Accept json
// @Produce json
// @Param id body views.ProductColorPhotosId true "ID продукта и цвета"
// @Success 200 {object} views.SWGSuccessResponse "Фотографии успешно удалены"
// @Failure 400 {object} views.SWGErrorResponse "Неверные данные"
// @Failure 500 {object} views.SWGErrorResponse "Ошибка на сервере"
// @Failure 502 {object} views.SWGErrorResponse "Ошибка взаимодействия с сервисом"
// @Router /api/productcolorphotos/delete [delete]
func (a *Apis) DeleteProductColorPhotos(c echo.Context) error {
	const op = "handlers.DeleteProductColorPhotos"

	var pcpId views.ProductColorPhotosId
	if err := c.Bind(&pcpId); err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad JSON"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := a.apiProduct.DeleteProductColorPhotos(ctx, pcpId.ProductId, pcpId.ColorId); err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "could not delete product color photos"})
	}

	return c.JSON(http.StatusOK, map[string]string{"answer": "product color photos deleted successfully"})
}

// GetAllProductColorPhotos godoc
// @Summary Получить все фотографии продуктов по цветам
// @Description Возвращает все записи product_color_photos
// @Tags product_color_photos
// @Produce json
// @Success 200 {object} views.SWGProductColorPhotosListResponse "Успешный запрос"
// @Failure 500 {object} views.SWGErrorResponse "Ошибка на сервере"
// @Failure 502 {object} views.SWGErrorResponse "Ошибка взаимодействия с сервисом"
// @Router /api/productcolorphotos/getall [get]
func (a *Apis) GetAllProductColorPhotos(c echo.Context) error {
	const op = "handlers.GetAllProductColorPhotos"

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	list, err := a.apiProduct.GetAllProductColorPhotos(ctx)
	if err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "failed to get product color photos"})
	}

	if len(list) == 0 {
		list = []views.ProductColorPhotos{}
	}

	return c.JSON(http.StatusOK, views.SWGProductColorPhotosListResponse{Items: list})
}

// GetPhotosByProductAndColor godoc
// @Summary Получить фотографии продукта по ID продукта и цвета
// @Description Возвращает массив фотографий для указанного продукта и цвета
// @Tags product_color_photos
// @Accept json
// @Produce json
// @Param id body views.ProductColorPhotosId true "ID продукта и цвета"
// @Success 200 {object} []string "Список фотографий"
// @Failure 400 {object} views.SWGErrorResponse "Неверные данные"
// @Failure 500 {object} views.SWGErrorResponse "Ошибка на сервере"
// @Failure 502 {object} views.SWGErrorResponse "Ошибка взаимодействия с сервисом"
// @Router /api/productcolorphotos/getphotos [post]
func (a *Apis) GetPhotosByProductAndColor(c echo.Context) error {
	const op = "handlers.GetPhotosByProductAndColor"

	var pcpId views.ProductColorPhotosId
	if err := c.Bind(&pcpId); err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad JSON"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	photos, err := a.apiProduct.GetPhotosByProductAndColor(ctx, pcpId.ProductId, pcpId.ColorId)
	if err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "failed to get photos"})
	}

	return c.JSON(http.StatusOK, photos)
}
