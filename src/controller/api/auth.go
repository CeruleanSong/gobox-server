package api

import (
	"context"
	"time"

	"github.com/CeruleanSong/gobox-server/src/database"
	"github.com/CeruleanSong/gobox-server/src/model"
	"github.com/CeruleanSong/gobox-server/src/util"
	"github.com/labstack/echo/v4"
	"github.com/valyala/fasthttp"
)

// AuthRegister a
func AuthRegister() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		var email string = c.FormValue("email")
		var password string = c.FormValue("password")

		println(email)

		if email == "" || password == "" {
			return c.JSON(fasthttp.StatusOK, util.ErrorINVALIDAUTH)
		}

		passwordSlice := []byte(password)
		hash := util.Hash(passwordSlice)

		db := database.Database()
		client, err := db.Get()
		if err != nil {
			return err
		}

		collection := client.Database("gobox").Collection("user")

		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		res, err := collection.InsertOne(ctx, &model.User{EMAIL: email, PASSWORD: hash})
		if err != nil {
			return err
		}

		return c.JSON(fasthttp.StatusOK, res)
	}
}
