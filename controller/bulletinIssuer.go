package controller

import (
	"github.com/abdlmunemdhaw/customs/model"
	Services "github.com/abdlmunemdhaw/customs/services"
	"github.com/gin-gonic/gin"
)

type BulletinIssuerController struct {
	BulletinIssuerService *Services.BulletinIssuerService
}

func NewBulletinIssuerController(service *Services.BulletinIssuerService) *BulletinIssuerController{
	return &BulletinIssuerController{BulletinIssuerService: service};
}
func (c *BulletinIssuerController) CreateBusiness(ctx *gin.Context){
	var bulletin_issuer model.BulletinIssuer;
	

	if err := ctx.ShouldBindJSON(&bulletin_issuer); err != nil {
		ctx.JSON(400,gin.H{
			"msg":err.Error(),
		})
		return;
	}
	bulletin_issuer.BulletinIssuerId = c.BulletinIssuerService.GetMaxId() + 1;
	if err := c.BulletinIssuerService.CreateBulletinIssuer(&bulletin_issuer); err != nil {
		ctx.JSON(400,gin.H{
			"msg":err.Error(),
		})
		return;
	}
	ctx.JSON(200,gin.H{
		"msg":"create bulletin_issuer succfully",
	})
}
func (c *BulletinIssuerController) GetBulletinIssueres(ctx *gin.Context){
	bulletin_issuer,err := c.BulletinIssuerService.GetAllBulletinIssueres();
	if err != nil {
		ctx.JSON(400,gin.H{
			"msg":err.Error(),
		})
	}
	ctx.JSON(200,gin.H{
		"msg":bulletin_issuer,
	})
}
func (c *BulletinIssuerController) UpdateBulletinIssuer(ctx *gin.Context){
	var bulletin_issuer_id string = ctx.Params.ByName("id");
	if bulletin_issuer_id == "" {
		ctx.JSON(400,gin.H{
			"msg":"undefined bulletin_issuer_id id",
		})
		return;
	}
	var bulletin_issuer map[string] interface{};
	if err := ctx.ShouldBindJSON(&bulletin_issuer); err != nil {
		ctx.JSON(400,gin.H{
			"msg":err.Error(),
		})
		return;
	}
	if err := c.BulletinIssuerService.UpdateBulletinIssuer(bulletin_issuer_id,bulletin_issuer); err != nil {
		ctx.JSON(400,gin.H{
			"msg":err.Error(),
		})
		return;
	}
	ctx.JSON(200,gin.H{
		"msg":"update data succfully",
	})
}
func (c *BulletinIssuerController) DeleteBulletinIssuer(ctx *gin.Context) {
	var bulletin_issuer_id string = ctx.Params.ByName("id");
	if (bulletin_issuer_id == ""){
		ctx.JSON(400,gin.H{
			"msg": "undefined bulletin_issuer_id ",
		})
		return;
	}
	if err := c.BulletinIssuerService.DeleteBulletinIssuer(bulletin_issuer_id ); err != nil {
		ctx.JSON(400,gin.H{
			"msg": err.Error(),
		})
		return;
	}

	ctx.JSON(200,gin.H{
		"msg": "delete succfully",
	})
}
