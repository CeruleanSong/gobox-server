package main

import (
	"github.com/CeruleanSong/gobox-server/src/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func initialize() *echo.Echo {
	e := echo.New()

	// e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

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

	// e.Use(middleware.CORS())

	api := e.Group("/api")
	new(controller.APIController).File(api)

	return e
}

func main() {
	e := initialize()

	e.Logger.Fatal(e.Start(":1323"))
}
