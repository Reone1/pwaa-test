package service

import (
	"errors"
	"log"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"pwaa-test.com/models/entity"
)

type PwaaService struct {}

func (service PwaaService) Create(pwaa *entity.Pwaa) error {
	err := mgm.Coll(pwaa).Create(pwaa)
	if err != nil {
		log.Panicf("create log error")
		return err
	}
	
	return nil
}

func (service PwaaService) GetOne(id string) (*entity.Pwaa, error) {
	pwaa := &entity.Pwaa{}
	coll := mgm.Coll(pwaa)
	err := coll.First(bson.M{"id": id}, pwaa)
	if err != nil {
		return nil, errors.New("not exist log")
	}
	return pwaa, nil
}