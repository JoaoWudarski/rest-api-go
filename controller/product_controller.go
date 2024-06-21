package controller

import (
	"fmt"
	"net/http"
	"rest-api-go/model"
	"rest-api-go/usecase"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		productUsecase: usecase,
	}
}

func (p_controller *productController) GetProducts(ctx *gin.Context) {

	products, err := p_controller.productUsecase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}

func (p_controller *productController) Save(ctx *gin.Context) {

	var product model.Product
	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	product, err = p_controller.productUsecase.SaveProducts(product)
	if err != nil {
		fmt.Printf("Erro ao salvar produto: %s\n", err)
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusCreated, product)
}
