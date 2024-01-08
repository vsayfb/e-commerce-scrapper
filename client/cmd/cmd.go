package cmd

import (
	"bufio"
	"fmt"
	"github.com/vsayfb/e-commerce-scrapper/search"
	"os"
)

func New() {

	scanner := bufio.NewScanner(os.Stdin)

	var page uint8 = 0

	for {

		fmt.Print("Search for a product: ")

		scanner.Scan()

		userInput := scanner.Text()

		s := search.New(userInput, page)

		if os.Args[2] == "sync" {
			products := s.SearchSync()

			for i, p := range products {
				fmt.Printf("%v - Site: %v - Price: %v - Title: %v \n", i, p.Site, p.Price, p.Title)
			}
		}

		page++
	}
}
