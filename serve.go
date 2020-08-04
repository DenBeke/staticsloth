package staticsloth

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoLog "github.com/labstack/gommon/log"
	logMiddleware "github.com/neko-neko/echo-logrus/v2"
	nekoLog "github.com/neko-neko/echo-logrus/v2/log"
)

// Serve craftboard server
func Serve(config *Config) {

	// prep echo
	e := echo.New()

	e.HideBanner = true

	// Gzip
	if config.Gzip {
		e.Use(middleware.Gzip())
	}

	// Cache Control
	e.Use(ServerHeaderMiddleware)

	// Cache Control
	if len(config.CacheControlPaths) > 0 {
		e.Use(NewCacheControlMiddleware(config.CacheControlPaths, config.CacheControlDuration).Handle)
	}

	// request logging
	if config.AccessLog {
		nekoLog.Logger().SetOutput(os.Stdout)
		nekoLog.Logger().SetLevel(echoLog.INFO)
		e.Logger = nekoLog.Logger()
		e.Use(logMiddleware.Logger())
	}

	// Real magic is happening here 🦥
	e.Static(config.PathPrefix, config.Directory)

	e.Logger.Fatal(e.Start(config.HTTPAddress))
}
