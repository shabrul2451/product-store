package database

import "Product-Store/model"

type SupplierRepository interface {
	Create(dto model.Supplier) error
	Get() []model.Supplier
	Update(dto model.Supplier) error
	Delete(id string) error
	GetById(id string) model.Supplier
}
