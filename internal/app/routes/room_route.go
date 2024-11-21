package routes

import (
	"awesomeProject/config"
	"github.com/gin-gonic/gin"
)

func RoomRoute(init *config.Initialize, g *gin.RouterGroup) *gin.RouterGroup {
	g.GET("", init.RoomCtrl.GetAll)
	g.POST("", init.RoomCtrl.Create)
	g.GET("/:id", init.RoomCtrl.GetByID)
	g.PUT("/:id", init.RoomCtrl.Update)
	g.DELETE("/:id", init.RoomCtrl.Delete)
	g.PUT("/:id/status", init.RoomCtrl.UpdateStatus)
	return g
}
