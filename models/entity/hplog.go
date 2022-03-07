package entity

import (
	"github.com/kamva/mgm/v3"
)

type HpLog struct {
	mgm.DefaultModel `bson:",inline" swaggerignore:"true"`
	Text string `json:"text" binding:"required"`
	Worth int `json:"worth" binding:"required"`
	UserId string `idx:"{userId},unique" json:"userId" swaggerignore:"true"`
	BottleId string `idx:"{bottleId},unique" json:"bottleId" swaggerignore:"true"`
}
