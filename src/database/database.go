package database

import (
	"context"
	"time"

	"github.com/CeruleanSong/gobox-server/src/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var d *db

type db struct {
	client *mongo.Client
	err    error
}

// Get grab a refrence to the mongo singleton
func (d *db) Get() (*mongo.Client, error) {
	return d.client, d.err
}

// create creates a singleton refrence to a mongo instance
func create() *db {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.URLDB))

	d := &db{
		client: client,
		err:    err,
	}

	return d
}

// Database returns a refrence to the mongo database
func Database() *db {
	if d == nil {
		client := create()
		d = &db{
			client: client.client,
			err:    client.err,
		}
	}
	return d
}
