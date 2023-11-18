package service

import "Product-Store/model"

type CategoriesServiceInterface interface {
	Create(dto model.Category) error
	Get() []model.Category
	Update(dto model.Category) error
	Delete(id string) error
	GetById(id string) model.Category
}
