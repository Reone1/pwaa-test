package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"pwaa-test.com/models/entity"
)

type BottleController struct {}
// ShowAccount godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Success      200  {object}  entity.Bottle
// @Router       /bottle [get]
func (b *BottleController) GetOne(c *gin.Context) {
	
	c.JSON(200, entity.Bottle{
		UserId: "hi",
	})
}


func (b *BottleController) Create(c *gin.Context) {
	fmt.Print(c.Request.Body)
	// 파라미터 분리하는 과정 
	c.JSON(200, gin.H{
		"message": "create bottle",
	})
}