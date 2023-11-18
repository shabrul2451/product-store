package service

import (
	"Product-Store/common"
	"Product-Store/model"
)

type ProductServiceInterface interface {
	Create(dto model.ProductDto) error
	Get() []model.Product
	Update(dto model.Product) error
	Delete(id string) error
	GetByFilter(filter common.ProductFilter) ([]model.Product, int64)
}
