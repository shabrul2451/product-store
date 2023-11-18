package database

import (
	"Product-Store/common"
	"Product-Store/model"
)

type ProductRepository interface {
	Create(dto model.Product) error
	Get() []model.Product
	Update(dto model.Product) error
	Delete(id string) error
	GetBySupplierIdAndProductName(id, name string) model.Product
	GetById(id string) model.Product
	GetByFilter(filter common.ProductFilter) ([]model.Product, int64)
}
