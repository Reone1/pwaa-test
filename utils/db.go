package utils

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Mongodb() (*mongo.Client, context.Context, context.CancelFunc) {
	if &DATABASE_URL == nil {
		panic("env constant Not Found")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(DATABASE_URL))
	if err != nil {
			panic(err)
	}

	err = client.Ping(ctx, readpref.Primary()) // Primary DB에 대한 연결 체크
	return client, ctx, cancel
}
