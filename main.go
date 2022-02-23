package main

import (
	"pwaa-test.com/models/db"
	"pwaa-test.com/router"
	"pwaa-test.com/utils"
)
func main (){
	utils.GetENV()
	db.SetDatabase()
	r := router.SetRouter()
	r.Run(":8080")
}