package utils

import (
	"encoding/binary"

	"github.com/btcsuite/btcutil/base58"
)

type Shortener struct {
	counter *Counter
}

func NewShortener(c *Counter) *Shortener {
	return &Shortener{counter: c}
}

func (s *Shortener) Generate() string {
	i := s.counter.Increment()
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, i)
	str := base58.Encode(b)
	return str[:10]
}
