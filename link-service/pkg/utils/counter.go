package utils

type Counter struct {
	min       uint64
	max       uint64
	current   uint64
	increment uint64
}

func NewCounter(min, max, increment uint64) *Counter {
	return &Counter{min: min, max: max, current: min, increment: increment}
}

func (c *Counter) Increment() uint64 {
	temp := c.current + 1
	if temp >= c.max {
		c.Reset()
		return c.current
	}
	c.current = temp
	return c.current

}

func (c *Counter) Reset() {
	c.min = c.max + 1
	c.max += c.increment
	c.current = c.min
}
