package service

import (
	"Product-Store/interfaces/database"
	"Product-Store/interfaces/service"
	"Product-Store/model"
	"time"
)

type productStockService struct {
	psr database.ProductStockRepository
}

func (p productStockService) Get() []model.ProductStock {
	return p.psr.Get()
}

func (p productStockService) Update(dto model.ProductStock) error {
	dto.UpdatedAt = time.Now()
	return p.psr.Update(dto)
}

func (p productStockService) Delete(id string) error {
	return p.psr.Delete(id)
}

func (p productStockService) GetById(id string) model.ProductStock {
	return p.psr.GetById(id)
}

func NewProductStockService(repository database.ProductStockRepository) service.ProductStockServiceInterface {
	return &productStockService{psr: repository}
}
