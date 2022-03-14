package entity

import (
	"github.com/kamva/mgm/v3"
)

type Pwaa struct {
	mgm.DefaultModel `bson:",inline"`
	Content string `json:"content" binding:"required"`
	Worth int `json:"worth" binding:"required"`
	UserId string `idx:"{userId},unique" json:"userId" bson:"userId"`
	BottleId string `idx:"{bottleId},unique" json:"bottleId" bson:"bottleId"`
}
