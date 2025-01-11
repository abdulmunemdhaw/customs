package controller

import (
	"github.com/abdlmunemdhaw/customs/model"
	Services "github.com/abdlmunemdhaw/customs/services"
	"github.com/gin-gonic/gin"
)

type MunicipalityController struct {
	MunicipalitiyService *Services.MunicipalityService;
}

func NewMunicipalityController(service *Services.MunicipalityService) *MunicipalityController {
	return &MunicipalityController{MunicipalitiyService: service};
}

func (obj *MunicipalityController) GetMunicipalities(ctx *gin.Context){
	municipalities,err := obj.MunicipalitiyService.GetAllMunicipality();
	if err != nil {
		ctx.JSON(400,gin.H{
			"msg":err.Error(),
		})
	}
	ctx.JSON(200,gin.H{
		"msg":municipalities,
	})
}
func (obj *MunicipalityController) UpdateMunicipality(ctx *gin.Context){
	var updates map[string] interface{};

	if err := ctx.ShouldBindJSON(&updates); err != nil {
		ctx.JSON(400,gin.H{
			"msg":err.Error(),
		})
		return;
	}
	var municipality_id string =ctx.Params.ByName("id");
	err := obj.MunicipalitiyService.UpdateMunicipality(municipality_id,updates);
	if err != nil {
		ctx.JSON(400,gin.H{
			"msg":err.Error(),
		})
		return;
	}
	ctx.JSON(200,gin.H{
		"msg":"updated succfully",
	})
}
func (obj *MunicipalityController) CreateMunicipality(ctx *gin.Context){
	var municipality model.Municipality;
	if err := ctx.ShouldBindJSON(&municipality); err != nil {
		ctx.JSON(400,gin.H{
			"msg":err.Error(),
		})
		return;
	}
	municipality.MunicipalityId = obj.MunicipalitiyService.GetMaxId() + 1;
	if err := obj.MunicipalitiyService.CreateMunicipality(&municipality); err != nil {
		ctx.JSON(400,gin.H{
			"msg":err.Error(),
		})
		return;
	}
	ctx.JSON(201,gin.H{
		"msg":"add succfully",
	})
}
func (obj *MunicipalityController) DeleteMunicipality(ctx *gin.Context){
	var municipality model.Municipality;
	var municipality_id string = ctx.Params.ByName("id");
	if municipality_id == "" {
		ctx.JSON(400,gin.H{
			"msg":"undefined id",
		})
		return;
	}
	if err := obj.MunicipalitiyService.DeleteMunicipality(municipality_id,&municipality); err != nil {
		ctx.JSON(400,gin.H{
			"msg":err.Error(),
		})
		return;
	}
	ctx.JSON(200,gin.H{
		"msg":"municipality delete succfully",
	})
}
