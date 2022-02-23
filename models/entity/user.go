package entity

import (
	"github.com/kamva/mgm/v3"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	NickName string `json:"nickName"`
	Mail string `json:"mail"`
	Key string `idx:"{key},unique" json:"key" binding:"required"`
	Type string `json:"type"`
	Bottle_list []Bottle `json:"bottleList"`
}

type Token struct {
	Key string `json:"key"`
	Type string `json:"type"`
}
