package inMemoryCache

import (
	"errors"
	"fmt"
	"time"
)

type Value struct {
	Value any
	Ttl   *time.Duration
}

type Cache struct {
	Storage map[string]*Value
}

func New() *Cache {
	return &Cache{make(map[string]*Value)}
}

func (c *Cache) Set(key string, value any, ttl time.Duration) {
	v := Value{
		Value: value,
		Ttl:   &ttl,
	}

	c.Storage[key] = &v
}

func (c *Cache) Get(key string) (any, error) {

	value, ok := c.Storage[key]
	if !ok {
		return nil, errors.New("value not found")
	}

	fmt.Println(*value.Ttl)

	return value.Value, nil
}

func (c *Cache) Delete(key string) error {
	_, ok := c.Storage[key]
	if !ok {
		return errors.New("value not found")
	}

	delete(c.Storage, key)
	return nil
}
