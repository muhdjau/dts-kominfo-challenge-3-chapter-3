package service

import (
	"challenge-chapter-3-sesi-3/mocks/entity"
	"challenge-chapter-3-sesi-3/mocks/repository"
	"errors"
)

type ProductService struct {
	Repository repository.ProductRepository
}

func (service ProductService) GetOneProduct(id uint) (*entity.Products, error) {
	product := service.Repository.FindById(id)
	if product == nil {
		return nil, errors.New("product not found")
	} else {
		return product, nil
	}
}

func (service ProductService) GetAll() ([]*entity.Products, error) {
	product := service.Repository.FindAll()
	if product == nil {
		return nil, errors.New("all product not found")
	} else {
		return product, nil
	}
}
