package handler

import (
	"github.com/dwarvesf/go23/ex7/pkg/config"
	"github.com/dwarvesf/go23/ex7/pkg/repo"
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
