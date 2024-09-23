package routes

import (
	"awesomeProject/config"
	"github.com/gin-gonic/gin"
)

func ServiceRoute(init *config.Initialize, g *gin.RouterGroup) *gin.RouterGroup {
	g.GET("", init.ServiceCtrl.GetAll)
	g.POST("", init.ServiceCtrl.Create)
	g.GET("/:id", init.ServiceCtrl.GetByID)
	g.PUT("/:id", init.ServiceCtrl.Update)
	g.DELETE("/:id", init.ServiceCtrl.Delete)
	return g
}
