package repo

import (
	"github.com/dwarvesf/go23/ex7/pkg/repo/product"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type pgRepo struct {
	db          *gorm.DB
	productRepo product.Repo
}

// NewPGRepo returns a new instance of a PGRepo.
func NewPGRepo(connectionStr string) Repo {
	db, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return &pgRepo{
		db:          db,
		productRepo: product.NewProductRepo(db),
	}
}

func (r pgRepo) AutoMigrate(models ...interface{}) error {
	for idx := range models {
		if err := r.db.AutoMigrate(models[idx]); err != nil {
			return err
		}
	}
	return nil
}

func (r pgRepo) Product() product.Repo {
	return r.productRepo
}

func (r pgRepo) DB() *gorm.DB {
	return r.db
}
