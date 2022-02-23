package entity

import (
	"github.com/kamva/mgm/v3"
)

type Bottle struct {
	mgm.DefaultModel `bson:",inline"`
	Title string `json:"title"`
	Description string `json:"description"`
	Maturity_date string `json:"maturityDate"`
	Log_list []Pwaa `json:"logList"`
	UserId string `json:"userId"`
	Total_worth int `json:"totalWorth"`
}