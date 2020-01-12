package main

import (
	"net/http"

	"github.com/CeruleanSong/gobox-server/src/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func initialize() *echo.Echo {
	e := echo.New()

	// client := database.Database()
	// db, err := client.Get()

	// if err == nil {
	// 	collection := db.Database("GoBox").Collection("numbers")
	// 	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	// 	res, _ := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	// 	id := res.InsertedID
	// 	println("id: " + id)
	// } else {
	// 	println("oof")
	// }

	e.Use(middleware.CORS())

	api := e.Group("/api/v1")
	new(controller.APIController).File(api)

	return e
}

func main() {
	e := initialize()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
