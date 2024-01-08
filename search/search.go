package search

import (
	"fmt"
	"github.com/vsayfb/e-commerce-scrapper/document"
	"github.com/vsayfb/e-commerce-scrapper/fetcher"
	"github.com/vsayfb/e-commerce-scrapper/product"
	"github.com/vsayfb/e-commerce-scrapper/resource"
	"sort"
	"sync"
)

type Searcher interface {
	SearchSync() []product.Product
	SearchAsync(ch chan<- product.Product, group *sync.WaitGroup)
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

	url := fmt.Sprintf("", s.keyword, s.page)

	r := resource.New(url, "Ebay", s.keyword, s.page)

	d := document.New(
		"",
		"",
		"",
		"",
		"",
	)

	r.AppendDoc(*d)

	f := fetcher.New(*r)

	fetchers = append(fetchers, f)

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

func (s Search) SearchAsync(ch chan<- product.Product, group *sync.WaitGroup) {
	//TODO implement me
	panic("implement me")
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
