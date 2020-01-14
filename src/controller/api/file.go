package api

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/CeruleanSong/gobox-server/src/config"
	"github.com/CeruleanSong/gobox-server/src/database"
	"github.com/CeruleanSong/gobox-server/src/model"
	"github.com/CeruleanSong/gobox-server/src/util"
	"github.com/gabriel-vasile/mimetype"
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
		name = file.Filename                                      // get file name
		url = config.URLROOT + config.PORT + "/download/" + token // create file url

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

		// kind, _ := filetype.Get(fileBytes)
		mime := mimetype.Detect(fileBytes)

		var dbEntry = &model.FileData{
			NAME:     name,
			ID:       token,
			BYTES:    file.Size,
			TYPE:     mime.String(),
			UPLOADED: time.Now(),
			// EXPIRES:  time.Now().Add(time.Hour * 24 * 90),
		}

		var res = &model.FileResponce{
			NAME:     name,
			ID:       token,
			URL:      url,
			BYTES:    file.Size,
			TYPE:     mime.String(),
			UPLOADED: time.Now(),
			EXPIRES:  time.Now().Add(time.Hour * 24 * 90),
		}

		collection := client.Database("gobox").Collection("fs.metadata")
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		collection.InsertOne(ctx, dbEntry)

		return c.JSON(fasthttp.StatusOK, res)
	}
}

// FileDownload a
func FileDownload() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		var fileType string
		var param string = c.Param("id")

		db := database.Database()
		client, err := db.Get()
		bucket, err := gridfs.NewBucket(client.Database("gobox"))
		if err != nil {
			return err
		}

		// Upload the file into the database
		str, err := bucket.OpenDownloadStream(param)
		if err != nil {
			return c.JSON(fasthttp.StatusOK, util.ErrorFILENOTFOUND)
		}

		collection := client.Database("gobox").Collection("fs.metadata")
		var result model.FileData

		filter := bson.M{"_id": param}
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		err = collection.FindOne(ctx, filter).Decode(&result)
		if err != nil {
			return err
		}

		/* set proper headers */
		c.Response().Header().Set("Content-Length", fmt.Sprintf("%d", result.BYTES))
		c.Response().Header().Set("Connection", "keep-alive")
		c.Response().Header().Set("Accept-Ranges", "bytes")
		c.Response().Header().Set("Content-Type", result.TYPE)
		c.Response().Header().Set("Content-Disposition", "inline; filename="+result.NAME)

		// return c.Blob(200, fileType, buf)
		return c.Stream(200, fileType, str)
		// return c.JSON(fasthttp.StatusOK, 0)
	}
}

// FileInfo a
func FileInfo() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		var param string = c.Param("id")

		db := database.Database()
		client, err := db.Get()
		if err != nil {
			return err
		}

		collection := client.Database("gobox").Collection("fs.metadata")
		var result model.FileData

		filter := bson.M{"_id": param}
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		err = collection.FindOne(ctx, filter).Decode(&result)
		if err != nil {
			return err
		}

		return c.JSON(fasthttp.StatusOK, result)
	}
}
