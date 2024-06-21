package main

import (
	"rest-api-go/controller"
	"rest-api-go/db"
	"rest-api-go/repository"
	"rest-api-go/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()
	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	productRepository := repository.NewProductRepository(dbConnection)
	productUsecase := usecase.NewProductUsecase(productRepository)
	productController := controller.NewProductController(productUsecase)

	server.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"health": "UP",
		})
	})

	server.GET("/products", productController.GetProducts)
	server.POST("/product", productController.Save)
	server.GET("/product/:productId", productController.FindById)

	server.Run(":8080")
}
