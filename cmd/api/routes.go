package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pedrogbraz/microservice-go/cmd/api/controllers"
	"github.com/pedrogbraz/microservice-go/internal/repositories"
)

func CategoryRoutes(router *gin.Engine) {
	// Define category-related routes here
	{
		categoryRoutes := router.Group("/categories")

		inMemoryCategoryRepository := repositories.NewInMemoryCategoryRepository()

		categoryRoutes.POST("/", func(ctx *gin.Context) {
			controllers.CreateCategory(ctx, inMemoryCategoryRepository)
		})

		categoryRoutes.GET("/", func(ctx *gin.Context) {
			controllers.ListCategories(ctx, inMemoryCategoryRepository)
		})

		categoryRoutes.DELETE("/:id", func(ctx *gin.Context) {
			controllers.DeleteCategory(ctx, inMemoryCategoryRepository)
		})
	}
}
