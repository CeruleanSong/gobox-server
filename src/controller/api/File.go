package api

import (
	"github.com/labstack/echo/v4"
	"github.com/valyala/fasthttp"
)

// FileUpload a
func FileUpload() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		return c.JSON(fasthttp.StatusOK, 0)
	}
}

// FileDownload a
func FileDownload() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		return c.JSON(fasthttp.StatusOK, 0)
	}
}
