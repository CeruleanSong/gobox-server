package api

import (
	"bytes"
	"context"
	"fmt"
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
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

// FileUpload a
func FileUpload() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		/* json */
		var name string
		var token string
		var url string

		/* generate random hash for token */
		g, nil := util.GenerateRandomBytes(4) // generate data for token
		token = fmt.Sprintf("%x", g)          // create token

		/* get the file from the request */
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}
		data, err := file.Open()
		if err != nil {
			return err
		}
		fileBytes := make([]byte, file.Size)
		data.Read(fileBytes)
		defer data.Close()

		/* file data */
		name = file.Filename           // get file name
		url = "/api/download/" + token // create file url

		// Create a connection to database & collection
		db := database.Database()
		client, err := db.Get()
		bucket, err := gridfs.NewBucket(client.Database("gobox"))
		if err != nil {
			return err
		}

		// Upload the file into the database
		err = bucket.UploadFromStreamWithID(token, name, bytes.NewReader(fileBytes))
		if err != nil {
			return err
		}

		kind, _ := filetype.Get(fileBytes)

		var dbEntry = &model.FileData{
			NAME:     name,
			ID:       token,
			BYTES:    file.Size,
			TYPE:     kind.MIME.Value,
			UPLOADED: time.Now(),
			EXPIRES:  time.Now().Add(time.Hour * 24 * 90),
		}

		var res = &model.FileResponce{
			NAME:     name,
			ID:       token,
			URL:      url,
			BYTES:    file.Size,
			TYPE:     kind.MIME.Value,
			UPLOADED: time.Now(),
			EXPIRES:  time.Now().Add(time.Hour * 24 * 90),
		}

		if err == nil {
			collection := client.Database("gobox").Collection("metadata")
			ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
			collection.InsertOne(ctx, dbEntry)
		} else {
			println("oof")
		}

		return c.JSON(fasthttp.StatusOK, res)
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

		var result model.FileData

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
