package controller

import (
	"github.com/abdlmunemdhaw/customs/model"
	Services "github.com/abdlmunemdhaw/customs/services"
	"github.com/gin-gonic/gin"
)

type BusinessController struct {
	BusinessService *Services.BusinessService
}

func NewBusinessController(service *Services.BusinessService) *BusinessController{
	return &BusinessController{BusinessService: service};
}
func (c *BusinessController) CreateBusiness(ctx *gin.Context){
	var business model.Business;
	

	if err := ctx.ShouldBindJSON(&business); err != nil {
		ctx.JSON(400,gin.H{
			"msg":err.Error(),
		})
		return;
	}
	business.BusinessId = c.BusinessService.GetMaxId() + 1;
	if err := c.BusinessService.CreateBusiness(&business); err != nil {
		ctx.JSON(400,gin.H{
			"msg":err.Error(),
		})
		return;
	}
	ctx.JSON(200,gin.H{
		"msg":"create sector succfully",
	})
}
func (c *BusinessController) GetBusiness(ctx *gin.Context){
	businesses,err := c.BusinessService.GetAllBusinesses();
	if err != nil {
		ctx.JSON(400,gin.H{
			"msg":err.Error(),
		})
	}
	ctx.JSON(200,gin.H{
		"msg":businesses,
	})
}
func (c *BusinessController) UpdateBusiness(ctx *gin.Context){
	var business_id string = ctx.Params.ByName("id");
	if business_id == "" {
		ctx.JSON(400,gin.H{
			"msg":"undefined field id",
		})
		return;
	}
	var business map[string] interface{};
	if err := ctx.ShouldBindJSON(&business); err != nil {
		ctx.JSON(400,gin.H{
			"msg":err.Error(),
		})
		return;
	}
	if err := c.BusinessService.UpdateBusiness(business_id,business); err != nil {
		ctx.JSON(400,gin.H{
			"msg":err.Error(),
		})
		return;
	}
	ctx.JSON(200,gin.H{
		"msg":"update data succfully",
	})
}
func (c *BusinessController) DeleteBusines(ctx *gin.Context) {
	var business_id string = ctx.Params.ByName("id");
	if (business_id == ""){
		ctx.JSON(400,gin.H{
			"msg": "undefined business_id",
		})
		return;
	}
	if err := c.BusinessService.DeleteBusines(business_id); err != nil {
		ctx.JSON(400,gin.H{
			"msg": err.Error(),
		})
		return;
	}

	ctx.JSON(200,gin.H{
		"msg": "delete succfully",
	})
}
