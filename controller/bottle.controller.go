package controllers

import (
	"errors"
	"fmt"
	"net/http"

	httputil "pwaa-test.com/module/utils"

	"github.com/gin-gonic/gin"
	"pwaa-test.com/models/entity"
	service "pwaa-test.com/services"
)

var bottleService = new(service.BottleService)

type BottleController struct {}

type GetBottleRequestQuery struct {
	BottleId string `json:"bottleId" example:"bottle ID"`
}
type GetBottleResponse struct {
	Title string `json:"title" example:"Bottle title"`
	TotalWorth int `json:"totalWorth" example:"2300000"`
	Maturity_date string `json:"maturityDate" example:"2022-03-04T03:16:49.767Z" binding:"require"`
	HplogList []entity.HpLog `json:"hplogList" binding:"require"`
}
// ShowAccount godoc
// @Summary      Show an Bottle
// @Description  get string by Bottle ID
// @Tags         bottle
// @Accept       json
// @Param				 query query GetBottleRequestQuery false "bottle's hplog list"
// @Success      200  {object}  GetBottleResponse
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /bottle [get]
// @Security ApiKeyAuth
func (b *BottleController) GetOne(c *gin.Context) {
	data, ok := c.Get("userId")
	if !ok {
		httputil.NewError(c, http.StatusBadRequest, errors.New("not authorized"))
		return	
	}
	userId := fmt.Sprintf("%v", data)
	query := c.Request.URL.Query()
	var totalWorth int = 0
	hplogs, err := hplogService.GetManyByBottle(userId, query.Get("bottleId"));
	if  err != nil {
		httputil.NewError(c, http.StatusNotFound, err)
		return
	}
	for _, log := range hplogs {
		totalWorth += log.Worth
	}
	bottle, err := bottleService.FindOne( query.Get("bottleId"))
	
	if err != nil {
		httputil.NewError(c, http.StatusNotFound, err)
		return
	}
	c.JSON(200, GetBottleResponse{
		Title: bottle.Title,
		TotalWorth: totalWorth,
		Maturity_date: bottle.Maturity_date,
		HplogList: hplogs,
	})
}

// ShowAccount godoc
// @Summary      유리병 목록 조회
// @Description  GET bottle list
// @Tags         bottle
// @Accept       json
// @Success      200  {array}  entity.Bottle
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /bottle/list [get]
// @Security ApiKeyAuth
func (b *BottleController) GetMany(c *gin.Context) {
	data, ok := c.Get("userId")
	if !ok {
		httputil.NewError(c, http.StatusBadRequest, errors.New("not authorized"))
		return	
	}
	userId := fmt.Sprintf("%v", data)
	bottles, err := bottleService.FindList(userId)
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return	
	}
	c.JSON(200, bottles)
}

type CreateBottleRequestBody struct {
	Title string `json:"title" example:"bottle title (optional)"`
	MaturityDate string `json:"maturityDate" example:"date JSON string"`
}
// ShowAccount godoc
// @Summary      Create BOTTLE by userID
// @Description  새로운 유리병을 생성합니다. <br /> 유리병의 이름을 입력할 수 있습니다. 아무값 없이 요청하면 "default" 이름을 갖게 됩니다.
// @Tags         bottle
// @Param        title body CreateBottleRequestBody false "Create Bottle request body"
// @Accept       json
// @Success      200  {object}  entity.Bottle
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /bottle [post]
// @Security ApiKeyAuth
func (b *BottleController) Create(c *gin.Context) {
	getData, ok := c.Get("userId")
	if !ok {
		httputil.NewError(c, http.StatusBadRequest, errors.New("not found UserId"))
		return	
	}
	userId := fmt.Sprintf("%v", getData)
	bottleService := new(service.BottleService)
	
	var body CreateBottleRequestBody
	
	if err := c.ShouldBindJSON(&body); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	bottleId, err := bottleService.Create(body.Title, userId, body.MaturityDate)
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(200, gin.H{
		"bottleId": bottleId,
	})
}