package User

import (
	"fmt"
	"log"

	"github.com/dghubble/oauth1"
	"go.mongodb.org/mongo-driver/mongo"
	TwitterModule "pwaa-test.com/user/module"
)
type  twitterTokenOption struct {
	key string
	secret string
	verifier string
}

var userCollection *mongo.Collection

func init() {
	if userCollection == nil {
		fmt.Print("go")
		// userCollection = utils.GetCollection("user")
	}
}

func GetUsers(ids ...string) []User {
	return []User{{
		ID: "user id string1",
		NickName: "user NickName string1",
		Type: "user type string1",
	},
	{
		ID: "user id string2",
		NickName: "user NickName string2",
		Type: "user type string2",
	}}
}

func createToken() Token {
	return Token{
		Key: "string",
		Type: "string",
	}
}

func kakaoAuth() Token {
	// kakao login
	return createToken()
}

func GetTwitterToken(callbackURL string) (string, string) {
	config := TwitterModule.SetTwitterConfig(callbackURL)

	requestToken, requestSecret, err := config.RequestToken()
	if err != nil {
		log.Panic("Twitter requestToken Error: ", err)
	}

	return requestToken, requestSecret
}

func TwitterAccessToken(oauth_token string, oauth_token_secret string, oauth_verifier string) *oauth1.Token {
	var option = twitterTokenOption{
		key: oauth_token,
		secret: oauth_token_secret,
		verifier: oauth_verifier,
	}

	accessToken, accessSecret, err := TwitterModule.Config.AccessToken(option.key, option.secret, option.verifier)

	if err != nil {
		panic("Twitter Access Token Error!")
	}

	token := TwitterModule.GetToken(accessToken, accessSecret)
	return token
}