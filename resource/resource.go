package resource

import "github.com/vsayfb/e-commerce-scrapper/document"

type Resource struct {
	Site string
	URL  string
	docs []document.Document
}

func New(url, site string) *Resource {

	r := &Resource{
		URL:  url,
		Site: site,
	}

	return r
}

func (r *Resource) AppendDoc(doc document.Document) {
	r.docs = append(r.docs, doc)
}
