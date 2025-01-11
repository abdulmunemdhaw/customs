package controller

import (
	"github.com/abdlmunemdhaw/customs/model"
	Services "github.com/abdlmunemdhaw/customs/services"
	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	CategoryService *Services.CategoryService
}

func NewCategoryController(service *Services.CategoryService) *CategoryController {
	return &CategoryController{CategoryService: service}
}
func (c *CategoryController) CreateCategory(ctx *gin.Context) {
	var category model.Category

	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}
	category.CategoryId = c.CategoryService.GetMaxId() + 1
	if err := c.CategoryService.CreateCategory(&category); err != nil {
		ctx.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"msg": "create category succfully",
	})
}
func (c *CategoryController) GetCategory(ctx *gin.Context) {
	category, err := c.CategoryService.GetAllCategories()
	if err != nil {
		ctx.JSON(400, gin.H{
			"msg": err.Error(),
		})
	}
	ctx.JSON(200, gin.H{
		"msg": category,
	})
}
func (c *CategoryController) UpdateCategory(ctx *gin.Context) {
	var category_id string = ctx.Params.ByName("id")
	if category_id == "" {
		ctx.JSON(400, gin.H{
			"msg": "undefined category id",
		})
		return
	}
	var category map[string]interface{}
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}
	if err := c.CategoryService.UpdateCategory(category_id, category); err != nil {
		ctx.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"msg": "update data succfully",
	})
}
func (c *CategoryController) DeleteCategory(ctx *gin.Context) {
	var category_id string = ctx.Params.ByName("id")
	if category_id == "" {
		ctx.JSON(400, gin.H{
			"msg": "undefined category_id",
		})
		return
	}
	if err := c.CategoryService.DeleteCategory(category_id); err != nil {
		ctx.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"msg": "delete succfully",
	})
}
