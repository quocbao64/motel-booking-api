package routes

import (
	"awesomeProject/config"
	"github.com/gin-gonic/gin"
)

func GeographyRoute(init *config.Initialize, g *gin.RouterGroup) *gin.RouterGroup {
	g.GET("/provinces", init.GeographyCtrl.GetAllProvince)
	g.GET("/provinces/:id", init.GeographyCtrl.GetProvinceByID)
	g.GET("/districts", init.GeographyCtrl.GetAllDistrict)
	g.GET("/districts/:id", init.GeographyCtrl.GetDistrictByID)
	g.GET("/wards", init.GeographyCtrl.GetAllWard)
	g.GET("/wards/:id", init.GeographyCtrl.GetWardByID)
	return g
}
