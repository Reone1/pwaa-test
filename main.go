package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"pwaa-test.com/docs"
	"pwaa-test.com/models/db"
	"pwaa-test.com/router"
	"pwaa-test.com/utils"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func main (){
	utils.GetENV()
	db.SetDatabase()
	r := router.SetRouter()
	docs.SwaggerInfo.BasePath = ""

	r.GET("/",func (c *gin.Context) {
		c.Redirect(302, "/swagger/index.html")
	})
	
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	r.Run(":8080")
}