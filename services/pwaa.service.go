package service

import (
	"errors"
	"fmt"
	"log"

	"github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/builder"
	"go.mongodb.org/mongo-driver/bson"
	"pwaa-test.com/models/entity"
)

type PwaaService struct {}

func (service *PwaaService) Create(userId, bottleId, Content string, worth int) error {
	pwaa := &entity.Pwaa{
		Content: Content,
		Worth: worth,
		UserId: userId,
		BottleId: bottleId,
	}
	err := mgm.Coll(pwaa).Create(pwaa)
	if err != nil {
		log.Panicf("create log error")
		return err
	}
	
	return nil
}

func (service *PwaaService) GetOne(pwaaId string) (*entity.Pwaa, error) {
	pwaa := &entity.Pwaa{}
	coll := mgm.Coll(pwaa)
	err := coll.FindByID(pwaaId, pwaa)
	if err != nil {
		return nil, errors.New("not exist log")
	}
	return pwaa, nil
}

func (service *PwaaService) GetManyByBottle(bottleId string) ([]entity.Pwaa, error) {
	result := []entity.Pwaa{}
	coll := mgm.Coll(&entity.Pwaa{})
	err := coll.SimpleFind(&result, bson.M{ "bottleId": bottleId})

	if err != nil {
		log.Panic(err)
		return nil, errors.New("cannot found bottle")
	}
	return result, nil
}


func (service *PwaaService) GetManyByUser(userId string) ([]entity.Pwaa, error) {
	result := []entity.Pwaa{}
	coll := mgm.Coll(&entity.Pwaa{})
	err := coll.SimpleFind(&result, bson.M{"userId": userId} )

	if err != nil {
		log.Panic(err)
		return nil, errors.New("cannot found bottle")
	}
	return result, nil
}
func (service *PwaaService) GetPriceOfPwaas(bottleId string) (int, error) {
	coll := mgm.Coll(&entity.Pwaa{})
	result := struct{
		Price int `json:"count"`
	}{}

	err := coll.SimpleAggregate(&result,builder.Group("count", bson.M{"$sum": "$worth"}))
	fmt.Print(result)
	if err != nil {
		return 0, err
	}
	return result.Price, nil
}