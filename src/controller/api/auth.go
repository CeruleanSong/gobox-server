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

// AuthRegister a
func AuthRegister() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		email := c.FormValue("email")
		password := c.FormValue("password")
		passwordSlice := []byte(password)

		/* check if empty */
		if email == "" || password == "" {
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
		_, err = collection.InsertOne(ctx, &model.User{EMAIL: email, PASSWORD: hash})
		if err != nil {
			return echo.ErrUnauthorized
		}

		/* generate token */
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["email"] = email
		claims["admin"] = false
		claims["exp"] = time.Now().Add(time.Hour * 720).Unix() // 30 days

		st, err := token.SignedString([]byte(config.SECRET))
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, map[string]string{
			"email": email,
			"token": st,
		})
	}
}

// AuthLogin a
func AuthLogin() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		email := c.FormValue("email")
		password := c.FormValue("password")
		passwordSlice := []byte(password)

		/* check if empty */
		if email == "" || password == "" {
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
		res := collection.FindOne(ctx, &bson.M{"_id": email})

		/* decode user */
		res.Decode(&result)
		if err != nil {
			return echo.ErrBadRequest
		}

		/* match provided password with database */
		if util.Compare(result.PASSWORD, passwordSlice) {
			/* generate token */
			token := jwt.New(jwt.SigningMethodHS256)
			claims := token.Claims.(jwt.MapClaims)
			claims["email"] = email
			claims["admin"] = false
			claims["exp"] = time.Now().Add(time.Hour * 720).Unix() // 30 days

			st, err := token.SignedString([]byte(config.SECRET))
			if err != nil {
				return err
			}
			return c.JSON(http.StatusOK, map[string]string{
				"email": email,
				"token": st,
			})
		}

		return echo.ErrUnauthorized
	}
}

// AuthVerify a
// func AuthVerify() echo.HandlerFunc {
// 	return func(c echo.Context) (err error) {
// 		var header string = c.Request().Header.Get("Authorization")
// 		var email string = c.FormValue("email")
// 		var password string = c.FormValue("password")
// 		passwordSlice := []byte(password)

// 		var authorization = strings.Split(header, " ")

// 		token, err := jwt.Parse(authorization[1], func(token *jwt.Token) (interface{}, error) {
// 			// Don't forget to validate the alg is what you expect:
// 			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 			}

// 			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
// 			return config.SECRET, nil
// 		})

// 		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 			fmt.Println(claims["foo"], claims["nbf"])
// 			db := database.Database()
// 			client, err := db.Get()
// 			if err != nil {
// 				return err
// 			}

// 			var result model.User

// 			collection := client.Database("gobox").Collection("user")
// 			ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

// 			filter := bson.M{"_id": email}

// 			res := collection.FindOne(ctx, &filter)

// 			res.Decode(&result)
// 			if err != nil {
// 				return err
// 			}

// 			if util.Valid(result.PASSWORD, passwordSlice) {
// 				// return c.Echo().Use(middleware.JWTConfig())
// 				return c.JSON(fasthttp.StatusOK, util.ErrorINVALIDAUTH)
// 			}

// 		} else {
// 			fmt.Println(err)
// 		}

// 		return c.JSON(fasthttp.StatusOK, util.ErrorINVALIDAUTH)
// 	}
// }
