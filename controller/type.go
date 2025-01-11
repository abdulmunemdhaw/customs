package controller

import (
	"github.com/abdlmunemdhaw/customs/model"
	Services "github.com/abdlmunemdhaw/customs/services"
	"github.com/gin-gonic/gin"
)

type TypeController struct {
	TypeService *Services.TypeService
}

func NewTypeController(typeService *Services.TypeService) *TypeController {
	return &TypeController{TypeService: typeService}
}

func (c *TypeController) CreateType(ctx *gin.Context){
	var typeTable model.Type;
	if err:= ctx.ShouldBindJSON(&typeTable); err != nil {
		ctx.JSON(400,gin.H{
			"msg": err.Error(),
		})
		return;
	}
	typeTable.TypeId = c.TypeService.GetMaxId() + 1;
	if err := c.TypeService.CreateType(&typeTable); err != nil {
		ctx.JSON(400,gin.H{
			"msg": err.Error(),
		})
		return;
	}
	ctx.JSON(201,gin.H{
		"msg": "create succfully",
	})
}

func (c *TypeController) GetTypes(ctx *gin.Context){
	types,err := c.TypeService.GetAllTypes();
	if  err != nil {
		ctx.JSON(400,gin.H{
			"msg": err.Error(),
		})
		return;
	}
	ctx.JSON(200,gin.H{
		"msg": types,
	})

}

func (c *TypeController) UpdateType(ctx *gin.Context){
	var type_table map[string] interface{};
	var type_id string = ctx.Params.ByName("id");

	if type_id == "" {
		ctx.JSON(400,gin.H{
			"msg": "undefined message",
		})
		return;
	}

	if err := ctx.ShouldBindJSON(&type_table); err != nil {
		ctx.JSON(400,gin.H{
			"msg": err.Error(),
		})
		return
	}

	if err := c.TypeService.UpdateType(type_id,type_table); err != nil {
		ctx.JSON(400,gin.H{
			"msg": err.Error(),
		})
		return;
	}

	ctx.JSON(200,gin.H{
		"msg": "update succfully",
	})
}

func (c *TypeController) DeleteType(ctx *gin.Context) {
	var type_id string = ctx.Params.ByName("id");
	if (type_id == ""){
		ctx.JSON(400,gin.H{
			"msg": "undefined type_id",
		})
		return;
	}
	if err := c.TypeService.DeleteType(type_id); err != nil {
		ctx.JSON(400,gin.H{
			"msg": err.Error(),
		})
		return;
	}

	ctx.JSON(200,gin.H{
		"msg": "delete succfully",
	})
}