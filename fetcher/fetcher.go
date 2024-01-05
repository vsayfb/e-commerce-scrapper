package fetcher

import (
	"github.com/vsayfb/e-commerce-scrapper/product"
	"github.com/vsayfb/e-commerce-scrapper/resource"
	"sync"
)

type Fetcher interface {
	FetchSync() []product.Product
	FetchAsync(ch chan<- product.Product, group *sync.WaitGroup)
}

type Fetch struct {
	resource resource.Resource
}

func New(r resource.Resource) *Fetch {
	f := Fetch{resource: r}

	return &f
}

func (f Fetch) FetchSync() []product.Product {
	return nil
}

func (f Fetch) FetchAsync(ch chan<- product.Product, group *sync.WaitGroup) {

}
