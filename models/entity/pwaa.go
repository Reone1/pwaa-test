package entity

import (
	"github.com/kamva/mgm/v3"
)

type Pwaa struct {
	mgm.DefaultModel `bson:",inline" swaggerignore:"true"`
	Content string `json:"content" binding:"required"`
	Worth int `json:"worth" binding:"required"`
	UserId string `idx:"{userId},unique" json:"userId" bson:"userId" swaggerignore:"true"`
	BottleId string `idx:"{bottleId},unique" json:"bottleId" bson:"bottleId" swaggerignore:"true"`
}
