package api

import "github.com/gin-gonic/gin"

type SupplierApi interface {
	Create(ctx gin.Context) error
	Get(ctx gin.Context) error
	Delete(ctx gin.Context) error
	Update(ctx gin.Context) error
	GetById(ctx gin.Context) error
}
