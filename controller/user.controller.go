package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"pwaa-test.com/models/entity"
	httputil "pwaa-test.com/module/utils"
	service "pwaa-test.com/services"
)

var userService = new(service.UserService)
type UserController struct {}

func (control UserController) GetUser(c *gin.Context) {
	fmt.Print("getUser controller called")
	c.JSON(200, gin.H{
		"message": c.Request.URL,
	})
}

func (control UserController) CreateOne(c *gin.Context) {
	if err := userService.Create(&entity.User{}); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(200,gin.H{
		"message": "success Create User",
	})
}