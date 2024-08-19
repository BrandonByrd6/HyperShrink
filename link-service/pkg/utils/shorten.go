package utils

import (
	"encoding/binary"

	"github.com/btcsuite/btcutil/base58"
)

type Shortener struct {
	Counter CounterInterface
}

func NewShortener(c CounterInterface) *Shortener {
	return &Shortener{Counter: c}
}

func (s *Shortener) Generate() string {
	i := s.Counter.Increment()
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, i)
	str := base58.Encode(b)
	return str[:10]
}
