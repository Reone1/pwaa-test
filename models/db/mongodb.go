package db

import (
	"log"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"pwaa-test.com/utils"
)

//GetConnection is for get mongo connection
func SetDatabase() {
	err := mgm.SetDefaultConfig(nil, utils.DATABASE_NAME, options.Client().ApplyURI(utils.DATABASE_URL))
	if err != nil {
		log.Fatal(err)
	}
}