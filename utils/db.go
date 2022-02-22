package utils

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)
var CTX *context.Context

func Mongodb() (*mongo.Client, context.CancelFunc) {
	if DATABASE_URL == "" {
		panic("env constant Not Found")
	}

	CTX, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(CTX, options.Client().ApplyURI(DATABASE_URL))
	if err != nil {
			panic(err)
	}
	
	pingErr := client.Ping(CTX, readpref.Primary()) // Primary DB에 대한 연결 체크
	
	if pingErr != nil {
		panic(pingErr)
	}
	return client, cancel
}
