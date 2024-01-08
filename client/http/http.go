package http

import (
	"fmt"
	"github.com/vsayfb/e-commerce-scrapper/product"
	"github.com/vsayfb/e-commerce-scrapper/search"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

func homePage(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.New("searchForm").Parse(`
			<!DOCTYPE html>
			<html>
			<head>
				<title>Search Form</title>
			</head>
			<body>
				<h1>Search Form</h1>
				<form method="post" action="/search">
					<label for="keyword">Enter keyword:</label>
					<input type="text" id="keyword" name="keyword" required>
					<input type="submit" value="Search">
				</form>
			</body>
			</html>
		`)

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

func searchHandler(w http.ResponseWriter, r *http.Request) {
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

	var products []product.Product

	if os.Args[2] == "sync" {
		products = s.SearchSync()
	}

	tmpl, err := template.New("searchResults").Parse(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Search Results</title>
		</head>
		<body>
			<h1>Search Results</h1>
			<ul>
				{{range .}}
					<li>
						<h2>{{.Title}}</h2>
						<p><strong>Site:</strong> {{.Site}}</p>
						<p><strong>Price:</strong> {{.Price}}</p>
						<p><strong>URL:</strong> <a href="{{.URL}}" target="_blank">{{.URL}}</a></p>
						<img src="{{.Image}}" alt="{{.Title}} Image">
					</li>
				{{end}}
			</ul>
		</body>
		</html>
	`)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func New(port uint16) {

	http.HandleFunc("/", homePage)
	http.HandleFunc("/search", searchHandler)

	p := ":" + strconv.Itoa(int(port))

	fmt.Println("Server is listening at", p)

	err := http.ListenAndServe(p, nil)

	if err != nil {
		return
	}
}
