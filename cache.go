package inMemoryCache

import "errors"

type Cache struct {
	Storage map[string]any
}

func New() *Cache {
	return &Cache{Storage: make(map[string]any)}
}

func (c *Cache) Set(key string, value any) {
	c.Storage[key] = value
}

func (c *Cache) Get(key string) (any, error) {
	value, ok := c.Storage[key]
	if !ok {
		return nil, errors.New("value not found")
	}

	return value, nil
}

func (c *Cache) Delete(key string) error {
	_, ok := c.Storage[key]
	if !ok {
		return errors.New("value not found")
	}

	delete(c.Storage, key)
	return nil
}
