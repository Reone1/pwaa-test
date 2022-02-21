package module

import (
	"github.com/dghubble/oauth1"
	"github.com/dghubble/oauth1/twitter"

	env "pwaa-test.com/utils"
)

var Config oauth1.Config
func SetTwitterConfig(url string) *oauth1.Config{
	var Config = oauth1.Config{
			ConsumerKey:    env.TWITTER_KEY,
			ConsumerSecret: env.TWITTER_SECRET,
			CallbackURL:    url,
			Endpoint:       twitter.AuthorizeEndpoint,
	}
	return &Config
}

func GetToken(accessToken string, accessSecret string) *oauth1.Token{
	return oauth1.NewToken(accessToken, accessSecret)
}
// http.Client will automatically authorize Requests
// var httpClient = config.Client(oauth1.NoContext, token)

// twitter client
// var TwitterClient = twitter.NewClient(httpClient)