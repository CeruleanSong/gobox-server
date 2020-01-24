// Package controller s
package controller

import (
	"github.com/CeruleanSong/gobox-server/src/controller/api"
	"github.com/labstack/echo/v4"
)

// APIController s
type APIController struct{}

// File collection of routes for files
func (c *APIController) File(g *echo.Group) {
	/********************** file **********************/

	{
		/* upload */
		g.Any("/upload", api.FileUpload())
		/* download */
		g.Any("/download/:id", api.FileDownload())
		/* delete file */
		g.Any("/delete/:id", api.FileDelete())
		/* info */
		g.Any("/info/:id", api.FileInfo())
	}
}

// Auth collection of routes relating to authentication
func (c *APIController) Auth(g *echo.Group) {
	/********************** auth **********************/

	authGroup := g.Group("/auth")
	{
		authGroup.Any("/register", api.AuthRegister())
		authGroup.Any("/login", api.AuthLogin())
	}
}

// Meta collection of routes regarding miscellaneous actions
func (c *APIController) Meta(g *echo.Group) {
	/********************** meta **********************/

	metaGroup := g.Group("/meta")
	{
		metaGroup.Any("/stats", api.Stats())
		metaGroup.Any("/hello", api.Hello())
	}
}
