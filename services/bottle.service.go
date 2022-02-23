package service

import (
	"errors"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"pwaa-test.com/models/entity"
)

type BottleService struct {
	bottle *entity.Bottle
}

func (service BottleService) Create(bottle *entity.Bottle) error {
	// 데이터 들어오면 가공하는 과정
	// userId 찾아서 반환하고 머시기하는거
	err := mgm.Coll(bottle).Create(bottle)
	
	if err != nil {
		return err
	}

	return nil
}

func (service BottleService) FindOne(id string) (*entity.Bottle, error) {
	bottle := &entity.Bottle{}
	coll := mgm.Coll(bottle)
	err := coll.First(bson.M{"id":id}, bottle)

	if err != nil {
		return nil, err
	}

	return bottle, nil
}

func (service BottleService) FindList(userId string) ([]entity.Bottle, error) {
	user := &entity.User{}
	coll := mgm.Coll(user)
	err := coll.First(bson.M{"id":userId}, user)
	
	if err != nil {
		return nil, errors.New("not found USER by Id")
	}
	
	return user.Bottle_list ,nil
}