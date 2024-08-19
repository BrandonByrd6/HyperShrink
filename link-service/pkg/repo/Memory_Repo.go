package repo

import (
	"errors"
	"sync"
	"time"

	"github.com/brandonbyrd6/link-service/pkg/models"
	"github.com/brandonbyrd6/link-service/pkg/utils"
)

type MemoryRepository struct {
	Urls      map[string]*models.Url
	lock      *sync.RWMutex
	shortener *utils.Shortener
}

func NewMemoryRepository(s *utils.Shortener) *MemoryRepository {
	return &MemoryRepository{
		Urls:      map[string]*models.Url{},
		lock:      &sync.RWMutex{},
		shortener: s,
	}
}

func (r *MemoryRepository) GetByShortUrl(ShortURL string) (*models.Url, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	url, ok := r.Urls[ShortURL]
	if !ok {
		return nil, errors.New("no url found")
	}
	return url, nil
}

func (r *MemoryRepository) CreateUrl(LongURL string, UserID string) (*models.Url, error) {
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

func (r *MemoryRepository) DeleteUrlByShortURL(ShortURL string) error {
	r.lock.RLock()
	defer r.lock.RUnlock()

	_, ok := r.Urls[ShortURL]
	if !ok {
		return errors.New("no url found")
	}
	delete(r.Urls, ShortURL)

	return nil
}
