package routes

import (
	"awesomeProject/config"
	"github.com/gin-gonic/gin"
)

func SignatureRoute(init *config.Initialize, g *gin.RouterGroup) *gin.RouterGroup {
	g.GET("", init.SignatureCtrl.GetAll)
	g.POST("", init.SignatureCtrl.Create)
	g.GET("/:id", init.SignatureCtrl.GetByID)
	g.PUT("/:id", init.SignatureCtrl.Update)
	g.DELETE("/:id", init.SignatureCtrl.Delete)
	return g
}
