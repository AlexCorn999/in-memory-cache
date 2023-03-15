package cache

type Cache struct {
	Storage map[string]any
}

func New() Cache {
	return Cache{Storage: make(map[string]any)}
}

func (c *Cache) Set(key string, value any) {
	c.Storage[key] = value
}

func (c *Cache) Get(key string) any {
	return c.Storage[key]
}

func (c *Cache) Delete(key string) {
	delete(c.Storage, key)
}
