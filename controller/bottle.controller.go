package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type BottleController struct {}
func (b *BottleController) GetOne(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": c.Request.URL,
	})
}

func (b *BottleController) Create(c *gin.Context) {
	fmt.Print(c.Request.Body)
	// 파라미터 분리하는 과정 
	c.JSON(200, gin.H{
		"message": "create bottle",
	})
}