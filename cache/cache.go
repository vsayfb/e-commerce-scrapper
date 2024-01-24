package cache

import "github.com/vsayfb/e-commerce-scrapper/product"

type CacheClient interface {
	Add(key string, data interface{}) bool
	GetProducts(key string) ([]product.Product, error)
}

type Cache struct {
	client CacheClient
}

func New(client CacheClient) *Cache {
	c := &Cache{
		client: client,
	}

	return c
}

func (c *Cache) Add(key string, data interface{}) bool {
	return c.client.Add(key, data)
}

func (c *Cache) GetProducts(key string) ([]product.Product, error) {
	return c.client.GetProducts(key)
}
