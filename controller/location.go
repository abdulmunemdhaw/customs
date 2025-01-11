package controller

import (
	"github.com/abdlmunemdhaw/customs/model"
	Services "github.com/abdlmunemdhaw/customs/services"
	"github.com/gin-gonic/gin"
)

type LocationConroller struct {
	LocationService *Services.LocationService
}

func NewLocationConroller(service *Services.LocationService) *LocationConroller {
	return &LocationConroller{LocationService: service}
}
func (c *LocationConroller) CreateLocation(ctx *gin.Context) {
	var location model.Location
	if err := ctx.ShouldBindJSON(&location); err != nil {
		ctx.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}
	location.LocationId = c.LocationService.GetMaxId() + 1
	if err := c.LocationService.CreateLocation(&location); err != nil {
		ctx.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"msg": "create location succfully",
	})
}
func (c *LocationConroller) GetLocations(ctx *gin.Context) {
	location, err := c.LocationService.GetLocations()
	if err != nil {
		ctx.JSON(400, gin.H{
			"msg": err.Error(),
		})
	}
	ctx.JSON(200, gin.H{
		"msg": location,
	})
}
func (c *LocationConroller) UpdateLocation(ctx *gin.Context) {
	var location_id string = ctx.Params.ByName("id")
	if location_id == "" {
		ctx.JSON(400, gin.H{
			"msg": "undefined hs_code id",
		})
		return
	}
	var hs_code map[string]interface{}
	if err := ctx.ShouldBindJSON(&hs_code); err != nil {
		ctx.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}
	if err := c.LocationService.UpdateLocation(location_id, hs_code); err != nil {
		ctx.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"msg": "update data succfully",
	})
}
func (c *LocationConroller) DeleteLocation(ctx *gin.Context) {
	var location_id string = ctx.Params.ByName("id")
	if location_id == "" {
		ctx.JSON(400, gin.H{
			"msg": "undefined location",
		})
		return
	}
	if err := c.LocationService.DeleteLocation(location_id); err != nil {
		ctx.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"msg": "delete succfully",
	})
}
