package routes

import (
	"awesomeProject/config"
	"github.com/gin-gonic/gin"
)

func UserRoute(init *config.Initialize, g *gin.RouterGroup) *gin.RouterGroup {
	//g.GET("", middleware.AuthMiddleware(), init.CustomerCtrl.GetAll)
	//g.POST("", init.CustomerCtrl.Create)
	return g
}
