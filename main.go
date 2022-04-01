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
// @description     pwaa 백엔드 REST API swagger.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      ec2-3-34-137-70.ap-northeast-2.compute.amazonaws.com:8080/
// @BasePath  /api/v1

// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
		// c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "http://localhost:8080, http://ec2-3-34-137-70.ap-northeast-2.compute.amazonaws.com:8080, https://ec2-3-34-137-70.ap-northeast-2.compute.amazonaws.com:8080, https://apis.studycfaws.ml")
		c.Header("Access-Control-Allow-Methods", "GET, DELETE, POST")

		if c.Request.Method == "OPTIONS" {
				c.AbortWithStatus(204)
				return
		}

		c.Next()
	}
}
func main (){
	utils.GetENV()
	db.SetDatabase()
	r := router.SetRouter()
	r.Use(CORSMiddleware())

	docs.SwaggerInfo.BasePath = ""

	r.GET("/",func (c *gin.Context) {
		c.Redirect(302, "/swagger/index.html")
	})
	
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	r.Run(":8080")
}