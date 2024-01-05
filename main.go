package main

import (
	"fmt"
	"github.com/vsayfb/e-commerce-scrapper/document"
	"github.com/vsayfb/e-commerce-scrapper/resource"
)

func main() {
	doc := document.New("", "", "", "", "")

	r := resource.New("", "")

	r.AppendDoc(*doc)

	fmt.Println(r)
}
