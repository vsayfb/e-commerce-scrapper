package cli

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/vsayfb/e-commerce-scrapper/client/http"
	"github.com/vsayfb/e-commerce-scrapper/product"
	"github.com/vsayfb/e-commerce-scrapper/search"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func New() {

	sync := os.Args[2] == "sync"

	if sync {
		openCLI(true, nil)
	}

	go http.New(5555)

	conn, _, connErr := websocket.DefaultDialer.Dial("ws://localhost:5555/search", nil)

	if connErr != nil {
		panic(connErr.Error())
	}

	defer func(conn *websocket.Conn) {

		err := conn.Close()

		if err != nil {
			log.Fatalln(err)
		}

	}(conn)

	go func() {

		for {

			_, message, readErr := conn.ReadMessage()

			if readErr != nil {
				log.Fatalln(readErr)
				return
			}

			var products []product.Product

			err := json.Unmarshal(message, &products)

			if err != nil {
				log.Fatalln(err)
				return
			}

			clearCLI()

			for i, p := range products {
				fmt.Printf("%v - Site: %v - Title: %v - Price: %v \n", i, p.Site, p.Title, p.Price)
			}

			fmt.Printf("\n Search for a product:")
		}

	}()

	openCLI(sync, conn)

}

func openCLI(sync bool, conn *websocket.Conn) {

	for {
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Printf("Search for a product: ")

		scanner.Scan()

		inp := scanner.Text()

		if inp == "q" {
			os.Exit(0)
		}

		if sync {
			s := search.New(inp, 0)

			products := s.SearchSync()

			for i, p := range products {
				fmt.Printf("%v - Site: %v - Price: %v - Title: %v \n", i, p.Site, p.Price, p.Title)
			}
		} else {

			writeErr := conn.WriteMessage(websocket.TextMessage, []byte(inp+"-"+"0"))

			if writeErr != nil {
				log.Fatalln(writeErr)
			}
		}
	}

}

func clearCLI() {
	var cmd *exec.Cmd

	switch runtime.GOOS {

	case "linux", "darwin":
		cmd = exec.Command("clear")

	case "windows":
		cmd = exec.Command("cli", "/c", "cls")

	default:
		fmt.Println("Unsupported OS")
		return
	}

	cmd.Stdout = os.Stdout

	err := cmd.Run()

	if err != nil {
		fmt.Println(err)

		return
	}
}
