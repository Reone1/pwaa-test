package service

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/dghubble/oauth1"
	"github.com/dghubble/oauth1/twitter"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"pwaa-test.com/models/entity"
	"pwaa-test.com/module/utils/jwt"
	"pwaa-test.com/utils"
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
	if err :=	coll.First(bson.M{"nickName": "testing"}, user); err != nil {
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

	if err :=	coll.First(bson.M{"nickName": "testing"}, user); err == nil {
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

func (service *UserService) FindById(id string) (*entity.User, error) {
	user := &entity.User{}
	err := mgm.Coll(user).FindByID(id, user)
	if err != nil {
		return nil ,err
	}
	return user, nil 
}	
func (service *UserService) FindKakaoUser(kakaoId string) (*entity.User, error) {
	user := &entity.User{}
	err := mgm.Coll(user).First(bson.M{"type":"kakao", "indentity": kakaoId}, user)
	if err != nil {
		return nil ,err
	}
	return user, nil 
}	

// 트위터 토큰 받아오기
func (service *UserService) GetTwitterAccessToken(oauth_token , oauth_token_secret , oauth_verifier ,callbackURL string) (*oauth1.Token, error)  {
	var config = oauth1.Config{
		ConsumerKey:    utils.TWITTER_KEY,
		ConsumerSecret: utils.TWITTER_SECRET,
		CallbackURL:    callbackURL,
		Endpoint:       twitter.AuthorizeEndpoint,
	}

	accessToken, accessSecret, err := config.AccessToken(oauth_token, oauth_token_secret, oauth_verifier)
	if err != nil {
		return nil, err
	}
	
	token := oauth1.NewToken(accessToken, accessSecret)
	return token, nil
}

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

func (service *UserService) GetKakaoOauthToken(grantType, clientId, redirectUri, code string) (string, error) {
	res, err := http.PostForm("https://kauth.kakao.com/oauth/token", url.Values{
		"grant_type": {grantType},
		"client_id": {clientId},
		"redirect_uri": {redirectUri},
		"code": {code},
	})

	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	str := string(respBody)
	return str, nil
}


func (service *UserService) GetKakaoUser(token string) (string, error) {
	req, err := http.NewRequest("get","https://kapi.kakao.com/v2/user/me", nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", "bearer " + token)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	var info struct {
		Id string `json:"id"`
	}
	if err := json.Unmarshal(respBody, &info); err != nil {
		return "", err
	}
	return info.Id, nil
}