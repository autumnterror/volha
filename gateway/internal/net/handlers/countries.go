package handlers

import (
	"context"
	"gateway/internal/utils/format"
	"gateway/internal/views"
	"github.com/labstack/echo/v4"
	"github.com/rs/xid"
	"log"
	"net/http"
	"time"
)

// CreateCountry godoc
// @Summary Создать страну
// @Description Добавляет новую страну
// @Tags country
// @Accept json
// @Produce json
// @Param country body views.Country true "Новая страна"
// @Success 200 {object} views.SWGSuccessResponse "Страна успешно создана"
// @Failure 400 {object} views.SWGErrorResponse "Неверный формат данных"
// @Failure 500 {object} views.SWGErrorResponse "Ошибка на сервере"
// @Failure 502 {object} views.SWGErrorResponse "Ошибка взаимодействия с сервисом"
// @Router /api/country/create [post]
func (a *Apis) CreateCountry(c echo.Context) error {
	const op = "handlers.CreateCountry"

	//delete cache dict
	if err := a.rds.CleanDictionaries(); err != nil {
		log.Println(format.Error(op, err))
	}

	var ctr views.Country
	if err := c.Bind(&ctr); err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad JSON"})
	}
	ctr.Id = xid.New().String()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := a.apiProduct.CreateCountry(ctx, &ctr); err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "could not create country"})
	}

	return c.JSON(http.StatusOK, map[string]string{"answer": "country created successfully"})
}

// UpdateCountry godoc
// @Summary Обновить страну
// @Description Обновляет данные существующей страны
// @Tags country
// @Accept json
// @Produce json
// @Param id query string true "ID страны"
// @Param country body views.Country true "Обновлённые данные страны"
// @Success 200 {object} views.SWGSuccessResponse "Страна успешно обновлена"
// @Failure 400 {object} views.SWGErrorResponse "Неверный ID или данные"
// @Failure 500 {object} views.SWGErrorResponse "Ошибка на сервере"
// @Failure 502 {object} views.SWGErrorResponse "Ошибка взаимодействия с сервисом"
// @Router /api/country/update [put]
func (a *Apis) UpdateCountry(c echo.Context) error {
	const op = "handlers.UpdateCountry"

	//delete cache dict
	if err := a.rds.CleanDictionaries(); err != nil {
		log.Println(format.Error(op, err))
	}

	id := c.QueryParam("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "missing id"})
	}

	var ctr views.Country
	if err := c.Bind(&ctr); err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad JSON"})
	}
	ctr.Id = id

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := a.apiProduct.UpdateCountry(ctx, &ctr); err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "could not update country"})
	}

	return c.JSON(http.StatusOK, map[string]string{"answer": "country updated successfully"})
}

// DeleteCountry godoc
// @Summary Удалить страну
// @Description Удаляет страну по ID
// @Tags country
// @Produce json
// @Param id query string true "ID страны"
// @Success 200 {object} views.SWGSuccessResponse "Страна успешно удалена"
// @Failure 400 {object} views.SWGErrorResponse "Неверный ID"
// @Failure 500 {object} views.SWGErrorResponse "Ошибка на сервере"
// @Failure 502 {object} views.SWGErrorResponse "Ошибка взаимодействия с сервисом"
// @Router /api/country/delete [delete]
func (a *Apis) DeleteCountry(c echo.Context) error {
	const op = "handlers.DeleteCountry"

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

	if err := a.apiProduct.DeleteCountry(ctx, id); err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "could not delete country"})
	}

	return c.JSON(http.StatusOK, map[string]string{"answer": "country deleted successfully"})
}

// GetAllCountries godoc
// @Summary Получить все страны
// @Description Возвращает список всех стран
// @Tags country
// @Produce json
// @Success 200 {object} views.SWGCountryListResponse "Успешный запрос"
// @Failure 500 {object} views.SWGErrorResponse "Ошибка на сервере"
// @Failure 502 {object} views.SWGErrorResponse "Ошибка взаимодействия с сервисом"
// @Router /api/country/getall [get]
func (a *Apis) GetAllCountries(c echo.Context) error {
	const op = "handlers.GetAllCountries"

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	list, err := a.apiProduct.GetAllCountries(ctx)
	if err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "failed to get countries"})
	}
	if len(list) == 0 {
		list = []views.Country{}
	}

	return c.JSON(http.StatusOK, list)
}
