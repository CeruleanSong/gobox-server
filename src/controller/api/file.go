package api

import (
	"bytes"
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/CeruleanSong/gobox-server/src/config"
	"github.com/CeruleanSong/gobox-server/src/database"
	"github.com/CeruleanSong/gobox-server/src/model"
	"github.com/CeruleanSong/gobox-server/src/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gabriel-vasile/mimetype"
	"github.com/labstack/echo/v4"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

// FileUpload uploads a file to the database
func FileUpload() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		/* json */
		var name string
		var token string
		var url string

		var ownerid = c.FormValue("user")
		var protectedval = c.FormValue("protected")
		var protectedflag bool

		if protectedval == "0" {
			protectedflag = false
		} else {
			protectedflag = true
		}

		var authHeader string = c.Request().Header.Get("Authorization")
		var authorization []string = strings.SplitN(authHeader, " ", 2)
		if authHeader != "" && len(authorization) == 2 {
			jwtToken, _ := jwt.ParseWithClaims(authorization[1], &model.Token{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(config.SECRET), nil
			})

			claims := jwtToken.Claims.(*model.Token) /* check if empty */
			ownerid = claims.USER

			if ownerid == "" {
				protectedflag = false
			} else {
				if protectedval == "" {
					return echo.ErrBadRequest
				}
			}
		}

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

		cTime := time.Now()

		var dbEntry = &model.FileData{
			ID:        token,
			NAME:      name,
			TYPE:      mime.String(),
			OWNERID:   ownerid,
			PROTECTED: protectedflag,
			DOWNLOADS: 0,
			VIEWS:     0,
			BYTES:     file.Size,
			UPLOADED:  cTime.Format("2006-01-02"),
			EXPIRES:   cTime.Add(time.Hour * 24 * 90).Format("2006-01-02"),
		}

		var res = &model.FileResponce{
			ID:        dbEntry.ID,
			NAME:      dbEntry.NAME,
			TYPE:      dbEntry.TYPE,
			DOWNLOADS: dbEntry.DOWNLOADS,
			VIEWS:     dbEntry.VIEWS,
			URL:       url,
			BYTES:     dbEntry.BYTES,
			UPLOADED:  dbEntry.UPLOADED,
			EXPIRES:   dbEntry.EXPIRES,
		}

		collection := client.Database("gobox").Collection("fs.metadata")
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		collection.InsertOne(ctx, dbEntry)

		return c.JSON(fasthttp.StatusOK, res)
	}
}

// FileDownload download a file
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
			return echo.ErrNotFound
		}

		collection := client.Database("gobox").Collection("fs.metadata")

		var infoResult model.FileData

		filter := bson.M{"_id": param}
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		err = collection.FindOne(ctx, filter).Decode(&infoResult)
		if err != nil {
			return err
		}

		{
			var authHeader string = c.Request().Header.Get("Authorization")
			var authorization []string = strings.SplitN(authHeader, " ", 2)
			var token *jwt.Token
			var claims *model.Token
			if authHeader != "" && len(authorization) == 2 {
				token, _ = jwt.ParseWithClaims(authorization[1], &model.Token{}, func(token *jwt.Token) (interface{}, error) {
					return []byte(config.SECRET), nil
				})

				claims = token.Claims.(*model.Token)
				claims.Valid()
			}

			if infoResult.PROTECTED == true {
				if len(authorization) == 2 {
					if !token.Valid || (claims.USER != infoResult.OWNERID) {
						return echo.ErrUnauthorized
					}
				} else {
					return echo.ErrUnauthorized
				}
			}
		}

		ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

		var result model.FileData
		filter = bson.M{"_id": bson.M{"$eq": param}}

		err = collection.FindOne(ctx, filter).Decode(&result)
		if err != nil {
			return err
		}

		update := bson.M{"$set": bson.M{"downloads": 1 + result.DOWNLOADS, "views": 1 + result.VIEWS}}
		collection.UpdateMany(ctx, filter, update)
		if err != nil {
			println("error")
			println(err.Error())
			return err
		}

		/* set proper headers */
		// println(result.DOWNLOADS)
		// fmt.Printf("%s-%d\n", result.TYPE, result.DOWNLOADS)
		c.Response().Header().Set("Content-Length", fmt.Sprintf("%x", result.BYTES))
		c.Response().Header().Set("Connection", "keep-alive")
		c.Response().Header().Set("Accept-Ranges", "bytes")
		c.Response().Header().Set("Content-Type", result.TYPE)
		c.Response().Header().Set("Content-Disposition", "inline; filename="+result.NAME)

		// return c.Blob(200, fileType, buf)
		return c.Stream(200, fileType, str)
		// return c.JSON(fasthttp.StatusOK, 0)
	}
}

