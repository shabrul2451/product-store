package service

import (
	"Product-Store/interfaces/database"
	"Product-Store/interfaces/service"
	"Product-Store/model"
	"github.com/google/uuid"
	"time"
)

type categoriesService struct {
	cr database.CategoriesRepository
}

func (c categoriesService) Create(dto model.Category) error {
	dto.Id = uuid.New().String()
	dto.CreatedAt = time.Now()
	dto.StatusId = 1
	return c.cr.Create(dto)
}

func (c categoriesService) Get() []model.Category {
	return c.cr.Get()
}

func (c categoriesService) Update(dto model.Category) error {
	dto.CreatedAt = time.Now()
	return c.cr.Update(dto)
}

func (c categoriesService) Delete(id string) error {
	return c.cr.Delete(id)
}

func (c categoriesService) GetById(id string) model.Category {
	return c.cr.GetById(id)
}

func NewCategoriesService(repository database.CategoriesRepository) service.CategoriesServiceInterface {
	return &categoriesService{cr: repository}
}
