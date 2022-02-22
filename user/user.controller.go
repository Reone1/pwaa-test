package user

import (
	"go.mongodb.org/mongo-driver/bson"
	"pwaa-test.com/utils"
)


func UserController() User {
	client, cancelContext := utils.Mongodb()
	defer cancelContext()
	collection := client.Database("sample-schema").Collection("user")
	result := User{}
	err := collection.FindOne(*utils.CTX, bson.D{}).Decode(&result)
	if err != nil{
		panic(err)
	}
	defer func() {
    if err = client.Disconnect(*utils.CTX); err != nil {
        panic(err)
    }
	}()
	return result
}