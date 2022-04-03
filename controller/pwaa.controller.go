package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	httputil "pwaa-test.com/module/utils"
	service "pwaa-test.com/services"
)

var PwaaService = new(service.PwaaService)
type PwaaController struct{}

type GetPwaaRequestBody struct {
	ID string `json:"id" example:"pwaa ID"`
}
// ShowAccount godoc
// @Summary      기록 세부사항 조회
// @Description  단일 기록 세부사항 조회
// @Tags         pwaa
// @Accept       json 
// @Param				 id   query GetPwaaRequestBody false "pwaa Request Body Data"
// @Success      200  {object}  entity.Pwaa
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /pwaa [get]
// @Security ApiKeyAuth
func (control *PwaaController) GetPwaa(c *gin.Context) {
	if _, ok := c.Get("userId"); !ok {
		httputil.NewError(c, http.StatusUnauthorized, errors.New("not Found UserId"))
	}

	ID := c.Request.URL.Query().Get("id")
	pwaa, err := PwaaService.GetOne(ID)
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
	}
	c.JSON(200, pwaa)
}

type CreatePwaaResponse struct {
	Message string `json:"message" example:"ok"`
}

type CreatePwaaRequestBody struct {
	BottleId string `json:"bottleId" example:"bottle ID"`
	Content string `json:"content" example:"each log text"`
	Worth int `json:"worth" example:"0"`
}
// ShowAccount godoc
// @Summary      로그 생성
// @Description  유리병에 단일 로그 생성 
// @Tags         pwaa
// @Accept       json
// @Param        body  body  CreatePwaaRequestBody false   "create pwaa"
// @Success      200  {object}  CreatePwaaResponse
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /pwaa [post]
// @Security ApiKeyAuth
func (control *PwaaController) CreatePwaa(c *gin.Context) {
	getId, ok := c.Get("userId")
	if !ok {
		httputil.NewError(c, http.StatusBadRequest, errors.New("not found UserId"))
		return	
	}
	userId := fmt.Sprintf("%v", getId)
	var body CreatePwaaRequestBody

	if err := c.ShouldBindJSON(&body); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	if err:= PwaaService.Create(userId, body.BottleId, body.Content, body.Worth); err != nil{
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	
 	c.JSON(200, gin.H{
		"message": "ok",
	})
}


type GetPwaasRequestBody struct { 
	BottleId string `json:"bottleId" example:"bottle ID"`
}
// ShowAccount godoc
// @Summary      기록 목록 조회
// @Description  다중기록 조회 (유리병 단위)
// @Tags         pwaa
// @Accept       json
// @Param        id   query GetPwaasRequestBody  false   "Bottle Id"
// @Success      200  {array}   entity.Pwaa
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /pwaa/list [get]
// @Security ApiKeyAuth
func (control *PwaaController) GetPwaaList(c *gin.Context) {
	if _, ok := c.Get("userId"); !ok {
		httputil.NewError(c, http.StatusBadRequest, errors.New("not found UserId"))
		return	
	}
	BottleId := c.Request.URL.Query().Get("bottleId")
	pwaas, err := PwaaService.GetManyByBottle( BottleId)
	if err != nil{
		httputil.NewError(c, http.StatusNotExtended, err)
		return
	}
	c.JSON(200, pwaas)
}
