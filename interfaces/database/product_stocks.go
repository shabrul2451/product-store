package database

import "Product-Store/model"

type ProductStockRepository interface {
	Get() []model.ProductStock
	Update(dto model.ProductStock) error
	Delete(id string) error
	GetById(id string) model.ProductStock
}
