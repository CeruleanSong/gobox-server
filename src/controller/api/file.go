package api

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/CeruleanSong/gobox-server/src/database"
	util "github.com/CeruleanSong/gobox-server/src/lib"
	"github.com/CeruleanSong/gobox-server/src/model"
	"github.com/h2non/filetype"
	"github.com/labstack/echo/v4"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson"
)

// FileUpload a
func FileUpload() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		/* var */
		var dir string
		/* json */
		var name string
		var token string
		var url string

		/* generate random hash for token */
		g, nil := util.GenerateRandomBytes(4) // generate data for token
		token = fmt.Sprintf("%x", g)          // create token
		dir = "./data"                        // get directory

		/* get the file from the request */
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}
		data, err := file.Open()
		if err != nil {
			return err
		}
		defer data.Close()

		/* file data */
		name = file.Filename // get file name

		url = "http://localhost:1323" + "/api/v2/file/download/" + token // create file url

		/* create directory */
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			err = os.MkdirAll(dir, 0755)
			if err != nil {
				return err
			}
		}

		/* create file */
		out, err := os.OpenFile(dir+"/"+token, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			return err
		}
		defer out.Close()

		_, err = io.Copy(out, data)
		if err != nil {
			return nil
		}

		bytes, err := ioutil.ReadAll(data)
		fmt.Printf("size:%d", len(bytes))

		var f = &model.File{
			NAME:  name,
			TOKEN: token,
			URL:   url,
		}

		client := database.Database()
		db, err := client.Get()

		if err == nil {
			collection := db.Database("gobox").Collection("file")
			ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
			res, _ := collection.InsertOne(ctx, f)
			id := res.InsertedID
			println("id: %s", id)
		} else {
			println("oof")
		}

		return c.JSON(fasthttp.StatusOK, f)
	}
}

// FileDownload a
func FileDownload() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		/* var */
		var dir string = "./data"
		/* json */
		var name string
		// var token string
		// var url string
		var fileType string
		var param string = c.Param("id")

		/* check data folder exists */
		file, err := os.Open(dir + "/" + param)
		if err != nil {
			return err
		}

		buf, _ := ioutil.ReadFile(file.Name())
		kind, err := filetype.Match(buf)
		if err != nil {
			fmt.Printf("err: %s", err)
		}

		fileType = kind.MIME.Value
		client := database.Database()
		db, err := client.Get()
		if err != nil {
			return err
		}

		var result model.File

		collection := db.Database("gobox").Collection("file")

		path := file.Name()
		rootName := filepath.Base(path)

		filter := bson.M{"_id": rootName}
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		err = collection.FindOne(ctx, filter).Decode(&result)
		if err != nil {
			println(fileType)
			return err
		}
		name = result.NAME

		info, err := file.Stat()
		c.Response().Header().Set("Content-Length", fmt.Sprintf("%d", info.Size()))
		c.Response().Header().Set("Connection", "keep-alive")
		c.Response().Header().Set("Accept-Ranges", "bytes")
		c.Response().Header().Set("Content-Disposition", "inline; filename="+name)

		// return c.Blob(200, fileType, buf)
		return c.Stream(200, fileType, (*os.File)(file))
		// return c.JSON(fasthttp.StatusOK, 0)
	}
}
