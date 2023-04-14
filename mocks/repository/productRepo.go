package repository

import "challenge-chapter-3-sesi-3/mocks/entity"

type ProductRepository interface {
	FindById(id uint) *entity.Products
	FindAll() []*entity.Products
}
