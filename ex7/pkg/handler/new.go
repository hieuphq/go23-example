package handler

import (
	"github.com/dwarvesf/go23/ex7/pkg/config"
	"github.com/dwarvesf/go23/ex7/pkg/model"
	"github.com/dwarvesf/go23/ex7/pkg/repo"
	"github.com/gin-gonic/gin"
)

// Handler is
type Handler struct {
	cfg config.Config
	r   repo.Repo
}

func New(cfg config.Config, r repo.Repo) *Handler {
	return &Handler{
		cfg: cfg,
		r:   r,
	}
}

func (h Handler) CreateProduct(ctx *gin.Context) {
	req := model.CreateProductRequest{}
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	product := model.Product{
		Name:        req.Name,
		Description: req.Description,
		SKU:         req.SKU,
		Price:       req.Price,
		Currency:    req.Currency,
	}

	p, err := h.r.Product().CreateProduct(product)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, p)
}
func (h Handler) ListProducts(ctx *gin.Context) {
	products, err := h.r.Product().GetProductList()
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, products)
}
