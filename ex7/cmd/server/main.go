package main

import (
	"fmt"
	"net/http"

	"github.com/dwarvesf/go23/ex7/pkg/config"
	"github.com/dwarvesf/go23/ex7/pkg/handler"
	"github.com/dwarvesf/go23/ex7/pkg/model"
	"github.com/dwarvesf/go23/ex7/pkg/repo"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	r := repo.NewPGRepo(cfg.DBUrl)

	r.AutoMigrate(&model.Product{}, &model.Vendor{})

	fmt.Println("Starting server...")
	srv := initServer(cfg, r)
	srv.Run(":" + cfg.Port)
	fmt.Println("Ending server...")
}

func initServer(cfg config.Config, r repo.Repo) *gin.Engine {
	srv := gin.Default()
	h := handler.New(cfg, r)

	srv.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	apiGroup := srv.Group("/api/v1")
	apiGroup.POST("/products", h.CreateProduct)
	apiGroup.GET("/products", h.ListProducts)

	return srv
}
