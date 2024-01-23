package source

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/vsayfb/e-commerce-scrapper/document"
	"gopkg.in/yaml.v2"
)

type Source struct {
	Website struct {
		Name      string              `yaml:"Name"`
		SourceURL string              `yaml:"SourceURL"`
		Docs      []document.Document `yaml:"Docs"`
	} `yaml:"Website"`
}

func GetSource() []Source {
	_, filename, _, _ := runtime.Caller(0)

	dir := filepath.Dir(filename)

	filePath := filepath.Join(dir, "source.yaml")

	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	var sources []Source

	err = yaml.Unmarshal(yamlFile, &sources)

	if err != nil {
		log.Fatal(err)
	}

	if len(sources) == 0 {
		log.Fatal("Empty source file.")
	}

	return make([]Source, 0)
}
