package client

import "github.com/vsayfb/e-commerce-scrapper/client/cli"
import "github.com/vsayfb/e-commerce-scrapper/client/http"

func New(client string) {

	if client == "cli" {
		cli.New()
	}

	http.New(5555)
}
