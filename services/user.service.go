package service

import (
	"errors"

	"github.com/kamva/mgm/v3"
	"pwaa-test.com/models/entity"
)

type UserSrvice struct {}

func (service UserSrvice) Create(user *entity.User)  error{
	coll := mgm.Coll(user)
	err :=	coll.FindByID(user.ID, user)
	if err != nil {
		return errors.New("already exist")
	}

	err = coll.Create(user)
	if err != nil {
		return errors.New("create user Error")
	}
	return err
}

func (service UserSrvice) Find(user *entity.User) (*entity.User, error) {
	coll := mgm.Coll(user)
	err := coll.FindByID(user.ID, user)
	if err != nil {
		return nil ,err
	}
	return user, nil 
}	