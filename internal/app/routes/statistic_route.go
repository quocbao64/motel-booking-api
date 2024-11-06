package routes

import (
	"awesomeProject/config"
	"github.com/gin-gonic/gin"
)

func StatisticRoute(init *config.Initialize, g *gin.RouterGroup) *gin.RouterGroup {
	g.GET("/user", init.StatisticCtrl.StatisticUser)
	g.GET("/room", init.StatisticCtrl.StatisticRoom)
	return g
}
