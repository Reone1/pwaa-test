package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	httputil "pwaa-test.com/module/utils"
	service "pwaa-test.com/services"
)

var hplogService = new(service.HpLogService)
type HplogController struct{}

func (control *HplogController) GetHplog(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": c.Request.URL.Path,
	})
}


func (control *HplogController) CreateHplog(c *gin.Context) {
	var body struct {
		UserId string `json:"userId"`
		BottleId string `json:"bottleId"`
		Text string `json:"text"`
		Worth int `json:"worth"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	if err:= hplogService.Create("", "", body.Text, body.Worth); err != nil{
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	
 	c.JSON(200, gin.H{
		"message": "create Log",
	})
}

func (control HplogController) GetHplogList(c *gin.Context) {
	var body struct {
		BottleId string `json:"bottleId"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	hplogs, err := hplogService.GetManyByBottle("621893bf296ce382ff06e70a")
	if err != nil{
		httputil.NewError(c, http.StatusNotExtended, err)
		return
	}
	c.JSON(200, hplogs)
}
func (control *HplogController) GetOne(c *gin.Context) {


	c.JSON(200, gin.H{
		"message": "log detail",
	})
}