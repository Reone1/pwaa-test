package entity

import (
	"github.com/kamva/mgm/v3"
)

type Bottle struct {
	mgm.DefaultModel `bson:",inline" swaggerignore:"true"`
	Title string `json:"title" example:"default" binding:"require"`
	Description string `json:"description" example:"" binding:"require"`
	Maturity_date string `json:"maturityDate" example:"" binding:"require"`
	UserId string `json:"userId" example:"Object ID" swaggerignore:"true"`
}