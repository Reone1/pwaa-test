package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	httputil "pwaa-test.com/module/utils"
	service "pwaa-test.com/services"
)

var hplogService = new(service.HpLogService)
type HplogController struct{}

type GetHplogRequestBody struct {
	ID string `json:"id" example:"hplog ID"`
}
// ShowAccount godoc
// @Summary      기록 세부사항 조회
// @Description  단일 기록 세부사항 조회
// @Tags         hplog
// @Accept       json 
// @Param				 id   query GetHplogRequestBody false "hplog Request Body Data"
// @Success      200  {object}  entity.HpLog
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /hplog [get]
// @Security ApiKeyAuth
func (control *HplogController) GetHplog(c *gin.Context) {
	if _, ok := c.Get("userId"); !ok {
		httputil.NewError(c, http.StatusUnauthorized, errors.New("not Found UserId"))
	}

	ID := c.Request.URL.Query().Get("id")
	hplog, err :=hplogService.GetOne(ID)
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
	}
	c.JSON(200, hplog)
}

type CreateHplogResponse struct {
	Message string `json:"message" example:"ok"`
}

type CreateHplogRequestBody struct {
	BottleId string `json:"bottleId" example:"bottle ID"`
	Text string `json:"text" example:"each log text"`
	Worth int `json:"worth" example:"0"`
}
// ShowAccount godoc
// @Summary      로그 생성
// @Description  유리병에 단일 로그 생성 
// @Tags         hplog
// @Accept       json
// @Param        body  body  CreateHplogRequestBody false   "create hplog"
// @Success      200  {object}  CreateHplogResponse
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /hplog [post]
// @Security ApiKeyAuth
func (control *HplogController) CreateHplog(c *gin.Context) {
	getId, ok := c.Get("userId")
	if !ok {
		httputil.NewError(c, http.StatusBadRequest, errors.New("not found UserId"))
		return	
	}
	userId := fmt.Sprintf("%v", getId)
	var body CreateHplogRequestBody

	if err := c.ShouldBindJSON(&body); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	if err:= hplogService.Create(userId, body.BottleId, body.Text, body.Worth); err != nil{
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	
 	c.JSON(200, gin.H{
		"message": "create Log",
	})
}


type GetHplistRequestBody struct { 
	BottleId string `json:"bottleId" example:"bottle ID"`
}
// ShowAccount godoc
// @Summary      기록 목록 조회
// @Description  다중기록 조회 (유리병 단위)
// @Tags         hplog
// @Accept       json
// @Param        id   query GetHplistRequestBody  false   "Bottle Id"
// @Success      200  {array}   entity.HpLog
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /hplog/list [get]
// @Security ApiKeyAuth
func (control HplogController) GetHplogList(c *gin.Context) {
	getId, ok := c.Get("userId")
	if !ok {
		httputil.NewError(c, http.StatusBadRequest, errors.New("not found UserId"))
		return	
	}
	userId := fmt.Sprintf("%v", getId)
	BottleId := c.Request.URL.Query().Get("bottleId")
	hplogs, err := hplogService.GetManyByBottle(userId, BottleId)
	if err != nil{
		httputil.NewError(c, http.StatusNotExtended, err)
		return
	}
	c.JSON(200, hplogs)
}
