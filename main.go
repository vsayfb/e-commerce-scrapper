package main

import (
	"fmt"
	"github.com/vsayfb/e-commerce-scrapper/document"
	"github.com/vsayfb/e-commerce-scrapper/fetcher"
	"github.com/vsayfb/e-commerce-scrapper/resource"
)

func main() {

	r := resource.New("", "", "", 222)

	doc := document.New("", "", "", "", "")

	r.AppendDoc(*doc)

	f := fetcher.New(*r)

	f.FetchSync()

	fmt.Println(r)
}
