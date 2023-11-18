package api

import "github.com/gin-gonic/gin"

type BrandApi interface {
	Create(ctx *gin.Context)
	Get(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	GetById(ctx *gin.Context)
}
