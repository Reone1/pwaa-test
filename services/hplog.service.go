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
	if userId == "" {
		userId = "62184020b83b2ef729a5a5d0"
	} 
	if bottleId == "" {
		bottleId = "621893bf296ce382ff06e70a"
	}
	
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
	err := coll.First(bson.M{"id": hplogId}, hpLog)
	if err != nil {
		return nil, errors.New("not exist log")
	}
	return hpLog, nil
}

func (service *HpLogService) GetManyByBottle(bottleId string) ([]entity.HpLog, error) {
	result := []entity.HpLog{}
	coll := mgm.Coll(&entity.HpLog{})
	err := coll.SimpleFind(&result, bson.M{"bottleid": bottleId})

	if err != nil {
		log.Panic(err)
		return nil, errors.New("cannot found bottle")
	}
	return result, nil
}


func (service *HpLogService) GetManyByUser(userId string) ([]entity.HpLog, error) {
	result := []entity.HpLog{}
	coll := mgm.Coll(&entity.HpLog{})
	err := coll.SimpleFind(&result, bson.M{"userId": userId} )

	if err != nil {
		log.Panic(err)
		return nil, errors.New("cannot found bottle")
	}
	return result, nil
}