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
	g.PUT("/balance", middleware.AuthMiddleware(), init.UserCtrl.UpdateBalance)
	g.GET("/:id", middleware.AuthMiddleware(), init.UserCtrl.GetByID)
	g.GET("", init.UserCtrl.GetAll)
	return g
}
