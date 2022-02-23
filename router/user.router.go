package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	controllers "pwaa-test.com/controller"
)

func init() {
	fmt.Print("User Router called")
	controller := new(controllers.UserController)

	logRouter := router.Group("/user")
	{
		logRouter.GET("/", controller.CreateOne)
		logRouter.GET("/list", controller.GetUser)	
	}
	
	answerRoute := logRouter.Group("/answer")
	{
		answerRoute.GET("/", func (c *gin.Context) {
			c.JSON(200, gin.H{
				"message":"hello",
			})
		})
	}
}