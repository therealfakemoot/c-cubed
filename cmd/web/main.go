package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	var s Site
	tmpls, err := filepath.Glob("./templates/*.*.tpl")
	if err != nil {
		log.Fatalln("couldn't get list of template files: %s", err)
	}
	log.Println(tmpls)

	ts, err := template.ParseFiles(tmpls...)
	if err != nil {
		log.Fatalln("couldn't parse template files: %s", err)
	}
	s.Templates = ts

	http.HandleFunc("/", s.RenderIndex)
	log.Fatal(http.ListenAndServe(":4000", nil))
}
