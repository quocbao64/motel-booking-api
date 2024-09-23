package routes

import (
	"awesomeProject/config"
	"github.com/gin-gonic/gin"
)

func AddressRoute(init *config.Initialize, g *gin.RouterGroup) *gin.RouterGroup {
	g.GET("/:id", init.AddressCtrl.GetFullAddress)
	return g
}
