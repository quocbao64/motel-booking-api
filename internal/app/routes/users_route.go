package routes

import (
	"awesomeProject/config"
	"awesomeProject/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoute(init *config.Initialize, g *gin.RouterGroup) *gin.RouterGroup {
	//g.GET("", middleware.AuthMiddleware(), init.CustomerCtrl.GetAll)
	g.POST("", init.UserCtrl.Create)
	g.PUT("/:id", middleware.AuthMiddleware(), init.UserCtrl.Update)
	return g
}
