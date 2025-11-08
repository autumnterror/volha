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

// CreateMaterial godoc
// @Summary Создать материал
// @Description Добавляет новый материал
// @Tags material
// @Accept json
// @Produce json
// @Param material body views.Material true "Новый материал"
// @Success 200 {object} views.SWGSuccessResponse "Материал успешно создан"
// @Failure 400 {object} views.SWGErrorResponse "Неверный формат данных"
// @Failure 500 {object} views.SWGErrorResponse "Ошибка на сервере"
// @Failure 502 {object} views.SWGErrorResponse "Ошибка взаимодействия с сервисом"
// @Router /api/material/create [post]
func (a *Apis) CreateMaterial(c echo.Context) error {
	const op = "handlers.CreateMaterial"

	//delete cache dict
	if err := a.rds.CleanDictionaries(); err != nil {
		log.Println(format.Error(op, err))
	}

	var m views.Material
	if err := c.Bind(&m); err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad JSON"})
	}
	m.Id = xid.New().String()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := a.apiProduct.CreateMaterial(ctx, &m); err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "could not create material"})
	}

	return c.JSON(http.StatusOK, map[string]string{"answer": "material created successfully"})
}

// UpdateMaterial godoc
// @Summary Обновить материал
// @Description Обновляет данные существующего материала
// @Tags material
// @Accept json
// @Produce json
// @Param id query string true "ID материала"
// @Param material body views.Material true "Обновлённые данные материала"
// @Success 200 {object} views.SWGSuccessResponse "Материал успешно обновлён"
// @Failure 400 {object} views.SWGErrorResponse "Неверный ID или данные"
// @Failure 500 {object} views.SWGErrorResponse "Ошибка на сервере"
// @Failure 502 {object} views.SWGErrorResponse "Ошибка взаимодействия с сервисом"
// @Router /api/material/update [put]
func (a *Apis) UpdateMaterial(c echo.Context) error {
	const op = "handlers.UpdateMaterial"

	//delete cache dict
	if err := a.rds.CleanDictionaries(); err != nil {
		log.Println(format.Error(op, err))
	}

	id := c.QueryParam("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "missing id"})
	}

	var m views.Material
	if err := c.Bind(&m); err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad JSON"})
	}
	m.Id = id

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := a.apiProduct.UpdateMaterial(ctx, &m); err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "could not update material"})
	}

	return c.JSON(http.StatusOK, map[string]string{"answer": "material updated successfully"})
}

// DeleteMaterial godoc
// @Summary Удалить материал
// @Description Удаляет материал по ID
// @Tags material
// @Produce json
// @Param id query string true "ID материала"
// @Success 200 {object} views.SWGSuccessResponse "Материал успешно удалён"
// @Failure 400 {object} views.SWGErrorResponse "Неверный ID"
// @Failure 500 {object} views.SWGErrorResponse "Ошибка на сервере"
// @Failure 502 {object} views.SWGErrorResponse "Ошибка взаимодействия с сервисом"
// @Router /api/material/delete [delete]
func (a *Apis) DeleteMaterial(c echo.Context) error {
	const op = "handlers.DeleteMaterial"

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

	if err := a.apiProduct.DeleteMaterial(ctx, id); err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "could not delete material"})
	}

	return c.JSON(http.StatusOK, map[string]string{"answer": "material deleted successfully"})
}

// GetAllMaterials godoc
// @Summary Получить все материалы
// @Description Возвращает список всех материалов
// @Tags material
// @Produce json
// @Success 200 {object} views.SWGMaterialListResponse "Успешный запрос"
// @Failure 500 {object} views.SWGErrorResponse "Ошибка на сервере"
// @Failure 502 {object} views.SWGErrorResponse "Ошибка взаимодействия с сервисом"
// @Router /api/material/getall [get]
func (a *Apis) GetAllMaterials(c echo.Context) error {
	const op = "handlers.GetAllMaterials"

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	list, err := a.apiProduct.GetAllMaterials(ctx)
	if err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "failed to get materials"})
	}

	if len(list) == 0 {
		list = []views.Material{}
	}

	return c.JSON(http.StatusOK, list)
}
