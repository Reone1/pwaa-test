package main

import (
	"github.com/gin-gonic/gin"
	User "pwaa-test.com/user"
	"pwaa-test.com/utils"
)
func main (){
	utils.GetENV()
	r := gin.Default()
	r.GET("/user", func (c *gin.Context) {
		user := User.UserController()
		c.JSON(200, user)
	})
	twitterRoute := r.Group("/twitter")
	{
		twitterRoute.GET("/request-token", func (c *gin.Context) {
			token, secret := User.GetTwitterToken(c.Query("callback_url"))
			c.JSON(200, gin.H{
				"oauth_token": token,
				"oauth_token_secret": secret,
			})
		})
		twitterRoute.GET("/access-token", func (c *gin.Context) {
			oauth_token:=c.Query("oauth_token")
			oauth_token_secret := c.Query("oauth_token_secret")
			oauth_verifier:= c.Query("oauth_verifier")
			token := User.TwitterAccessToken(oauth_token,oauth_token_secret,oauth_verifier)
			c.JSON(200, gin.H{ "token": token })
		})
	}
	r.Run(":8080")
}