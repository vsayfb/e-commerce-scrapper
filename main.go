package main

import (
	"fmt"
	"github.com/vsayfb/e-commerce-scrapper/client"
	"os"
)

func main() {

	args := os.Args

	if len(args) < 2 || (args[1] != "cmd" && args[1] != "http") {
		fmt.Println("You have to choose an option, cmd or http.")

		os.Exit(1)
	}

	if len(args) < 3 && (args[1] != "sync" && args[1] != "async") {
		fmt.Println("You have to choose an option, sync or async.")

		os.Exit(1)
	}

	client.New(args[1])
}
