package cache

type cache struct {
	mapa map[string]interface{}
}

func New() *cache {
	return &cache{
		mapa: make(map[string]interface{}),
	}
}

func (c *cache) Set(key string, value interface{}) {
	c.mapa[key] = value
}

func (c *cache) Get(key string) interface{} {
	if c.mapa[key] == nil {
		return "No record found"
	}
	return c.mapa[key]
}

func (c *cache) Delete(key string) {
	delete(c.mapa, key)
}
