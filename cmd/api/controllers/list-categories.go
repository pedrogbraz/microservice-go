package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedrogbraz/microservice-go/internal/repositories"
	"github.com/pedrogbraz/microservice-go/internal/usecases"
)

func ListCategories(ctx *gin.Context, repository repositories.ICategoryRepository) {
	useCases := usecases.NewListCategoriesUseCase(repository)
	categories, err := useCases.Execute()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, (gin.H{
			"success": false,
			"error":   err.Error()}))
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success":    true,
		"categories": categories,
	})
}
