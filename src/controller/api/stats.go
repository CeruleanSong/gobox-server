package api

import (
	"context"
	"time"

	"github.com/CeruleanSong/gobox-server/src/database"
	"github.com/labstack/echo/v4"
	"github.com/valyala/fasthttp"
)

// Stats a
func Stats() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		db := database.Database()
		client, err := db.Get()
		collection := client.Database("gobox").Collection("metadata")

		// filter := bson.M{"_id": param}
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		res, err := collection.EstimatedDocumentCount(ctx)
		if err != nil {
			return err
		}

		print(res)

		return c.JSON(fasthttp.StatusOK, res)
	}
}
