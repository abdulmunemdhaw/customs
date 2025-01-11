package controller

import (
	"github.com/abdlmunemdhaw/customs/model"
	Services "github.com/abdlmunemdhaw/customs/services"
	"github.com/gin-gonic/gin"
)

type SectorController struct {
	SectorService *Services.SectorService;
}

func NewSectorController(sectorService *Services.SectorService) *SectorController{
	return &SectorController{SectorService: sectorService}
}

func (c *SectorController) CreateSector(ctx *gin.Context){
	var sector model.Sector;
	if err := ctx.ShouldBindJSON(&sector); err != nil {
		ctx.JSON(400,gin.H{
			"msg":err.Error(),
		})
		return;
	}
	sector.SectorId = c.SectorService.GetMaxId() + 1;
	if err := c.SectorService.CreateSector(&sector); err != nil {
		ctx.JSON(400,gin.H{
			"msg":err.Error(),
		})
		return;
	}
	ctx.JSON(200,gin.H{
		"msg":"create sector succfully",
	})
}

func (c *SectorController) GetAllSector(ctx *gin.Context){
	sectors,err := c.SectorService.GetAllSector();
	if (err != nil){
		ctx.JSON(400,gin.H{
			"msg":err.Error(),
		})
		return;
	}
	ctx.JSON(200,gin.H{
		"msg":sectors,
	})
}

func (c *SectorController) UpdateSector(ctx *gin.Context){
	var sector_id string = ctx.Params.ByName("id");
	if sector_id == "" {
		ctx.JSON(400,gin.H{
			"msg":"undefined sector id",
		})
		return;
	}
	var sector map[string] interface{};
	if err := ctx.ShouldBindJSON(&sector); err != nil {
		ctx.JSON(400,gin.H{
			"msg":err.Error(),
		})
		return;
	}
	if err := c.SectorService.UpdateSector(sector_id,sector); err != nil {
		ctx.JSON(400,gin.H{
			"msg":err.Error(),
		})
		return;
	}
	ctx.JSON(200,gin.H{
		"msg":"update data succfully",
	})
}

func (c *SectorController) DeleteSector(ctx *gin.Context) {
	var sector_id string = ctx.Params.ByName("id");
	if (sector_id == ""){
		ctx.JSON(400,gin.H{
			"msg": "undefined sector_id",
		})
		return;
	}
	if err := c.SectorService.DeleteSector(sector_id); err != nil {
		ctx.JSON(400,gin.H{
			"msg": err.Error(),
		})
		return;
	}

	ctx.JSON(200,gin.H{
		"msg": "delete succfully",
	})
}
