package product

import (
	"gorm.io/gorm"
)

type GormProductRepository struct {
	DB *gorm.DB
}

func NewGormProductRepository(db *gorm.DB) Repository {
	return &GormProductRepository{DB: db}
}

func (r *GormProductRepository) GetAll() ([]Product, error) {
	var products []Product
	result := r.DB.Find(&products)
	return products, result.Error
}

func (r *GormProductRepository) GetByID(id uint) (Product, error) {
	var product Product
	result := r.DB.First(&product, id)
	return product, result.Error
}

func (r *GormProductRepository) Create(product Product) (Product, error) {
	result := r.DB.Create(&product)
	return product, result.Error
}

func (r *GormProductRepository) Update(product Product) (Product, error) {
	result := r.DB.Save(&product)
	return product, result.Error
}

func (r *GormProductRepository) Delete(id uint) error {
	result := r.DB.Delete(&Product{}, id)
	return result.Error
}