// FileDelete delete a file & its metadata
func FileDelete() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		var param string = c.Param("id")

		db := database.Database()
		client, err := db.Get()

		collection := client.Database("gobox").Collection("fs.metadata")

		var infoResult model.FileData
		filter := bson.M{"_id": param}
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		err = collection.FindOne(ctx, filter).Decode(&infoResult)
		if err != nil {
			return err
		}

		{
			var authHeader string = c.Request().Header.Get("Authorization")
			var authorization []string = strings.SplitN(authHeader, " ", 2)
			var token *jwt.Token
			var claims *model.Token
			if authHeader != "" && len(authorization) == 2 {
				token, _ = jwt.ParseWithClaims(authorization[1], &model.Token{}, func(token *jwt.Token) (interface{}, error) {
					return []byte(config.SECRET), nil
				})

				claims = token.Claims.(*model.Token)
				claims.Valid()
			}

			if infoResult.PROTECTED == true {
				if len(authorization) == 2 {
					if !token.Valid || (claims.USER != infoResult.OWNERID) {
						return echo.ErrUnauthorized
					}
				} else {
					return echo.ErrUnauthorized
				}
			}
		}

		filter = bson.M{"_id": param}
		filterChunks := bson.M{"files_id": param}

		if false {
			// do nothing
		} else {
			// delete metadata
			collection := client.Database("gobox").Collection("fs.metadata")
			ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
			res, err := collection.DeleteOne(ctx, filter)
			if err != nil {
				return echo.ErrBadRequest
			}

			// delete chunks
			collection = client.Database("gobox").Collection("fs.chunks")
			ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
			res, err = collection.DeleteMany(ctx, filterChunks)
			if err != nil {
				return echo.ErrBadRequest
			}

			// delete files metadata
			collection = client.Database("gobox").Collection("fs.files")
			ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
			res, err = collection.DeleteOne(ctx, filter)
			if err != nil {
				return echo.ErrBadRequest
			}

			message := func() string {
				if err != nil {
					return err.Error()
				}
				return ""
			}()

			return c.JSON(fasthttp.StatusOK, map[string]string{
				"file":    param,
				"success": strconv.FormatBool(res.DeletedCount == 1),
				"message": message,
			})
		}

		message := func() string {
			if err != nil {
				return err.Error()
			}
			return ""
		}()

		return c.JSON(fasthttp.StatusOK, map[string]string{
			"file":    param,
			"success": "false",
			"message": message,
		})
	}
}

// FileInfo display file information
func FileInfo() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		var param string = c.Param("id")

		db := database.Database()
		client, err := db.Get()
		if err != nil {
			return err
		}

		collection := client.Database("gobox").Collection("fs.metadata")

		var infoResult model.FileData

		filter := bson.M{"_id": param}
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		err = collection.FindOne(ctx, filter).Decode(&infoResult)
		if err != nil {
			return err
		}

		{
			var authHeader string = c.Request().Header.Get("Authorization")
			var authorization []string = strings.SplitN(authHeader, " ", 2)
			var token *jwt.Token
			var claims *model.Token
			if authHeader != "" && len(authorization) == 2 {
				token, _ = jwt.ParseWithClaims(authorization[1], &model.Token{}, func(token *jwt.Token) (interface{}, error) {
					return []byte(config.SECRET), nil
				})

				claims = token.Claims.(*model.Token)
				claims.Valid()
			}

			if infoResult.PROTECTED == true {
				if len(authorization) == 2 {
					if !token.Valid || (claims.USER != infoResult.OWNERID) {
						return echo.ErrUnauthorized
					}
				} else {
					return echo.ErrUnauthorized
				}
			}
		}

		var result model.FileData

		filter = bson.M{"_id": param}
		ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
		err = collection.FindOne(ctx, filter).Decode(&result)
		if err != nil {
			return err
		}

		filter = bson.M{"_id": bson.M{"$eq": param}}
		update := bson.M{"$set": bson.M{"views": 1 + result.VIEWS}}
		collection.UpdateMany(ctx, filter, update)
		if err != nil {
			return err
		}
		return c.JSON(fasthttp.StatusOK, result)
	}
}
