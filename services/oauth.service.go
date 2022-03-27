package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/dghubble/oauth1"
	"github.com/dghubble/oauth1/twitter"
	"pwaa-test.com/utils"
)


type  TwitterTokenOption struct {
	Key string
	Secret string
	Verifier string
}

type OAuthService struct {}

func (service *OAuthService) GetTwitterAccessToken(oauth_token, oauth_token_secret, oauth_verifier, callbackURL string) (*oauth1.Token, error) {
// 트위터 토큰 받아오기
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
func (service *OAuthService) GetTwitterCode(code string) {}


// 트위터 accessToken 받기
func (service *OAuthService) GetTwitterAuthToken(callbackURL string) (string, string, error) {
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

// 카카오 토큰 받기
func (service *OAuthService) GetKakaoOauthToken(code string) (string, error) {
	res, err := http.PostForm("https://kauth.kakao.com/oauth/token", url.Values{
		"grant_type": {"authorization_code"},
		"client_id": {utils.KAKAO_CLIENT_ID},
		"redirect_uri": {"http://ec2-3-34-137-70.ap-northeast-2.compute.amazonaws.com:8080/oauth/kakao"},
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
	type KakaoRes struct {
		Token string `json:"access_token"`
	}
	kakaoResp := KakaoRes{}
	if err := json.Unmarshal([]byte(str), &kakaoResp);err != nil{
		return "",err
	}
	return kakaoResp.Token, nil
}

// 카카오 유저 찾기
func (service *OAuthService) GetKakaoUser(token string) (string, error) {
	req, err := http.NewRequest("GET","https://kapi.kakao.com/v2/user/me", nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", "Bearer " + token)

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
		Id int `json:"id"`
	}
	if err := json.Unmarshal(respBody, &info); err != nil {
		return "", err
	}
	return fmt.Sprint(info.Id), nil
}