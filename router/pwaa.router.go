package router

import (
	controllers "pwaa-test.com/controller"
	authGaurd "pwaa-test.com/module/utils/auth-gaurd"
)

func init() {
	controller := new(controllers.PwaaController)
	hplogRouter := router.Group("/pwaa")
	hplogRouter.Use(authGaurd.AuthMiddleware)
	{
		hplogRouter.GET("/", controller.GetPwaa)
		
		hplogRouter.GET("/list", controller.GetPwaaList)
		hplogRouter.POST("/", controller.CreatePwaa)
	}
}