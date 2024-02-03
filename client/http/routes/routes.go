package routes

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"

	"github.com/vsayfb/e-commerce-scrapper/client/response"
	"github.com/vsayfb/e-commerce-scrapper/search"
)

func Homepage(w http.ResponseWriter, _ *http.Request) {
	fp := path.Join("client", "http", "static", "html", os.Args[2]+".html")

	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		fmt.Println(err)

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	keyword := queryParams.Get("keyword")

	p := queryParams.Get("page")

	if keyword == "" && p == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	keyword = strings.TrimSpace(keyword)

	re := regexp.MustCompile("[^a-zA-Z0-9]+")

	keyword = re.ReplaceAllString(keyword, "+")

	if keyword == "" && p == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	page, parsingErr := strconv.ParseUint(p, 10, 8)

	if parsingErr != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	s := search.New(keyword, uint8(page))

	products := s.SearchSync()

	resp := response.New(keyword, uint8(page), products)

	bytes, encodingErr := json.Marshal(resp)

	if encodingErr != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(bytes)
}
