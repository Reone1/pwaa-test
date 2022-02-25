package controllers

import (
	"errors"
	"net/http"

	httputil "pwaa-test.com/module/utils"

	"github.com/gin-gonic/gin"
	"pwaa-test.com/models/entity"
	service "pwaa-test.com/services"
)

type BottleController struct {}

// ShowAccount godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         bottles
// @Accept       json
// @Success      200  {object}  entity.Bottle
// @Router       /bottle [get]
func (b *BottleController) GetOne(c *gin.Context) {
	c.JSON(200, entity.Bottle{
		UserId: "hi",
	})
}


// ShowAccount godoc
// @Summary      유리병 목록 조회
// @Description  GET bottle list
// @Tags         bottles
// @Accept       json
// @Success      200  {object}  []entity.Bottle
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /bottles [get]
func (b *BottleController) GetMany(c *gin.Context) {
	var err error = errors.New("string")
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(200, gin.H{
		"result": [...]entity.Bottle{{
			UserId: "hi",
		}},
	})
}

// ShowAccount godoc
// @Summary      Create BOTTLE by userID
// @Description  새로운 유리병을 생성합니다.\n 유리병의 이름을 입력할 수 있습니다. 아무값 없이 요청하면 "default" 이름을 갖게 됩니다.
// @Tags         bottles
// @Param        title body string false "Bottle name"
// @Accept       json
// @Success      200  {object}  entity.Bottle
// @Router       /bottle [post]
func (b *BottleController) Create(c *gin.Context) {

	bottleService := new(service.BottleService)

	var body struct {
		title string
		userId string
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	bottleId, err := bottleService.Create("test", "62184020b83b2ef729a5a5d0")
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	// 파라미터 분리하는 과정 
	c.JSON(200, gin.H{
		"bottleId": bottleId,
	})
}