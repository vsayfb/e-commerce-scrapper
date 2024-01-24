package routes

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"

	"github.com/vsayfb/e-commerce-scrapper/search"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
	file := "client/http/html/" + os.Args[2] + ".html"

	tmpl, err := template.ParseFiles(file)
	if err != nil {
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
	err := r.ParseForm()
	if err != nil {
		return
	}

	keyword := r.Form.Get("keyword")

	s := search.New(keyword, 0)

	products := s.SearchSync()

	bytes, encodingErr := json.Marshal(products)

	if encodingErr != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(bytes)
}
