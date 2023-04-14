package repository

import (
	"challenge-chapter-3-sesi-3/mocks/entity"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (repository *ProductRepositoryMock) FindById(id uint) *entity.Products {
	arguments := repository.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil
	} else {
		product := arguments.Get(0).(entity.Products)
		return &product
	}

}

func (repository *ProductRepositoryMock) FindAll() []*entity.Products {

	arguments := repository.Mock.Called()

	if arguments.Get(0) == nil {
		return nil
	}

	product := arguments.Get(0).([]*entity.Products)
	return product

}
