package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Crie um roteador Gin com os middlewares padrão (logger e recovery)
	router := gin.Default()

	// Defina um endpoint GET simples
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
		})
	})

	// Configure as rotas de categoria
	CategoryRoutes(router)

	// Inicie o servidor na porta 8080 (padrão)
	// O servidor escutará em 0.0.0.0:8080 (localhost:8080 no Windows)
	router.Run(":8080")
}
