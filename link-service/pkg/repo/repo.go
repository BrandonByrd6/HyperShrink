package repo

import "github.com/brandonbyrd6/link-service/pkg/models"

type Repository interface {
	GetByShortUrl(ShortURL string) (*models.Url, error)
	CreateUrl(LongURL string, UserID string) (*models.Url, error)
	DeleteUrlByShortURL(ShortURL string) error
}
