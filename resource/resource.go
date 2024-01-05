package resource

import "github.com/vsayfb/e-commerce-scrapper/document"

type Resource struct {
	Site    string
	URL     string
	Keyword string
	Page    uint8
	Docs    []document.Document
}

func New(url, site, keyword string, page uint8) *Resource {

	r := &Resource{
		URL:     url,
		Site:    site,
		Keyword: keyword,
		Page:    page,
	}

	return r
}

func (r *Resource) AppendDoc(doc document.Document) {
	r.Docs = append(r.Docs, doc)
}
