package routes

import (
	"awesomeProject/config"
	"github.com/gin-gonic/gin"
)

func AuthRoute(init *config.Initialize, g *gin.RouterGroup) *gin.RouterGroup {
	g.POST("/login", init.AuthCtrl.Login)
	g.POST("/register", init.AuthCtrl.Register)
	return g

}
