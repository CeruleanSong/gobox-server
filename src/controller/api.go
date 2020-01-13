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
		gFile := g.Group("/file")

		{
			/* upload */
			gFile.Any("/upload", api.FileUpload())
			/* download */
			gFile.Any("/download/:id", api.FileDownload())
		}
	}
}

// // File a
// func (c *APIController) File(group *echo.Group) {
// 	//
// }
