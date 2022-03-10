package entity

import (
	"github.com/kamva/mgm/v3"
)

type Bottle struct {
	mgm.DefaultModel `bson:",inline" swaggerignore:"true"`
	Title string `json:"title" bson:"title" example:"default" binding:"require"`
	Index string `json:"index" example:"1" binding:"require"`
	Type string `json:"type" example:"1" binding:"require"`
	IsOpen bool `json:"isOpen" bson:"isOpen" example:"false" binding:"require"`
	ImgUri string `json:"imgUri" bson:"imgUri" example:"any.img.uri" binding:"require"`
	Description string `json:"description" example:"" binding:"require"`
	Maturity_date string `json:"maturityDate" bson:"maturityDate" example:"" binding:"require"`
	UserId string `json:"userId" bson:"userId" example:"Object ID" swaggerignore:"true"`
}