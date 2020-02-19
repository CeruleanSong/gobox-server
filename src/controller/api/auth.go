package api

import (
	"context"
	"net/http"
	"time"

	"github.com/CeruleanSong/gobox-server/src/config"
	"github.com/CeruleanSong/gobox-server/src/database"
	"github.com/CeruleanSong/gobox-server/src/model"
	"github.com/CeruleanSong/gobox-server/src/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

// AuthRegister register a new user
func AuthRegister() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		user := c.FormValue("user")
		password := c.FormValue("password")
		passwordSlice := []byte(password)

		/* check if empty */
		if user == "" || password == "" {
			return echo.ErrUnauthorized
		}

		hash := util.Hash(passwordSlice)

		/* get refernece to database */
		db := database.Database()
		client, err := db.Get()
		if err != nil {
			return err
		}

		/* add user to database */
		collection := client.Database("gobox").Collection("user")
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		_, err = collection.InsertOne(ctx, &model.User{USER: user, PASSWORD: hash})
		if err != nil {
			return echo.ErrUnauthorized
		}

		/* generate token */
		claims := model.Token{}

		claims.USER = user
		claims.ADMIN = false
		claims.ExpiresAt = time.Now().Add(time.Hour * 720).Unix()
		claims.Issuer = "gobox"

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		st, err := token.SignedString([]byte(config.SECRET))
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, map[string]string{
			"user":  user,
			"token": st,
		})
	}
}

// AuthLogin get an authorization token for a particular user
func AuthLogin() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		user := c.FormValue("user")
		password := c.FormValue("password")
		passwordSlice := []byte(password)

		/* check if empty */
		if user == "" || password == "" {
			return echo.ErrUnauthorized
		}

		/* get refernece to database */
		db := database.Database()
		client, err := db.Get()
		if err != nil {
			return err
		}

		/* check user is in database */
		var result model.User
		collection := client.Database("gobox").Collection("user")
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		res := collection.FindOne(ctx, &bson.M{"_id": user})

		/* decode user */
		res.Decode(&result)
		if err != nil {
			return echo.ErrBadRequest
		}

		/* match provided password with database */
		if util.Compare(result.PASSWORD, passwordSlice) {
			/* generate token */
			claims := model.Token{}

			claims.USER = user
			claims.ADMIN = false
			claims.ExpiresAt = time.Now().Add(time.Hour * 720).Unix()
			claims.Issuer = "gobox"

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

			st, err := token.SignedString([]byte(config.SECRET))
			if err != nil {
				return err
			}
			return c.JSON(http.StatusOK, map[string]string{
				"user":  user,
				"token": st,
			})
		}

		return echo.ErrUnauthorized
	}
}
