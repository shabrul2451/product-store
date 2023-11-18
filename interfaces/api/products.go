package api

import "github.com/gin-gonic/gin"

type ProductApi interface {
	Create(ctx *gin.Context)
	Get(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
	GetById(ctx *gin.Context)
	GetByFilters(ctx *gin.Context)
}
