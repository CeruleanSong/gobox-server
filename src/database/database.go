package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var d *db

type db struct {
	client *mongo.Client
	err    error
}

func (d *db) Get() (*mongo.Client, error) {
	return d.client, d.err
}

func create() *db {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	d := &db{
		client: client,
		err:    err,
	}

	return d
}

// Database s
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
