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
		fmt.Printf("Erro ao buscar produtos: %s\n", err)
		return model.Product{}, err
	}

	product.ID = id
	return product, nil
}

func (p_usecase *ProductUsecase) FindById(id int) (model.Product, error) {
	product, err := p_usecase.productRepository.FindById(id)
	if err != nil {
		fmt.Printf("Erro ao buscar produto por id: %s\n", err)
		return model.Product{}, nil
	}

	return product, nil
}
