package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"pwaa-test.com/models/entity"
	utils "pwaa-test.com/utils"
)

type BottleService struct {}

func (service *BottleService) Create(bottleType, title, userId string, date time.Time) (*entity.Bottle, error) {
	bottles, err := service.FindList(userId, "");

	if err != nil {
		return  nil, err
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
		return nil, err
	}

	return bottle, nil
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

func (service *BottleService) FindList(userId , isOpen string) ([]entity.Bottle, error) {
	bottles := []entity.Bottle{}
	var err error
	if isOpen == "" {
		err = mgm.Coll(&entity.Bottle{}).SimpleFind(&bottles, bson.M{"userId": userId})
	} else {
		openStatus := false
		if isOpen == "true" {
			openStatus = true;
		}
		err = mgm.Coll(&entity.Bottle{}).SimpleFind(&bottles, bson.M{"userId": userId, "isOpen": openStatus})
	}
	
	if err != nil {
		return nil, errors.New("not found USER by Id")
	}
	
	return bottles ,nil
}

func (service *BottleService) UpdateIsOpenStatus(userId, bottleId string) error {
	bottle := &entity.Bottle{}
	if err := mgm.Coll(&entity.Bottle{}).FindByID(bottleId, bottle); err != nil {
		return err
	}
	bottle.IsOpen = !bottle.IsOpen
	var totalWorth int = 0
	pwaaService := new(PwaaService)
	pwaas, err := pwaaService.GetManyByBottle(userId)
	if err != nil {
		return err
	}
	for _, pwaa := range pwaas {
		totalWorth += pwaa.Worth
	}
	utils.ImgPathStr(totalWorth)
	if err := mgm.Coll(bottle).Update(bottle); err != nil {
		return err
	}
	return nil
}

func (service *BottleService) UpdateImgUri(userId, bottleId, imgUri string) error {
	bottle := &entity.Bottle{}
	if err := mgm.Coll(&entity.Bottle{}).FindByID(bottleId, bottle); err != nil {
		return err
	}
	bottle.ImgUri = imgUri
	if err := mgm.Coll(bottle).Update(bottle); err != nil {
		return err
	}
	return nil
}