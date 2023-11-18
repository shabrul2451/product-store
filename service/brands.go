package service

import (
	"Product-Store/interfaces/database"
	"Product-Store/interfaces/service"
	"Product-Store/model"
	"time"
)

type brandService struct {
	br database.BrandRepository
}

func (b brandService) Get() []model.Brand {
	return b.br.Get()
}

func (b brandService) Update(dto model.Brand) error {
	dto.CreatedAt = time.Now()
	return b.br.Update(dto)
}

func (b brandService) Delete(id string) error {
	return b.br.Delete(id)
}

func (b brandService) GetById(id string) model.Brand {
	return b.br.GetById(id)
}

func NewBrandService(repository database.BrandRepository) service.BrandsServiceInterface {
	return &brandService{br: repository}
}
