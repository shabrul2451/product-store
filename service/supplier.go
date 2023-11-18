package service

import (
	"Product-Store/interfaces/database"
	"Product-Store/interfaces/service"
	"Product-Store/model"
	"github.com/google/uuid"
	"time"
)

type supplierService struct {
	sr database.SupplierRepository
}

func (s supplierService) Create(dto model.Supplier) error {
	dto.Id = uuid.New().String()
	dto.CreatedAt = time.Now()
	dto.StatusId = 1
	return s.sr.Create(dto)
}

func (s supplierService) Get() []model.Supplier {
	return s.sr.Get()
}

func (s supplierService) Update(dto model.Supplier) error {
	dto.CreatedAt = time.Now()
	return s.sr.Update(dto)
}

func (s supplierService) Delete(id string) error {
	return s.sr.Delete(id)
}

func (s supplierService) GetById(id string) model.Supplier {
	return s.sr.GetById(id)
}

func NewSupplierService(repository database.SupplierRepository) service.SupplierServiceInterface {
	return &supplierService{sr: repository}
}
