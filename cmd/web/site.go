package main

import (
	"html/template"
	"log"
	"net/http"
)

type Site struct {
	Templates *template.Template
}

func (s *Site) RenderIndex(w http.ResponseWriter, r *http.Request) {
	err := s.Templates.ExecuteTemplate(w, "index.page.tpl", nil)
	if err != nil {
		log.Println("error executing template for index: ", err)
		http.Error(w, "error rendering template", 500)
		return
	}
}
