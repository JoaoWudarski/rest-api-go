package usecase

import (
	"fmt"
	"rest-api-go/model"
	"rest-api-go/repository"
)

type ProductUsecase struct {
	productRepository repository.ProductRepository
}

func NewProductUsecase(productRepository repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		productRepository: productRepository,
	}
}

func (p_usecase *ProductUsecase) GetProducts() ([]model.Product, error) {
	return p_usecase.productRepository.GetAllProducts()
}

func (p_usecase *ProductUsecase) SaveProducts(product model.Product) (model.Product, error) {
	id, err := p_usecase.productRepository.SaveProduct(product)
	if err != nil || id == -1 {
		fmt.Printf("Erro ao buscar produtos: %s", err)
		return model.Product{}, err
	}

	product.ID = id
	return product, nil
}
