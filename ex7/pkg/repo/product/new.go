package product

import (
	"github.com/dwarvesf/go23/ex7/pkg/model"
	"gorm.io/gorm"
)

type Repo interface {
	CreateProduct(p model.Product) (*model.Product, error)
	UpdateProduct(p model.Product) (*model.Product, error)
	GetProductByID(ID int) (*model.Product, error)
	GetProductBySKU(sku string) (*model.Product, error)
	GetProductList() ([]model.Product, error)
	DeleteProductByID(ID int) error
}

// NewProductRepo returns a new instance of a ProductRepo.
func NewProductRepo(db *gorm.DB) Repo {
	return &pgRepo{
		DB: db,
	}
}

type pgRepo struct {
	DB *gorm.DB
}

func (r pgRepo) CreateProduct(p model.Product) (*model.Product, error) {
	return &p, r.DB.Create(&p).Error
}

func (r pgRepo) UpdateProduct(p model.Product) (*model.Product, error) {
	return &p, r.DB.Save(&p).Error
}

func (r pgRepo) GetProductByID(ID int) (*model.Product, error) {
	p := model.Product{}
	return &p, r.DB.Table("products").Where("id = ?", ID).First(&p).Error
}

func (r pgRepo) GetProductBySKU(sku string) (*model.Product, error) {
	p := model.Product{}
	return &p, r.DB.Table("products").Where("sku = ?", sku).First(&p).Error
}

func (r pgRepo) GetProductList() ([]model.Product, error) {
	rs := []model.Product{}
	return rs, r.DB.Table("products").Find(&rs).Error
}

func (r pgRepo) DeleteProductByID(ID int) error {
	return nil
}
