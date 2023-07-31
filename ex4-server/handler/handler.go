package handler

import (
	"github.com/dwarvesf/go23/ex4/handler/model"
	"github.com/gin-gonic/gin"
)

type A interface {
	X()
	Y()
}

type B interface {
	Z()
}

type impl struct {
}

func Auth(c *gin.Context) {
	r := model.LoginRequest{}
	err := c.ShouldBind(&r)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
	})
}
