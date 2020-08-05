package staticsloth

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// NewBlockMiddleware instantiates a new Block middleware with the given paths that will be blocked
func NewBlockMiddleware(blockPaths []string) *BlockMiddleware {
	m := &BlockMiddleware{
		blockPaths: blockPaths,
	}
	return m
}

// BlockMiddleware is a middleware for blocking paths.
type BlockMiddleware struct {
	blockPaths []string
}

// Handle middleware blocks requests that match with one of the paths.
func (m *BlockMiddleware) Handle(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		path := c.Request().URL.Path
		if isPathMatch(path, m.blockPaths) {
			return c.String(http.StatusForbidden, "Forbidden")
		}

		return next(c)
	}
}
