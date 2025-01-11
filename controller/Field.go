package controller

import (
	"github.com/abdlmunemdhaw/customs/model"
	Services "github.com/abdlmunemdhaw/customs/services"
	"github.com/gin-gonic/gin"
)

type FieldController struct {
	FieldServices *Services.FieldService;
}

func NewFieldController(service *Services.FieldService) *FieldController {
	return &FieldController{FieldServices: service};
}
func (c *FieldController) CreateField(ctx *gin.Context){
	var field model.Field;
	
	if err := ctx.ShouldBindJSON(&field); err != nil {
		ctx.JSON(400,gin.H{
			"msg":err.Error(),
		})
		return;
	}
	field.FieldId = c.FieldServices.GetMaxId() + 1;
	if err := c.FieldServices.CreateSector(&field); err != nil {
		ctx.JSON(400,gin.H{
			"msg":err.Error(),
		})
		return;
	}
	ctx.JSON(200,gin.H{
		"msg":"create sector succfully",
	})
}
func (c *FieldController) GetFields(ctx *gin.Context){
	fields,err := c.FieldServices.GetAllFields();
	if err != nil {
		ctx.JSON(400,gin.H{
			"msg":err.Error(),
		})
	}
	ctx.JSON(200,gin.H{
		"msg":fields,
	})
}
func (c *FieldController) UpdateField(ctx *gin.Context){
	var field_id string = ctx.Params.ByName("id");
	if field_id == "" {
		ctx.JSON(400,gin.H{
			"msg":"undefined field id",
		})
		return;
	}
	var field map[string] interface{};
	if err := ctx.ShouldBindJSON(&field); err != nil {
		ctx.JSON(400,gin.H{
			"msg":err.Error(),
		})
		return;
	}
	if err := c.FieldServices.UpdateField(field_id,field); err != nil {
		ctx.JSON(400,gin.H{
			"msg":err.Error(),
		})
		return;
	}
	ctx.JSON(200,gin.H{
		"msg":"update data succfully",
	})
}
func (c *FieldController) DeleteField(ctx *gin.Context) {
	var field_id string = ctx.Params.ByName("id");
	if (field_id == ""){
		ctx.JSON(400,gin.H{
			"msg": "undefined field_id",
		})
		return;
	}
	if err := c.FieldServices.DeleteField(field_id); err != nil {
		ctx.JSON(400,gin.H{
			"msg": err.Error(),
		})
		return;
	}

	ctx.JSON(200,gin.H{
		"msg": "delete succfully",
	})
}
