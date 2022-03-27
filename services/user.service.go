package service

import (
	"errors"
	"log"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"pwaa-test.com/models/entity"
	"pwaa-test.com/module/utils/jwt"
)

type UserService struct {}


func (service *UserService) Create(inputs *entity.User)  error {
	user := &entity.User{}
	coll := mgm.Coll(user)

	if err :=	coll.First(bson.M{"nickName": inputs.NickName}, user); err == nil {
		log.Print("user confilict")
		return errors.New("nickname exist")
	}

	if err := coll.Create(&entity.User{
		Type:inputs.Type,
		NickName: inputs.NickName,
		Key: inputs.Key ,
	}); err != nil {
		return errors.New("create user Error")
	}
	return nil
}

func (service *UserService) FindById(id string) (*entity.User, error) {
	user := &entity.User{}
	err := mgm.Coll(user).FindByID(id, user)
	if err != nil {
		return nil ,err
	}
	return user, nil 
}	

func (service *UserService) FindOauthUser(UserType, userKey string) (*entity.User, error) {
	user := &entity.User{}
	err := mgm.Coll(user).First(bson.M{"type":UserType, "key": userKey}, user)
	if err != nil {
		return nil ,err
	}
	return user, nil 
}	

// 유저 토큰 받기
func (service *UserService) GetToken(id string) (string, error) {
	jwtModule := new(jwt.Module)
	token, err := jwtModule.CreateToken(id)
	return token, err
}

func (service *UserService) FindBykey(userType, key string) (*entity.User, error) {
	user := &entity.User{}
	coll := mgm.Coll(user)
	if err :=	coll.First(bson.M{"type":userType, "key": key}, user); err != nil {
		return nil, errors.New("cannot find user")
	}
	
	return user, nil
}

func (serivce *UserService) FindByNickName( nickName string ) (*entity.User, error) {
	user := &entity.User{}
	coll := mgm.Coll(user)
	if err :=	coll.First(bson.M{"nickName": nickName}, user); err != nil {
		return nil, errors.New("cannot find user")
	}
	
	return user, nil
}