package service

import "Product-Store/model"

type SupplierServiceInterface interface {
	Create(dto model.Supplier) error
	Get() []model.Supplier
	Update(dto model.Supplier) error
	Delete(id string) error
	GetById(id string) model.Supplier
}
