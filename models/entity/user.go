package entity

import (
	"github.com/kamva/mgm/v3"
)

type User struct {
	mgm.DefaultModel `bson:",inline" swaggerignore:"true"`
	NickName string `json:"nickName" bson:"nickName" binding:"require"`
	Mail string `json:"mail"`
	Key string `idx:"{key},unique" json:"key" binding:"required"`
	Type string `json:"type" binding:"require"`
}
