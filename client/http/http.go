package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vsayfb/e-commerce-scrapper/client/http/routes"
	"github.com/vsayfb/e-commerce-scrapper/client/http/socket"
	"net/http"
	"os"
	"strconv"
)

func New(port uint16) {

	r := mux.NewRouter()

	http.FileServer(http.Dir("./html"))

	r.HandleFunc("/", routes.Homepage)

	if os.Args[2] == "sync" {
		r.HandleFunc("/search", routes.SearchHandler).Queries("keyword", "{keyword:[a-zA-Z0-9]+}")
	} else {
		r.HandleFunc("/search", socket.ConcurrentSearchHandler)
	}

	p := ":" + strconv.Itoa(int(port))

	fmt.Println("Server is listening at", p)

	err := http.ListenAndServe(p, r)

	if err != nil {
		panic(err.Error())
	}
}
