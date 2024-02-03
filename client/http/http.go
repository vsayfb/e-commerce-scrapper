package http

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vsayfb/e-commerce-scrapper/client/http/routes"
	"github.com/vsayfb/e-commerce-scrapper/client/http/socket"
)

func New(port uint16) {
	r := mux.NewRouter()

	fp := path.Join("client", "http", "static")

	fs := http.FileServer(http.Dir(fp))

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	r.HandleFunc("/", routes.Homepage)

	if os.Args[2] == "sync" {
		r.HandleFunc("/search", routes.SearchHandler)
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
