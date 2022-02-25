package service

import (
	"errors"
	"log"

	"github.com/dghubble/oauth1"
	"github.com/dghubble/oauth1/twitter"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"pwaa-test.com/models/entity"
	"pwaa-test.com/utils"
)

type UserService struct {}

type  twitterTokenOption struct {
	key string
	secret string
	verifier string
}

func (service *UserService) Create(inputs *entity.User)  error{
	user := &entity.User{}
	coll := mgm.Coll(user)

	if err :=	coll.First(bson.M{"nickname": "testing"}, user); err == nil {
		log.Print("user confilict")
		return errors.New("already exist")
	}

	err := coll.Create(&entity.User{
		Type:"kakao",
		Key: "string",
		NickName: "testing",
		Mail: "testing@mail",
	})
	if err != nil {
		return errors.New("create user Error")
	}
	return err
}

func (service *UserService) Find(id string) (*entity.User, error) {
	user := &entity.User{}
	err := mgm.Coll(user).FindByID(id, user)
	if err != nil {
		return nil ,err
	}
	return user, nil 
}	
// func (service *UserService) getTwitterAccessToken(oauth_token string, oauth_token_secret string, oauth_verifier string) (*oauth1.Token, error)  {
// 	var option = twitterTokenOption{
// 		key: oauth_token,
// 		secret: oauth_token_secret,
// 		verifier: oauth_verifier,
// 	}
// 	accessToken, accessSecret, err := config.AccessToken(option.key, option.secret, option.verifier)
// 	if err != nil {
// 		panic("Twitter Access Token Error!")
// 		return nil, err
// 	}
// 	token := oauth1.NewToken(accessToken, accessSecret)
// 	return token, nil
// }

func (service *UserService) GetTwitterAuthToken(callbackURL string) (string, string, error) {
	var config = oauth1.Config{
		ConsumerKey:    utils.TWITTER_KEY,
		ConsumerSecret: utils.TWITTER_SECRET,
		CallbackURL:    callbackURL,
		Endpoint:       twitter.AuthorizeEndpoint,
	}

	requestToken, requestSecret, err := config.RequestToken()

	if err != nil {
		return "", "", err
	}
	return requestToken, requestSecret, nil
}