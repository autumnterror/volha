package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

// CheckPw godoc
// @Summary Проверить пароль
// @Description проверяет пароль из cookie
// @Tags auth
// @Produce json
// @Success 200 {object} views.SWGSuccessResponse "Успешный запрос"
// @Failure 500 {object} views.SWGErrorResponse "Ошибка на сервере"
// @Failure 502 {object} views.SWGErrorResponse "Ошибка взаимодействия с сервисом"
// @Router /api/auth/check [get]
func (a *Apis) CheckPw(c echo.Context) error {
	const op = "handlers.CheckPw"
	log.Println(op)

	cookie, err := c.Cookie("admin_pw")

	if err != nil || cookie.Value != a.cfg.AdminPW {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "unauthorized",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"answer": "pw cool",
	})
}
