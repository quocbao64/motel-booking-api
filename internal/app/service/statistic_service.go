package service

import (
	"awesomeProject/internal/app/constant"
	"awesomeProject/internal/app/pkg"
	"awesomeProject/internal/app/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

type StatisticService interface {
	StatisticUser(c *gin.Context)
	StatisticRoom(c *gin.Context)
}

type StatisticServiceImpl struct {
	statisticRepo repository.StatisticRepository
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

func StatisticServiceInit(repo repository.StatisticRepository) *StatisticServiceImpl {
	return &StatisticServiceImpl{statisticRepo: repo}
}
