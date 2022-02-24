package entity

import (
	"github.com/kamva/mgm/v3"
)

type Pwaa struct {
	mgm.DefaultModel `bson:",inline" swaggerignore:"true"`
	Text string `json:"text" binding:"required"`
	Worth int `json:"worth" binding:"required"`
	UserId string `idx:"{userId},unique" json:"userId" binding:"required"`
	Bottle_Id string `idx:"{bottleId},unique" json:"bottleId" binding:"required"`
}
