package service

import (
	"awesomeProject/internal/app/constant"
	"awesomeProject/internal/app/pkg"
	"awesomeProject/internal/app/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type StatisticService interface {
	StatisticUser(c *gin.Context)
	StatisticRoom(c *gin.Context)
	Statistic(c *gin.Context)
}

type StatisticServiceImpl struct {
	statisticRepo repository.StatisticRepository
}

type StatisticParams struct {
	Year  int `json:"year" form:"year"`
	Month int `json:"month" form:"month"`
}

func (repo StatisticServiceImpl) StatisticUser(c *gin.Context) {
	data, err := repo.statisticRepo.StatisticUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, err, pkg.Null()))
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (repo StatisticServiceImpl) StatisticRoom(c *gin.Context) {
	data, err := repo.statisticRepo.StatisticRoom()

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, err, pkg.Null()))
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (repo StatisticServiceImpl) Statistic(c *gin.Context) {
	params := StatisticParams{}
	err := c.BindQuery(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, err, pkg.Null()))
		return
	}

	filter := &repository.StatisticFilter{
		Year:  params.Year,
		Month: params.Month,
	}
	fmt.Println(filter)
	data, err := repo.statisticRepo.Statistic(filter)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, err, pkg.Null()))
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func StatisticServiceInit(repo repository.StatisticRepository) *StatisticServiceImpl {
	return &StatisticServiceImpl{statisticRepo: repo}
}
