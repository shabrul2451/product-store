package api

import (
	"Product-Store/service"
	"Product-Store/utils/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(LoggerMiddleware())
	brandGroupV1 := router.Group("/api/v1/brands")
	{
		brandApi := NewBrandApi(service.NewBrandService(database.NewBrandRepository()))
		brandGroupV1.GET("", brandApi.Get)
		brandGroupV1.POST("", brandApi.Create)
		brandGroupV1.PUT("/:id", brandApi.Update)
		brandGroupV1.DELETE("/:id", brandApi.Delete)
		brandGroupV1.GET("/:id", brandApi.GetById)
	}

	productGroupV1 := router.Group("/api/v1/products")
	{
		productApi := NewProductApi(service.NewProductService(database.NewProductRepository(), database.NewSupplierRepository(), database.NewProductStockRepository(), database.NewBrandRepository(), database.NewCategoriesRepository()))
		productGroupV1.GET("/by_filters", productApi.GetByFilters)
		productGroupV1.POST("", productApi.Create)
	}

	categoryGroupV1 := router.Group("/api/v1/categories")
	{
		categoriesApi := NewCategoriesApi(service.NewCategoriesService(database.NewCategoriesRepository()))
		categoryGroupV1.GET("/tree", categoriesApi.GetTree)
	}

	return router
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		elapsed := time.Since(start)
		fmt.Printf("%s %s %s %s\n",
			c.Request.Method,
			c.Request.RequestURI,
			c.Writer.Status(),
			elapsed,
		)
	}
}
