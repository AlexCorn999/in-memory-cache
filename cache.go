package inMemoryCache

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Value struct {
	Value             any
	Created           time.Time
	DefaultExpiration int64
}

type Cache struct {
	sync.RWMutex
	cleanupInterval time.Duration
	storage         map[string]*Value
}

func New() *Cache {

	db := &Cache{
		storage:         make(map[string]*Value),
		cleanupInterval: time.Second * 2,
	}

	db.StartGC()

	return db
}

// установка значений по ключу
func (c *Cache) Set(key string, value any, ttl time.Duration) {

	c.Lock()

	defer c.Unlock()

	v := Value{
		Value:             value,
		Created:           time.Now(),
		DefaultExpiration: time.Now().Add(ttl).UnixNano(),
	}

	c.storage[key] = &v
}

// получение значения по ключу
func (db *Cache) Get(key string) (any, error) {

	db.RLock()

	defer db.RUnlock()

	value, ok := db.storage[key]
	if !ok {
		return nil, errors.New("value not found")
	}

	return value.Value, nil
}

// удаление ключа
func (db *Cache) Delete(key string) error {

	db.Lock()

	defer db.Unlock()

	if _, ok := db.storage[key]; !ok {
		return errors.New("value not found")
	}

	delete(db.storage, key)
	return nil
}

// начало очистки для всех хранилищ
func (db *Cache) StartGC() {
	go db.GC()
}

func (db *Cache) GC() {

	for {
		// ожидаем время установленное в cleanupInterval
		<-time.After(db.cleanupInterval)

		if db.storage == nil {
			return
		}

		// Ищем элементы с истекшим временем жизни и удаляем из хранилища
		if keys := db.expiredKeys(); len(keys) != 0 {
			db.clearItems(keys)

		}

	}

}

// expiredKeys возвращает список "просроченных" ключей
func (db *Cache) expiredKeys() (keys []string) {

	db.RLock()

	defer db.RUnlock()

	for k, i := range db.storage {
		if time.Now().UnixNano() > i.DefaultExpiration && i.DefaultExpiration > 0 {
			fmt.Println("xxxxxxxxx")
			keys = append(keys, k)
		}
	}

	return
}

// clearItems удаляет ключи из переданного списка, в нашем случае "просроченные"
func (db *Cache) clearItems(keys []string) {

	db.Lock()

	defer db.Unlock()

	for _, k := range keys {
		delete(db.storage, k)
	}
}
