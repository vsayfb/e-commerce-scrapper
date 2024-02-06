package search

import (
	"encoding/json"
	"fmt"
	"sort"
	"sync"

	"github.com/vsayfb/e-commerce-scrapper/cache"
	"github.com/vsayfb/e-commerce-scrapper/cache/redis"
	"github.com/vsayfb/e-commerce-scrapper/fetcher"
	"github.com/vsayfb/e-commerce-scrapper/product"
	"github.com/vsayfb/e-commerce-scrapper/resource"
	"github.com/vsayfb/e-commerce-scrapper/source"
)

type Searcher interface {
	SearchSync() []product.Product
	SearchAsync(ch chan<- product.Product)
}

func New(keyword string, page uint8) *Search {
	s := &Search{
		keyword: keyword, page: page,
	}

	return s
}

type Search struct {
	keyword string
	page    uint8
}

func (s Search) SearchSync() []product.Product {
	c := cache.New(redis.New())

	key := fmt.Sprintf("%v-%v", s.keyword, s.page)

	result, err := c.GetProducts(key)

	if err == nil {
		return result
	}

	fetchers := make([]fetcher.Fetch, 0)

	sources := source.GetSource()

	for _, source := range sources {

		url := fmt.Sprintf(source.Website.SourceURL, s.keyword, s.page)

		r := resource.New(url, source.Website.Name, s.keyword, s.page)

		for _, d := range source.Website.Docs {
			r.AppendDoc(d)
		}

		fetchers = append(fetchers, *fetcher.New(*r))
	}

	products := make([]product.Product, 0)

	var wg sync.WaitGroup

	wg.Add(len(fetchers))

	for _, f := range fetchers {
		go func(fetcher fetcher.Fetch) {
			defer wg.Done()

			res := fetcher.FetchSync()

			products = append(products, res...)
		}(f)
	}

	wg.Wait()

	sort.Slice(products, func(i, j int) bool {
		return products[i].NumPrice < products[j].NumPrice
	})

	go func() {
		bytes, err := json.Marshal(products)
		if err != nil {
			fmt.Print("Marshal error", err)
		} else {
			c.Add(key, bytes)
		}
	}()

	return products
}

func (s Search) SearchAsync(ch chan<- product.Product) {
	fetchers := make([]*fetcher.Fetch, 0)

	sources := source.GetSource()

	for _, source := range sources {

		url := fmt.Sprintf(source.Website.SourceURL, s.keyword, s.page)

		r := resource.New(url, source.Website.Name, s.keyword, s.page)

		for _, d := range source.Website.Docs {
			r.AppendDoc(d)
		}

		fetchers = append(fetchers, fetcher.New(*r))
	}

	var wg sync.WaitGroup

	go func() {
		for _, f := range fetchers {
			wg.Add(1)

			go f.FetchAsync(ch, &wg)
		}

		wg.Wait()

		close(ch)
	}()
}
