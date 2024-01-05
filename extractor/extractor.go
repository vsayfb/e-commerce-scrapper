package extractor

import "github.com/PuerkitoBio/goquery"

type Doc *goquery.Document
type Selection *goquery.Selection

type Extractor interface {
	Main(d Doc) Selection
	Title(s Selection) string
	Price(s Selection) string
	NumPrice(s Selection) float64
	URL(s Selection) string
	Image(s Selection) string
}
