package main

import (
	"github.com/CeruleanSong/gobox-server/src/config"
	"github.com/CeruleanSong/gobox-server/src/controller"
	"github.com/CeruleanSong/gobox-server/src/controller/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func initialize() *echo.Echo {
	e := echo.New()

	// e.Use(middleware.Logger())
	// e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	// e.Use(middleware.JWT([]byte(config.SECRET)))

	apiRoute := e.Group("/api")
	new(controller.APIController).File(apiRoute)

	{
		e.Any("/download/:id", api.FileDownload())
	}

	return e
}

func main() {
	e := initialize()

	e.Logger.Fatal(e.Start(config.PORT))
}
