package repo

import (
	"github.com/dwarvesf/go23/ex7/pkg/repo/product"
	"gorm.io/gorm"
)

// Repo is the interface that wraps the basic methods for a repository.
type Repo interface {
	DB() *gorm.DB
	AutoMigrate(models ...interface{}) error
	Product() product.Repo
}
