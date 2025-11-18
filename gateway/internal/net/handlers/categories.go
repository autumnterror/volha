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

// CreateCategory godoc
// @Summary Создать категорию
// @Description Добавляет новую категорию
// @Tags category
// @Accept json
// @Produce json
// @Param category body views.Category true "Новая категория"
// @Success 200 {object} views.SWGSuccessResponse "Категория успешно создана"
// @Failure 400 {object} views.SWGErrorResponse "Неверный формат данных"
// @Failure 500 {object} views.SWGErrorResponse "Ошибка на сервере"
// @Failure 502 {object} views.SWGErrorResponse "Ошибка взаимодействия с сервисом"
// @Router /api/category/create [post]
func (a *Apis) CreateCategory(c echo.Context) error {
	const op = "handlers.CreateCategory"

	//delete cache dict
	if err := a.rds.CleanDictionaries(); err != nil {
		log.Println(format.Error(op, err))
	}

	var cat views.Category
	if err := c.Bind(&cat); err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad JSON"})
	}
	cat.Id = xid.New().String()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := a.apiProduct.CreateCategory(ctx, &cat); err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "could not create category"})
	}

	return c.JSON(http.StatusOK, map[string]string{"id": cat.Id})
}

// UpdateCategory godoc
// @Summary Обновить категорию
// @Description Обновляет информацию о категории
// @Tags category
// @Accept json
// @Produce json
// @Param id query string true "ID категории"
// @Param category body views.Category true "Обновлённая категория"
// @Success 200 {object} views.SWGSuccessResponse "Категория успешно обновлена"
// @Failure 400 {object} views.SWGErrorResponse "Неверный ID или формат"
// @Failure 500 {object} views.SWGErrorResponse "Ошибка на сервере"
// @Failure 502 {object} views.SWGErrorResponse "Ошибка взаимодействия с сервисом"
// @Router /api/category/update [put]
func (a *Apis) UpdateCategory(c echo.Context) error {
	const op = "handlers.UpdateCategory"

	//delete cache dict
	if err := a.rds.CleanDictionaries(); err != nil {
		log.Println(format.Error(op, err))
	}

	id := c.QueryParam("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "missing id"})
	}

	var cat views.Category
	if err := c.Bind(&cat); err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad JSON"})
	}
	cat.Id = id

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := a.apiProduct.UpdateCategory(ctx, &cat); err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "could not update category"})
	}

	return c.JSON(http.StatusOK, map[string]string{"answer": "category updated successfully"})
}

// DeleteCategory godoc
// @Summary Удалить категорию
// @Description Удаляет категорию по ID
// @Tags category
// @Produce json
// @Param id query string true "ID категории"
// @Success 200 {object} views.SWGSuccessResponse "Категория успешно удалена"
// @Failure 400 {object} views.SWGErrorResponse "Неверный ID"
// @Failure 500 {object} views.SWGErrorResponse "Ошибка на сервере"
// @Failure 502 {object} views.SWGErrorResponse "Ошибка взаимодействия с сервисом"
// @Router /api/category/delete [delete]
func (a *Apis) DeleteCategory(c echo.Context) error {
	const op = "handlers.DeleteCategory"

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

	if err := a.apiProduct.DeleteCategory(ctx, id); err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "could not delete category"})
	}

	return c.JSON(http.StatusOK, map[string]string{"answer": "category deleted successfully"})
}

// GetAllCategories godoc
// @Summary Получить все категории
// @Description Возвращает список всех категорий
// @Tags category
// @Produce json
// @Success 200 {object} views.SWGCategoryListResponse "Успешный запрос"
// @Failure 500 {object} views.SWGErrorResponse "Ошибка на сервере"
// @Failure 502 {object} views.SWGErrorResponse "Ошибка взаимодействия с сервисом"
// @Router /api/category/getall [get]
func (a *Apis) GetAllCategories(c echo.Context) error {
	const op = "handlers.GetAllCategories"

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	list, err := a.apiProduct.GetAllCategories(ctx)
	if err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "failed to get categories"})
	}
	if len(list) == 0 {
		list = []views.Category{}
	}

	return c.JSON(http.StatusOK, list)
}
