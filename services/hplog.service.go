package service

import (
	"errors"
	"log"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"pwaa-test.com/models/entity"
)

type HpLogService struct {}

func (service *HpLogService) Create(userId, bottleId, text string, worth int) error {
	hpLog := &entity.HpLog{
		Text: text,
		Worth: worth,
		UserId: userId,
		BottleId: bottleId,
	}
	err := mgm.Coll(hpLog).Create(hpLog)
	if err != nil {
		log.Panicf("create log error")
		return err
	}
	
	return nil
}

func (service *HpLogService) GetOne(hplogId string) (*entity.HpLog, error) {
	hpLog := &entity.HpLog{}
	coll := mgm.Coll(hpLog)
	err := coll.FindByID(hplogId, hpLog)
	if err != nil {
		return nil, errors.New("not exist log")
	}
	return hpLog, nil
}

func (service *HpLogService) GetManyByBottle(userId, bottleId string) ([]entity.HpLog, error) {
	result := []entity.HpLog{}
	coll := mgm.Coll(&entity.HpLog{})
	err := coll.SimpleFind(&result, bson.M{"userid": userId, "bottleid": bottleId})

	if err != nil {
		log.Panic(err)
		return nil, errors.New("cannot found bottle")
	}
	return result, nil
}


func (service *HpLogService) GetManyByUser(userId string) ([]entity.HpLog, error) {
	result := []entity.HpLog{}
	coll := mgm.Coll(&entity.HpLog{})
	err := coll.SimpleFind(&result, bson.M{"userid": userId} )

	if err != nil {
		log.Panic(err)
		return nil, errors.New("cannot found bottle")
	}
	return result, nil
}