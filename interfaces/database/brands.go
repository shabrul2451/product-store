package database

import "Product-Store/model"

type BrandRepository interface {
	Get() []model.Brand
	Update(dto model.Brand) error
	Delete(id string) error
	GetById(id string) model.Brand
}
