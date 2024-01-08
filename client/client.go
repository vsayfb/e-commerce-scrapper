package client

import "github.com/vsayfb/e-commerce-scrapper/client/cmd"
import "github.com/vsayfb/e-commerce-scrapper/client/http"

func New(client string) {

	if client == "cmd" {
		cmd.New()
	}

	http.New(5555)
}
