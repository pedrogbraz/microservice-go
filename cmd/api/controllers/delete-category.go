package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pedrogbraz/microservice-go/internal/repositories"
	"github.com/pedrogbraz/microservice-go/internal/usecases"
)

func DeleteCategory(ctx *gin.Context, repository repositories.ICategoryRepository) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, (gin.H{
			"success": false,
			"error":   err.Error()}))
		return
	}

	useCases := usecases.NewDeleteCategoryUseCase(repository)
	err = useCases.Execute(id)

	if err != nil {
		if errors.Is(err, repositories.ErrCategoryNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, (gin.H{
				"success": false,
				"error":   err.Error()}))
			return
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, (gin.H{
			"success": false,
			"error":   err.Error()}))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
