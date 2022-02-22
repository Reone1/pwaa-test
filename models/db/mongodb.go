package db

import (
	"log"

	"github.com/goonode/mogo"
	"pwaa-test.com/utils"
)

var mongoConnection *mogo.Connection = nil

//GetConnection is for get mongo connection
func GetConnection() *mogo.Connection {
	if mongoConnection == nil {
		config := &mogo.Config{
			ConnectionString: utils.DATABASE_URL,
			Database:         utils.DATABASE_NAME,
		}
		mongoConnection, err := mogo.Connect(config)
		if err != nil {
			log.Fatal(err)
		} else {
			return mongoConnection
		}
	}
	return mongoConnection
}