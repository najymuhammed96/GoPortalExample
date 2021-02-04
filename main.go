package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func main() {
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", do)
	http.ListenAndServe(":8080", nil)
}

func do(w http.ResponseWriter, r *http.Request) {
	tmpl := parseTemplates()
	datama := map[string]interface{}{"head": "go portal",
		"foot": "Develoed by Najy", "text": "Go Portal"}
	err := tmpl.ExecuteTemplate(w, "index.html", datama)
	if err != nil {
		println(err.Error())
	}
}

func parseTemplates() *template.Template {
	templ := template.New("")
	err := filepath.Walk("templates", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			_, err = templ.ParseFiles(path)
			if err != nil {
				log.Println(err)
			}
		}

		return err
	})

	if err != nil {
		panic(err)
	}
	return templ
}
