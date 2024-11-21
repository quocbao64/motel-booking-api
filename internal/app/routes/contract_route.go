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
	g.POST("/liquidity", init.ContractCtrl.Liquidity)
	g.POST("/booking", init.ContractCtrl.CreateFromBookingRequest)
	g.POST("/cancel", init.ContractCtrl.CancelContract)
	return g
}
