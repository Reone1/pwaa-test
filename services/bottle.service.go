package service

import (
	"errors"
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"pwaa-test.com/models/entity"
)

type BottleService struct {}

func (service *BottleService) Create(title string, userId string) (string, error) {
	bottle := &entity.Bottle{
		Title: title,
		UserId: userId,
		Description: "default bottle description",
		Maturity_date: time.Now().AddDate(0,0,3).UTC().String(),
	}
	err := mgm.Coll(bottle).Create(bottle)
	
	if err != nil {
		return "", err
	}

	return bottle.ID.String(), nil
}

func (service *BottleService) FindOne(bottleId string) (*entity.Bottle, error) {
	bottle := &entity.Bottle{}
	coll := mgm.Coll(bottle)
	err := coll.FindByID(bottleId, bottle)

	if err != nil {
		return nil, err
	}

	return bottle, nil
}

func (service *BottleService) FindList(userId string) ([]entity.Bottle, error) {
	bottles := []entity.Bottle{}
	err := mgm.Coll(&entity.Bottle{}).SimpleFind(&bottles, bson.M{"userId": userId})
	
	if err != nil {
		return nil, errors.New("not found USER by Id")
	}
	
	return bottles ,nil
}