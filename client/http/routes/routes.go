package routes

import (
	"encoding/json"
	"fmt"
	"github.com/vsayfb/e-commerce-scrapper/search"
	"html/template"
	"net/http"
	"os"
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
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err := r.ParseForm()

	if err != nil {
		return
	}

	keyword := r.Form.Get("keyword")

	s := search.New(keyword, 0)

	products := s.SearchSync()

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(products)

	if err != nil {

		http.Error(w, fmt.Sprintf("Error encoding JSON: %v", err), http.StatusInternalServerError)

		return
	}
}
