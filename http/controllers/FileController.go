package controllers

import (
	"Altea/services/managers"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

type FileController struct {
	fileManager *managers.FileManager
}

func NewFileController(fileManager *managers.FileManager) *FileController {
	return &FileController{fileManager: fileManager}
}

func (c *FileController) Get(ctx *gin.Context) {
	path := ctx.Param("path")
	path = filepath.Clean(path)
	content, itemType, err := c.fileManager.GetPath(path)

	if err != nil {
		_ = ctx.AbortWithError(500, err)
		return
	}
	if itemType.IsNotFound() {
		ctx.JSON(404, gin.H{"error": "not found"})
		return
	}

	if itemType.IsFile() {
		ctx.Writer.WriteHeader(200)
		ctx.Writer.Write(content.([]byte))
		return
	}
	ctx.JSON(200, content)
}

func (c *FileController) Put(ctx *gin.Context) {
	path := ctx.Param("path")
	path = filepath.Clean(path)
	err := c.fileManager.WritePath(path, ctx.Request.Body)
	if err != nil {
		_ = ctx.AbortWithError(500, err)
		return
	}
	c.Get(ctx)
}

func (c *FileController) Delete(ctx *gin.Context) {
	path := ctx.Param("path")
	path = filepath.Clean(path)

	err := c.fileManager.DeletePath(path)
	if err != nil {
		_ = ctx.AbortWithError(500, err)
		return
	}
	ctx.Writer.WriteHeader(204)
	return
}
