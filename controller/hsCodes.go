package controller

import (
	"github.com/abdlmunemdhaw/customs/model"
	Services "github.com/abdlmunemdhaw/customs/services"
	"github.com/gin-gonic/gin"
)

type HsCodeController struct {
	HsCodeService *Services.HsCodeService
}

func NewHsCodeController(service *Services.HsCodeService) *HsCodeController {
	return &HsCodeController{HsCodeService: service}
}
func (c *HsCodeController) CreateHsCode(ctx *gin.Context) {
	var hs_code model.HsCode

	if err := ctx.ShouldBindJSON(&hs_code); err != nil {
		ctx.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}
	hs_code.HsCodeId = c.HsCodeService.GetMaxId() + 1
	if err := c.HsCodeService.CreateHsCode(&hs_code); err != nil {
		ctx.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"msg": "create sector succfully",
	})
}
func (c *HsCodeController) GetHsCodes(ctx *gin.Context) {
	hs_codes, err := c.HsCodeService.GetAllHsCodes()
	if err != nil {
		ctx.JSON(400, gin.H{
			"msg": err.Error(),
		})
	}
	ctx.JSON(200, gin.H{
		"msg": hs_codes,
	})
}
func (c *HsCodeController) UpdateHsCode(ctx *gin.Context) {
	var hs_code_id string = ctx.Params.ByName("id")
	if hs_code_id == "" {
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
	if err := c.HsCodeService.UpdateHsCode(hs_code_id, hs_code); err != nil {
		ctx.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"msg": "update data succfully",
	})
}
func (c *HsCodeController) DeleteHsCode(ctx *gin.Context) {
	var hs_code_id string = ctx.Params.ByName("id")
	if hs_code_id == "" {
		ctx.JSON(400, gin.H{
			"msg": "undefined hs_code_id",
		})
		return
	}
	if err := c.HsCodeService.DeleteHsCode(hs_code_id); err != nil {
		ctx.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"msg": "delete succfully",
	})
}
