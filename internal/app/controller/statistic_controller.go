package controller

import (
	"awesomeProject/internal/app/service"
	"github.com/gin-gonic/gin"
)

type StatisticController interface {
	StatisticUser(c *gin.Context)
}

type StatisticControllerImpl struct {
	statisticService service.StatisticService
}

func (controller StatisticControllerImpl) StatisticUser(c *gin.Context) {
	controller.statisticService.StatisticUser(c)
}

func StatisticControllerInit(statisticService service.StatisticService) *StatisticControllerImpl {
	return &StatisticControllerImpl{statisticService: statisticService}
}
