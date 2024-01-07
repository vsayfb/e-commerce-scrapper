package extractor

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/vsayfb/e-commerce-scrapper/document"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type Doc goquery.Document

type Selection *goquery.Selection

type Extractor interface {
	Title(sel Selection) string
	Price(sel Selection) string
	NumPrice(sel Selection) (error, float64)
	URL(sel Selection) string
	Image(sel Selection) string
}

func New(doc document.Document, responseBody io.ReadCloser) *Extract {

	d, err := goquery.NewDocumentFromReader(responseBody)

	if err != nil {
		return &Extract{}
	}

	items := d.Selection.Find(doc.MainTag)

	e := Extract{MainSelection: items, doc: doc}

	return &e
}

type Extract struct {
	MainSelection *goquery.Selection
	doc           document.Document
}

func (e Extract) Title(sel Selection) string {
	title := ""

	(*sel).Find(e.doc.TitleTag).Each(func(i int, s *goquery.Selection) {
		title += s.Text()
	})

	return title

}

func (e Extract) Price(sel Selection) string {

	price := ""

	(*sel).Find(e.doc.PriceTag).Each(func(i int, s *goquery.Selection) {
		price += s.Text() + " "
	})

	return price

}

func (e Extract) NumPrice(sel Selection) (float64, error) {
	priceStr := e.Price(sel)

	split := strings.Split(priceStr, " ")

	re := regexp.MustCompile("[^0-9,.]")

	priceStr = split[0]

	priceStr = re.ReplaceAllString(priceStr, "")

	m := make(map[rune]uint8)

	separator := ' '

	for _, c := range priceStr {
		if c == ',' || c == '.' {

			if separator == ' ' {
				separator = c
			}

			m[c]++
		}
	}

	// format is "1.000.000" or "1,000,000" or "20.00"
	if len(m) == 1 {

		if m[separator] == 1 {
			// format is "20.00" or "20,00"

			priceStr = strings.ReplaceAll(priceStr, ",", ".")
		} else {
			// format is "1.000.000" or "1,000,000"

			priceStr = strings.ReplaceAll(priceStr, string(separator), "")
		}

	} else if len(m) > 1 {

		// format is "1,540.30" or "1.540,30" or "1,234,567,890.12" or "1.234.567.890,12"

		priceStr = strings.ReplaceAll(priceStr, string(separator), "")

		priceStr = strings.Replace(priceStr, ",", ".", 1)
	}

	v, err := strconv.ParseFloat(priceStr, 64)

	if err != nil {
		return -1, err
	}

	return v, nil
}

func (e Extract) URL(sel Selection) string {

	url, _ := (*sel).Find(e.doc.URLTag).Attr("href")

	return url
}

func (e Extract) Image(sel Selection) string {

	img, _ := (*sel).Find(e.doc.ImageTag).Attr("src")

	return img

}
