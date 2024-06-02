package services

import "golang-hactiv8-lab/chapter6/challange/domain/product"

type ProductService struct {
	repo product.Repository
}

func NewProductService(repo product.Repository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAll() ([]product.Product, error) {
	return s.repo.GetAll()
}

func (s *ProductService) GetByID(id uint) (product.Product, error) {
	return s.repo.GetByID(id)
}

func (s *ProductService) Create(p product.Product) (product.Product, error) {
	return s.repo.Create(p)
}

func (s *ProductService) Update(p product.Product) (product.Product, error) {
	return s.repo.Update(p)
}

func (s *ProductService) Delete(id uint) error {
	return s.repo.Delete(id)
}
