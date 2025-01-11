package controller

import (
	"github.com/abdlmunemdhaw/customs/model"
	Services "github.com/abdlmunemdhaw/customs/services"
	"github.com/gin-gonic/gin"
)

type EconomicBulletinController struct {
	EconomicBulletinService *Services.EconomicBulletinService
}

func NewEconomicBulletinController(service *Services.EconomicBulletinService) *EconomicBulletinController {
	return &EconomicBulletinController{EconomicBulletinService: service}
}
func (c *EconomicBulletinController) CreateEconomicBulletin(ctx *gin.Context) {
	var economic_bulletin model.EconomicBulletin

	if err := ctx.ShouldBindJSON(&economic_bulletin); err != nil {
		ctx.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}
	economic_bulletin.EconomicBulletinId = c.EconomicBulletinService.GetMaxId() + 1
	if err := c.EconomicBulletinService.CreateEconomicBulletin(&economic_bulletin); err != nil {
		ctx.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"msg": "create economic_bulletin succfully",
	})
}
func (c *EconomicBulletinController) GetEconomicBulletins(ctx *gin.Context) {
	category, err := c.EconomicBulletinService.GetAllEconomicBulletins()
	if err != nil {
		ctx.JSON(400, gin.H{
			"msg": err.Error(),
		})
	}
	ctx.JSON(200, gin.H{
		"msg": category,
	})
}
func (c *EconomicBulletinController) UpdateEconomicBulletin(ctx *gin.Context) {
	var economic_bulletin_id string = ctx.Params.ByName("id")
	if economic_bulletin_id == "" {
		ctx.JSON(400, gin.H{
			"msg": "undefined  economic_bulletin_id",
		})
		return
	}
	var economic_bulletin map[string]interface{}
	if err := ctx.ShouldBindJSON(&economic_bulletin); err != nil {
		ctx.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}
	if err := c.EconomicBulletinService.UpdateEconomicBulletin(economic_bulletin_id, economic_bulletin); err != nil {
		ctx.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"msg": "update data succfully",
	})
}
func (c *EconomicBulletinController) DeleteEconomicBulletin(ctx *gin.Context) {
	var economic_bulletin_id string = ctx.Params.ByName("id")
	if economic_bulletin_id == "" {
		ctx.JSON(400, gin.H{
			"msg": "undefined category_id",
		})
		return
	}
	if err := c.EconomicBulletinService.DeleteEconomicBulletin(economic_bulletin_id); err != nil {
		ctx.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"msg": "delete succfully",
	})
}
