package product

type Repository interface {
	GetAll() ([]Product, error)
	GetByID(id uint) (Product, error)
	Create(product Product) (Product, error)
	Update(product Product) (Product, error)
	Delete(id uint) error
}
