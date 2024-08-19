package utils

type CounterInterface interface {
	Increment() uint64
	GetCurrent() uint64
	Reset()
}

type ShortenerInterface interface {
	Generate() string
}
