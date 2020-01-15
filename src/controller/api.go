// Package controller s
package controller

import (
	"github.com/CeruleanSong/gobox-server/src/controller/api"
	"github.com/labstack/echo/v4"
)

// APIController s
type APIController struct{}

// File a
func (c *APIController) File(g *echo.Group) {
	{
		/* TARGET: '/api/v2/f' */
		// gFile := g.Group("")
		{
			/* upload */
			g.Any("/upload", api.FileUpload())
			/* download */
			g.Any("/download/:id", api.FileDownload())
			/* info */
			g.Any("/info/:id", api.FileInfo())
		}

		authGroup := g.Group("/auth")
		{
			authGroup.Any("/register", api.AuthRegister())
			authGroup.Any("/login", api.AuthLogin())
		}

		// gFile := g.Group("")
		{
			g.Any("/stats", api.Stats())
		}
	}
}

// // File a
// func (c *APIController) File(group *echo.Group) {
// 	//
// }
