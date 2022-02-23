package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type PwaaController struct{}

func (control PwaaController) GetPwaa(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": c.Request.URL.Path,
	})
}

func (control PwaaController) CreatePwaa(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "create Log",
	})
}

func (control PwaaController) GetPwaaList(c *gin.Context) {
	fmt.Print("GetPwaaList called")
	c.JSON(200, gin.H{
		"message": "pwaa list",
	})
}