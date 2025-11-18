package handlers

import (
	"context"
	"gateway/internal/utils/format"
	"gateway/internal/views"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/xid"
)

// CreateColor godoc
// @Summary Создать цвет
// @Description Добавляет новый цвет
// @Tags color
// @Accept json
// @Produce json
// @Param color body views.Color true "Новый цвет"
// @Success 200 {object} views.SWGSuccessResponse "Цвет успешно создан"
// @Failure 400 {object} views.SWGErrorResponse "Неверный формат данных"
// @Failure 500 {object} views.SWGErrorResponse "Ошибка на сервере"
// @Failure 502 {object} views.SWGErrorResponse "Ошибка взаимодействия с сервисом"
// @Router /api/color/create [post]
func (a *Apis) CreateColor(c echo.Context) error {
	const op = "handlers.CreateColor"

	//delete cache dict
	if err := a.rds.CleanDictionaries(); err != nil {
		log.Println(format.Error(op, err))
	}

	var clr views.Color
	if err := c.Bind(&clr); err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad JSON"})
	}
	clr.Id = xid.New().String()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := a.apiProduct.CreateColor(ctx, &clr); err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "could not create color"})
	}

	return c.JSON(http.StatusOK, map[string]string{"id": clr.Id})
}

// UpdateColor godoc
// @Summary Обновить цвет
// @Description Обновляет данные существующего цвета
// @Tags color
// @Accept json
// @Produce json
// @Param id query string true "ID цвета"
// @Param color body views.Color true "Обновлённый цвет"
// @Success 200 {object} views.SWGSuccessResponse "Цвет успешно обновлён"
// @Failure 400 {object} views.SWGErrorResponse "Неверный ID или данные"
// @Failure 500 {object} views.SWGErrorResponse "Ошибка на сервере"
// @Failure 502 {object} views.SWGErrorResponse "Ошибка взаимодействия с сервисом"
// @Router /api/color/update [put]
func (a *Apis) UpdateColor(c echo.Context) error {
	const op = "handlers.UpdateColor"

	//delete cache dict
	if err := a.rds.CleanDictionaries(); err != nil {
		log.Println(format.Error(op, err))
	}

	id := c.QueryParam("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "missing id"})
	}

	var clr views.Color
	if err := c.Bind(&clr); err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad JSON"})
	}
	clr.Id = id

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := a.apiProduct.UpdateColor(ctx, &clr); err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "could not update color"})
	}

	return c.JSON(http.StatusOK, map[string]string{"answer": "color updated successfully"})
}

// DeleteColor godoc
// @Summary Удалить цвет
// @Description Удаляет цвет по ID
// @Tags color
// @Produce json
// @Param id query string true "ID цвета"
// @Success 200 {object} views.SWGSuccessResponse "Цвет успешно удалён"
// @Failure 400 {object} views.SWGErrorResponse "Неверный ID"
// @Failure 500 {object} views.SWGErrorResponse "Ошибка на сервере"
// @Failure 502 {object} views.SWGErrorResponse "Ошибка взаимодействия с сервисом"
// @Router /api/color/delete [delete]
func (a *Apis) DeleteColor(c echo.Context) error {
	const op = "handlers.DeleteColor"

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

	if err := a.apiProduct.DeleteColor(ctx, id); err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "could not delete color"})
	}

	return c.JSON(http.StatusOK, map[string]string{"answer": "color deleted successfully"})
}

// GetAllColors godoc
// @Summary Получить все цвета
// @Description Возвращает список всех доступных цветов
// @Tags color
// @Produce json
// @Success 200 {object} views.SWGColorListResponse "Успешный запрос"
// @Failure 500 {object} views.SWGErrorResponse "Ошибка на сервере"
// @Failure 502 {object} views.SWGErrorResponse "Ошибка взаимодействия с сервисом"
// @Router /api/color/getall [get]
func (a *Apis) GetAllColors(c echo.Context) error {
	const op = "handlers.GetAllColors"

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	list, err := a.apiProduct.GetAllColors(ctx)
	if err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "failed to get colors"})
	}

	if len(list) == 0 {
		list = []views.Color{}
	}

	return c.JSON(http.StatusOK, list)
}
