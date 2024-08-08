package handlers

import (
	"github.com/brandonbyrd6/link-service/pkg/utils"
	"gorm.io/gorm"
)

type Handler struct {
	DB        *gorm.DB //TODO: Remove with Repositories, needs Services?
	shortener *utils.Shortener
}

func NewHandler(db *gorm.DB, shortener *utils.Shortener) *Handler {
	return &Handler{DB: db, shortener: shortener}
}
