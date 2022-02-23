package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserController struct {}

func (control UserController) GetUser(c *gin.Context) {
	fmt.Print("getUser controller called")
	c.JSON(200, gin.H{
		"message": c.Request.URL,
	})
}

func (control UserController) CreateOne(c *gin.Context) {
	fmt.Print("CreateOne controller called")
}