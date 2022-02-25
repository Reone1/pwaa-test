package entity

import (
	"github.com/kamva/mgm/v3"
)

type Bottle struct {
	mgm.DefaultModel `bson:",inline" swaggerignore:"true"`
	Title string `json:"title"`
	Description string `json:"description"`
	Maturity_date string `json:"maturityDate"`
	UserId string `json:"userId"`
}