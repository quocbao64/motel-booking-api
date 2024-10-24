package routes

import (
	"awesomeProject/config"
	"github.com/gin-gonic/gin"
)

func TransactionRoute(init *config.Initialize, g *gin.RouterGroup) *gin.RouterGroup {
	g.GET("", init.TransactionCtrl.GetAll)
	g.POST("", init.TransactionCtrl.Create)
	g.GET("/:id", init.TransactionCtrl.GetByID)
	g.PUT("/:id", init.TransactionCtrl.Update)
	g.DELETE("/:id", init.TransactionCtrl.Delete)
	return g
}
