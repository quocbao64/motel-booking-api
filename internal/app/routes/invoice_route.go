package routes

import (
	"awesomeProject/config"
	"github.com/gin-gonic/gin"
)

func InvoiceRoute(init *config.Initialize, g *gin.RouterGroup) *gin.RouterGroup {
	g.GET("", init.InvoiceCtrl.GetAll)
	g.POST("", init.InvoiceCtrl.Create)
	g.GET("/:id", init.InvoiceCtrl.GetByID)
	g.PUT("/:id", init.InvoiceCtrl.Update)
	g.DELETE("/:id", init.InvoiceCtrl.Delete)
	return g
}
