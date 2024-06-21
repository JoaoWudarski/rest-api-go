package usecase

import (
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
