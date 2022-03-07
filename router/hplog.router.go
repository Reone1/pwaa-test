package router

import (
	controllers "pwaa-test.com/controller"
	authGaurd "pwaa-test.com/module/utils/auth-gaurd"
)

func init() {
	controller := new(controllers.HplogController)
	hplogRouter := router.Group("/hplog")
	hplogRouter.Use(authGaurd.AuthMiddleware)
	{
		hplogRouter.GET("/", controller.GetHplog)
		
		hplogRouter.GET("/list", controller.GetHplogList)
		hplogRouter.POST("/", controller.CreateHplog)
	}
}