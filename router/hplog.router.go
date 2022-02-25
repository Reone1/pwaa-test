package router

import controllers "pwaa-test.com/controller"

func init() {
	controller := new(controllers.HplogController)
	hplogRouter := router.Group("/hplog")
	{
		hplogRouter.GET("/", controller.GetHplog)
		hplogRouter.GET("/list", controller.GetHplogList)
		hplogRouter.POST("/", controller.CreateHplog)
	}
}