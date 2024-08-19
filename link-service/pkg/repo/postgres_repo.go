package repo

import (
	"time"

	"github.com/brandonbyrd6/link-service/pkg/models"
	"github.com/brandonbyrd6/link-service/pkg/utils"
	"gorm.io/gorm"
)

type PostgresRepository struct {
	db        *gorm.DB
	shortener *utils.Shortener
}

func NewPostgresRepository(db *gorm.DB, s *utils.Shortener) *PostgresRepository {
	return &PostgresRepository{
		db:        db,
		shortener: s,
	}
}

func (p *PostgresRepository) GetByShortUrl(ShortURL string) (*models.Url, error) {
	url := models.Url{}
	res := p.db.Where("short_url = ?", ShortURL).First(&url)
	if res.Error != nil {
		return nil, res.Error
	}
	return &url, nil
}

func (p *PostgresRepository) CreateUrl(LongURL string, UserID string) (*models.Url, error) {
	surl := p.shortener.Generate()
	url := models.Url{
		UserId:    UserID,
		LongUrl:   LongURL,
		ShortUrl:  surl,
		ExpiresAt: time.Now().Add(time.Duration(time.Duration.Hours(60))),
	}

	if res := p.db.Create(&url); res.Error != nil {
		return nil, res.Error
	}

	return &url, nil
}

func (p *PostgresRepository) DeleteUrlByShortURL(ShortURL string) error {
	url := models.Url{}

	res := p.db.Where("short_url = ?", ShortURL).Delete(&url)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
