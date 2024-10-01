package routes

import (
	"awesomeProject/config"
	"github.com/gin-gonic/gin"
)

func ServiceDemandRoute(init *config.Initialize, g *gin.RouterGroup) *gin.RouterGroup {
	g.GET("", init.ServicesDemandCtrl.GetAll)
	g.POST("", init.ServicesDemandCtrl.Create)
	g.GET("/:id", init.ServicesDemandCtrl.GetByID)
	g.PUT("/:id", init.ServicesDemandCtrl.Update)
	g.DELETE("/:id", init.ServicesDemandCtrl.Delete)
	return g
}
