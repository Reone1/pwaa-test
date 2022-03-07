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

type  twitterTokenOption struct {
	key string
	secret string
	verifier string
}
// 애플 토큰 발급 과정 => 로그인 타입 및 토큰
// 트위터 토큰 발급 과정 => 로그인 타입 및 토큰
// 애플 계정 정보 불러오기 => call API object Id 
// 트위터 계정 정보 불러오기 => call API object Id 
// jwt 토큰 생성
// jwt 토큰 저장
// auth guard 설정
// 로그아웃
// 테스트 유저용 토큰 생성
//
func (serivce *UserService) TestLogin() (string, error) {
	jwtModule := new(jwt.Module)
	user := &entity.User{}
	coll := mgm.Coll(user)
	if err :=	coll.First(bson.M{"nickname": "testing"}, user); err != nil {
		return "", errors.New("cannot find user")
	}
	
	token, err := jwtModule.CreateToken(user.ID.Hex())
	if err != nil {
		return "", err
	}
	return token, nil
}

func (service *UserService) Create(inputs *entity.User) error {
	user := &entity.User{}
	coll := mgm.Coll(user)

	if err :=	coll.First(bson.M{"nickname": "testing"}, user); err == nil {
		log.Print("user confilict")
		return errors.New("nickname exist")
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

// 트위터 토큰 받아오기
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

// func (service *UserService) GetTwitterAuthToken(callbackURL string) (string, string, error) {
// 	var config = oauth1.Config{
// 		ConsumerKey:    utils.TWITTER_KEY,
// 		ConsumerSecret: utils.TWITTER_SECRET,
// 		CallbackURL:    callbackURL,
// 		Endpoint:       twitter.AuthorizeEndpoint,
// 	}

// 	requestToken, requestSecret, err := config.RequestToken()

// 	if err != nil {
// 		return "", "", err
// 	}
// 	return requestToken, requestSecret, nil
// }