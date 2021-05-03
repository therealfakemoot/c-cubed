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
	err := s.Templates.Execute(w, nil)
	if err != nil {
		log.Println("error executing template for index: ", err)
	}
}
