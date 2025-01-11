package controller

import (
	"github.com/abdlmunemdhaw/customs/model"
	Services "github.com/abdlmunemdhaw/customs/services"
	"github.com/gin-gonic/gin"
)

type DocumentController struct {
	DocumentService *Services.DocumentService
}

func NewDocumentController(service *Services.DocumentService) *DocumentController {
	return &DocumentController{DocumentService: service}
}
func (c *DocumentController) CreateDocument(ctx *gin.Context) {
	var document model.Document

	if err := ctx.ShouldBindJSON(&document); err != nil {
		ctx.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}
	document.DocumentId = c.DocumentService.GetMaxId() + 1
	if err := c.DocumentService.CreateDocument(&document); err != nil {
		ctx.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"msg": "create document succfully",
	})
}
func (c *DocumentController) GetDocument(ctx *gin.Context) {
	document, err := c.DocumentService.GetAllDocuments()
	if err != nil {
		ctx.JSON(400, gin.H{
			"msg": err.Error(),
		})
	}
	ctx.JSON(200, gin.H{
		"msg": document,
	})
}
func (c *DocumentController) UpdateDocument(ctx *gin.Context) {
	var document_id string = ctx.Params.ByName("id")
	if document_id == "" {
		ctx.JSON(400, gin.H{
			"msg": "undefined document id",
		})
		return
	}
	var document map[string]interface{}
	if err := ctx.ShouldBindJSON(&document); err != nil {
		ctx.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}
	if err := c.DocumentService.UpdateDocument(document_id, document); err != nil {
		ctx.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"msg": "update data succfully",
	})
}
func (c *DocumentController) DeleteDocument(ctx *gin.Context) {
	var document_id string = ctx.Params.ByName("id")
	if document_id == "" {
		ctx.JSON(400, gin.H{
			"msg": "undefined document_id",
		})
		return
	}
	if err := c.DocumentService.DeleteDocument(document_id); err != nil {
		ctx.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"msg": "delete succfully",
	})
}
