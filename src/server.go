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

	/********************** middleware **********************/

	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	/********************** routes **********************/

	// add routes for api
	apiRoute := e.Group("/api")
	new(controller.APIController).File(apiRoute)

	// top level shortcuts
	{
		/** api **/
		e.Any("/download/:id", api.FileDownload())
		e.Any("/info/:id", api.FileInfo())
	}

	return e
}

func main() {
	e := initialize()

	e.Logger.Fatal(e.Start(config.PORT))
}
