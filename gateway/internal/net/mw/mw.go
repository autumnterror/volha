package mw

import (
	"gateway/config"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/xid"
)

func AdminAuth(cfg *config.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cookie, err := c.Cookie("admin_pw")
			if err != nil || cookie.Value != cfg.AdminPW {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "unauthorized",
				})
			}
			return next(c)
		}
	}
}

func CheckId() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			id := c.QueryParam("id")
			if id == "" {
				return next(c)
			}
			if _, err := xid.FromString(id); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{
					"error": "bad id",
				})
			}
			return next(c)
		}
	}
}
