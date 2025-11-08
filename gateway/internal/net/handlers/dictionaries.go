package handlers

import (
	"context"
	"gateway/internal/utils/format"
	"gateway/internal/views"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetAllDictionariesByCategory godoc
// @Summary Получить все справочники по категории
// @Description Возвращает список всех доступных справочных данных: бренды, материалы, страны и цвета и размеры по определенной категории
// @Tags dictionaries
// @Produce json
// @Param id query string true "ID категории"
// @Success 200 {object} views.Dictionaries "Успешный запрос"
// @Failure 500 {object} views.SWGErrorResponse "Ошибка на сервере"
// @Failure 502 {object} views.SWGErrorResponse "Ошибка взаимодействия с сервисом"
// @Router /api/dictionaries/getall/category [get]
func (a *Apis) GetAllDictionariesByCategory(c echo.Context) error {
	const op = "handlers.GetAllDictionariesByCategory"

	id := c.QueryParam("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "missing id"})
	}

	//check cache
	if ds, err := a.rds.GetDictionariesByCategory(id); err == nil {
		log.Println("read dict from cache", ds)
		return c.JSON(http.StatusOK, ds)
	}
	log.Println("no dict on cache")

	//if not in cache get
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	data, err := a.apiProduct.GetAllDictionariesByCategory(ctx, id)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			switch st.Code() {
			case codes.Internal:
				log.Println(format.Error(op, err))
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal service error"})
			default:
				log.Println(format.Error(op, err))
				return c.JSON(http.StatusBadGateway, map[string]string{"error": "failed to get dictionaries"})
			}
		}
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "gateway error"})
	}

	if len(data.Materials) == 0 {
		data.Materials = []views.Material{}
	}
	if len(data.Colors) == 0 {
		data.Colors = []views.Color{}
	}
	if len(data.Brands) == 0 {
		data.Brands = []views.Brand{}
	}
	if len(data.Countries) == 0 {
		data.Countries = []views.Country{}
	}
	//set in cache
	if err := a.rds.SetDictionariesByCategory(id, data); err != nil {
		log.Println(format.Error(op, err))
	}
	return c.JSON(http.StatusOK, data)
}

// GetAllDictionaries godoc
// @Summary Получить все справочники
// @Description Возвращает список всех доступных справочных данных: бренды, категории, материалы, страны и цвета
// @Tags dictionaries
// @Produce json
// @Success 200 {object} views.Dictionaries "Успешный запрос"
// @Failure 500 {object} views.SWGErrorResponse "Ошибка на сервере"
// @Failure 502 {object} views.SWGErrorResponse "Ошибка взаимодействия с сервисом"
// @Router /api/dictionaries/getall [get]
func (a *Apis) GetAllDictionaries(c echo.Context) error {
	const op = "handlers.GetAllDictionaries"
	//check cache
	if ds, err := a.rds.GetDictionaries(); err == nil {
		log.Println("read dict from cache", ds)
		return c.JSON(http.StatusOK, ds)
	}
	log.Println("no dict on cache")

	//if not in cache get
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	data, err := a.apiProduct.GetAllDictionaries(ctx)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			switch st.Code() {
			case codes.Internal:
				log.Println(format.Error(op, err))
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal service error"})
			default:
				log.Println(format.Error(op, err))
				return c.JSON(http.StatusBadGateway, map[string]string{"error": "failed to get dictionaries"})
			}
		}
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "gateway error"})
	}

	if len(data.Categories) == 0 {
		data.Categories = []views.Category{}
	}
	if len(data.Materials) == 0 {
		data.Materials = []views.Material{}
	}
	if len(data.Colors) == 0 {
		data.Colors = []views.Color{}
	}
	if len(data.Brands) == 0 {
		data.Brands = []views.Brand{}
	}
	if len(data.Countries) == 0 {
		data.Countries = []views.Country{}
	}
	//set in cache
	if err := a.rds.SetDictionaries(data); err != nil {
		log.Println(format.Error(op, err))
	}
	return c.JSON(http.StatusOK, data)
}
