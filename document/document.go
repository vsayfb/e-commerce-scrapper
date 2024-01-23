package document

type Document struct {
	MainTag  string `yaml:"MainTag"`
	TitleTag string `yaml:"TitleTag"`
	PriceTag string `yaml:"PriceTag"`
	ImageTag string `yaml:"ImageTag"`
	URLTag   string `yaml:"URLTag"`
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
