package services

import "www.github.com/biskitsx/Redis-Go/repositories"

type catalogService struct {
	productRepo repositories.ProductRepository
}

func NewCatalogService(productRepo repositories.ProductRepository) CatalogService {
	return &catalogService{productRepo}
}

func (s catalogService) GetProducts() (products []Product, err error) {
	productsDB, err := s.productRepo.GetAll()
	if err != nil {
		return nil, err
	}

	for _, product := range productsDB {
		products = append(products, Product(product))
	}
	return products, err
}
