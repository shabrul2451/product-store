package service

import (
	"Product-Store/common"
	"Product-Store/interfaces/database"
	"Product-Store/interfaces/service"
	"Product-Store/model"
	"errors"
	"github.com/google/uuid"
	"time"
)

type productService struct {
	productRepository      database.ProductRepository
	supplierRepository     database.SupplierRepository
	productStockRepository database.ProductStockRepository
	brandRepository        database.BrandRepository
	categoryRepository     database.CategoriesRepository
}

func (p productService) GetByFilter(filter common.ProductFilter) ([]model.Product, int64) {
	res, total := p.productRepository.GetByFilter(filter)
	var finalResult []model.Product
	if filter.IsVerified {
		for _, each := range res {
			supplier := p.supplierRepository.GetById(each.SupplierId)
			if filter.IsVerified && supplier.IsVerifiedSupplier {
				finalResult = append(finalResult, each)
			}
		}
	} else {
		finalResult = res
	}

	return finalResult, total
}

func (p productService) Create(dto model.ProductDto) error {
	product := p.productRepository.GetBySupplierIdAndProductName(dto.SupplierId, dto.Name)
	if product.Id != "" {
		return errors.New("product already exist")
	}
	dto.Id = uuid.New().String()
	dto.StatusId = 1
	brand := p.brandRepository.GetById(dto.BrandId)
	if brand.Id == "" {
		return errors.New("brand is invalid")
	}
	category := p.categoryRepository.GetById(dto.CategoryId)
	if category.Id == "" {
		return errors.New("category is invalid")
	}
	supplier := p.supplierRepository.GetById(dto.SupplierId)
	if supplier.Id == "" {
		return errors.New("supplier is invalid")
	}
	stock := model.ProductStock{
		Id:            uuid.New().String(),
		ProductId:     dto.Id,
		StockQuantity: dto.StockQuantity,
		UpdatedAt:     time.Now(),
	}
	go p.productStockRepository.Update(stock)

	product = dto.ConvertToProduct()
	err := p.productRepository.Create(product)
	if err != nil {
		return err
	}
	return nil
}

func (p productService) Get() []model.Product {
	return p.productRepository.Get()
}

func (p productService) Update(dto model.Product) error {
	return p.productRepository.Update(dto)
}

func (p productService) Delete(id string) error {
	return p.productRepository.Delete(id)
}

func NewProductService(productRepository database.ProductRepository, repository database.SupplierRepository, stockRepository database.ProductStockRepository, brandRepository database.BrandRepository, categoriesRepository database.CategoriesRepository) service.ProductServiceInterface {
	return &productService{
		productRepository:      productRepository,
		supplierRepository:     repository,
		productStockRepository: stockRepository,
		brandRepository:        brandRepository,
		categoryRepository:     categoriesRepository,
	}
}
