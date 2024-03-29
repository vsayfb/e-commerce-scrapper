package fetcher

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/vsayfb/e-commerce-scrapper/extractor"
	"github.com/vsayfb/e-commerce-scrapper/product"
	"github.com/vsayfb/e-commerce-scrapper/resource"
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

func (f Fetch) FetchAsync(ch chan<- product.Product, wg *sync.WaitGroup) {
	defer func() {
		fmt.Println(f.resource.Site + "done.")

		wg.Done()
	}()

	resp, err := http.Get(f.resource.URL)
	if err != nil {
		fmt.Println("Product fetching error", err)
	}

	for _, doc := range f.resource.Docs {
		e := extractor.New(doc, resp.Body)

		e.MainSelection.Each(func(_ int, s *goquery.Selection) {
			p := product.Product{Site: f.resource.Site}

			p.Title = e.Title(s)
			p.Price = e.Price(s)
			p.URL = e.URL(s)
			p.Image = e.Image(s)
			p.NumPrice, err = e.NumPrice(s)

			if err == nil {
				ch <- p
			}
		})
	}
}

func (f Fetch) FetchSync() []product.Product {
	products := make([]product.Product, 0)

	resp, err := http.Get(f.resource.URL)
	if err != nil {
		fmt.Println("Product fetching error", err)
	}

	for _, doc := range f.resource.Docs {
		e := extractor.New(doc, resp.Body)

		e.MainSelection.Each(func(_ int, s *goquery.Selection) {
			p := product.Product{Site: f.resource.Site}

			p.Title = e.Title(s)
			p.Price = e.Price(s)
			p.URL = e.URL(s)
			p.Image = e.Image(s)
			p.NumPrice, err = e.NumPrice(s)

			if err == nil {
				products = append(products, p)
			}
		})
	}

	return products
}
