package staticsloth

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
)

// NewCacheControlMiddleware instantiates a new CacheControl middleware with:
//   the given paths to apply the cache control to
//   and the give 'max-age' duration
func NewCacheControlMiddleware(cacheControlPaths []string, cacheControlDuration int) *CacheControlMiddleware {
	m := &CacheControlMiddleware{
		cacheControlPaths:       cacheControlPaths,
		cacheControlHeaderValue: fmt.Sprintf("max-age=%d", cacheControlDuration),
	}
	return m
}

// CacheControlMiddleware is a middleware for the Cache-Control header.
type CacheControlMiddleware struct {
	cacheControlPaths       []string
	cacheControlHeaderValue string
}

// Handle middleware adds a `CacheControl` header to the response.
func (m *CacheControlMiddleware) Handle(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		path := c.Request().URL.Path
		if isPathMatch(path, m.cacheControlPaths) {
			c.Response().Header().Set("Cache-Control", m.cacheControlHeaderValue)
		}

		return next(c)
	}
}

func isPathMatch(path string, cacheControlPaths []string) bool {

	for _, p := range cacheControlPaths {
		if strings.HasPrefix(path, p) {
			return true
		}
	}
	return false
}
