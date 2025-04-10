package cache

import "github.com/Pashgunt/Validator/internal/contract"

type Cache struct {
	data map[string]contract.ConstraintInterface
}

func NewCache() *Cache {
	return &Cache{data: make(map[string]contract.ConstraintInterface)}
}

func (c Cache) Get(key string) contract.ConstraintInterface {
	return c.data[key]
}

func (c Cache) Set(key string, value contract.ConstraintInterface) {
	c.data[key] = value
}

func (c Cache) Exist(key string) bool {
	_, isset := c.data[key]

	return isset
}

func (c Cache) Delete(key string) {
	delete(c.data, key)
}
