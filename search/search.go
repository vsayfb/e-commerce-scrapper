package search

import (
	"fmt"
	"sort"
	"sync"

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

	products := make([]product.Product, 0)

	for _, f := range fetchers {

		res := f.FetchSync()

		if len(products) == 0 {
			products = res
		} else {
			products = mergeSort(products, res)
		}
	}

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

func mergeSort(p1, p2 []product.Product) []product.Product {
	res := make([]product.Product, 0)

	sort.Slice(p1, func(i, j int) bool {
		return p1[i].NumPrice < p1[j].NumPrice
	})

	sort.Slice(p2, func(i, j int) bool {
		return p2[i].NumPrice < p2[j].NumPrice
	})

	i := 0
	j := 0

	for i < len(p1) && j < len(p2) {
		if p1[i].NumPrice <= p2[j].NumPrice {
			res = append(res, p1[i])
			i++
		} else {
			res = append(res, p2[j])
			j++
		}
	}

	res = append(res, p1[i:]...)
	res = append(res, p2[j:]...)

	return res
}
