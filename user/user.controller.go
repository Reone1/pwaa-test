package User

import (
	"go.mongodb.org/mongo-driver/bson"
	"pwaa-test.com/utils"
)


func UserController() User {
	client, ctx, cancelContext := utils.Mongodb()
	defer cancelContext()
	collection := client.Database("sample-schema").Collection("user")
	result := User{}
	err := collection.FindOne(ctx, bson.D{}).Decode(&result)
	if err != nil{
		panic(err)
	}
	defer func() {
    if err = client.Disconnect(ctx); err != nil {
        panic(err)
    }
	}()
	return result
}