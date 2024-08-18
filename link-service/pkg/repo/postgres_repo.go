package repo

import (
	"errors"
	"time"

	"github.com/brandonbyrd6/link-service/pkg/models"
	"github.com/brandonbyrd6/link-service/pkg/utils"
	"gorm.io/gorm"
)

type PostgresRepository struct {
	db        *gorm.DB
	shortener *utils.Shortener
}

func NewPostgresRepository(db *gorm.DB, s *utils.Shortener) PostgresRepository {
	return PostgresRepository{
		db:        db,
		shortener: s,
	}
}

func (p PostgresRepository) GetByShortUrl(ShortURL string) (*models.Url, error) {
	url := models.Url{}

	if !ok {
		return nil, errors.New("No Url Found")
	}
	return url, nil
}

func (p PostgresRepository) CreateUrl(LongURL string, UserID string) (*models.Url, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	surl := r.shortener.Generate()
	Url := models.Url{
		ID:        uint(r.shortener.Counter.GetCurrent()),
		UserId:    UserID,
		LongUrl:   LongURL,
		ShortUrl:  surl,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ExpiresAt: time.Now().Add(time.Duration(time.Duration.Hours(60))),
	}
	r.Urls[surl] = &Url
	return &Url, nil
}

func (p PostgresRepository) DeleteUrlByShortURL(ShortURL string) error {
	r.lock.RLock()
	defer r.lock.RUnlock()

	_, ok := r.Urls[ShortURL]
	if !ok {
		return errors.New("No Url Found")
	}
	delete(r.Urls, ShortURL)

	return nil
}
