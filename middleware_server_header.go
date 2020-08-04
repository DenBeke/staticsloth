package staticsloth

import (
	"github.com/labstack/echo/v4"
)

// ServerHeaderMiddleware middleware adds a `Server` header to the response.
func ServerHeaderMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "StaticSloth")
		return next(c)
	}
}
