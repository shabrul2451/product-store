package service

import "Product-Store/model"

type BrandsServiceInterface interface {
	Get() []model.Brand
	Update(dto model.Brand) error
	Delete(id string) error
	GetById(id string) model.Brand
}
