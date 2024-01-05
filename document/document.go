package document

type Document struct {
	MainTag  string
	TitleTag string
	PriceTag string
	ImageTag string
	URLTag   string
}

func New(main, title, price, image, url string) *Document {

	doc := Document{
		MainTag:  main,
		TitleTag: title,
		PriceTag: price,
		ImageTag: image,
		URLTag:   url,
	}

	return &doc

}
