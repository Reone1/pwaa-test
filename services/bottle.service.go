package service

import (
	"errors"
	"fmt"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"pwaa-test.com/models/entity"
)

type BottleService struct {}

func (service *BottleService) Create(bottleType, title, userId, date string) (string, error) {
	bottles, err := service.FindList(userId);

	if err != nil {
		return "", err
	}
	length := len(bottles)

	bottle := &entity.Bottle{
		Title: title,
		UserId: userId,
		Type: bottleType,
		Index:fmt.Sprint(length + 1),
		IsOpen: false,
		Description: "default bottle description",
		Maturity_date: date,
	}
	
	if err := mgm.Coll(bottle).Create(bottle); err != nil {
		return "", err
	}

	return bottle.ID.Hex(), nil
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

func (service *BottleService) UpdateIsOpen(userId, bottleId string) error {
	bottle := &entity.Bottle{}
	if err := mgm.Coll(&entity.Bottle{}).FindByID(bottleId, bottle); err != nil {
		return err
	}
	bottle.IsOpen = !bottle.IsOpen
	if err := mgm.Coll(bottle).Update(bottle); err != nil {
		return err
	}
	return nil
}

func (service *BottleService) UpdateImgUri(userId, bottleId, UpdateImgUri string) error {
	bottle := &entity.Bottle{}
	if err := mgm.Coll(&entity.Bottle{}).FindByID(bottleId, bottle); err != nil {
		return err
	}
	bottle.ImgUri = UpdateImgUri
	if err := mgm.Coll(bottle).Update(bottle); err != nil {
		return err
	}
	return nil
}