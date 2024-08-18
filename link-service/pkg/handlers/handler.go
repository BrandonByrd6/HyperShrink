package handlers

import "github.com/brandonbyrd6/link-service/pkg/repo"

type Handler struct {
	r repo.Repository
}

func NewHandler(r repo.Repository) *Handler {
	return &Handler{r}
}
