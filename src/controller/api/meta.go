package api

import (
	"context"
	"strconv"
	"time"

	"github.com/CeruleanSong/gobox-server/src/database"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/valyala/fasthttp"
)

var (
	upgrader = websocket.Upgrader{}
)

// Stats a
func Stats() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		db := database.Database()
		client, err := db.Get()

		collection := client.Database("gobox").Collection("fs.metadata")
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		files, err := collection.EstimatedDocumentCount(ctx)
		if err != nil {
			return echo.ErrInternalServerError
		}

		collection = client.Database("gobox").Collection("user")
		ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
		users, err := collection.EstimatedDocumentCount(ctx)
		if err != nil {
			return echo.ErrInternalServerError
		}

		return c.JSON(fasthttp.StatusOK, map[string]string{
			"files": strconv.Itoa(int(files)),
			"users": strconv.Itoa(int(users)),
		})
	}
}

// Hello a
func Hello() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		// todo
		return c.JSON(200, map[string]string{"sucess": "true"})
	}
}
