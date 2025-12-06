package context

type Context struct {
	Numbers map[string]float64
}

func NewContext(numbers map[string]float64) *Context {
	return &Context{
		Numbers: numbers,
	}
}

func (c *Context) GetNumber(key string) (float64, bool) {
	v, ok := c.Numbers[key]
	return v, ok
}
