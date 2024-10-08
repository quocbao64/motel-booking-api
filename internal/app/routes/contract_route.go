package routes

import (
	"awesomeProject/config"
	"github.com/gin-gonic/gin"
)

func ContractRoute(init *config.Initialize, g *gin.RouterGroup) *gin.RouterGroup {
	g.GET("", init.ContractCtrl.GetAll)
	g.POST("", init.ContractCtrl.Create)
	g.GET("/:id", init.ContractCtrl.GetByID)
	g.PUT("/:id", init.ContractCtrl.Update)
	g.DELETE("/:id", init.ContractCtrl.Delete)
	return g
}
