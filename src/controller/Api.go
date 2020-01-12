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
		g.Any("/file", api.FileUpload())
	}
}

// // File a
// func (c *APIController) File(group *echo.Group) {
// 	//
// }
