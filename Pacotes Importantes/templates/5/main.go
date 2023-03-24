package main

import (
	"net/http"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	templates := []string{
		"header.html",
		"template.html",
		"footer.html",
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles(templates...))
		err := t.Execute(w, Cursos{
			{"Test", 12},
			{"Teste", 15},
			{"Teste", 18},
		})
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8080", nil)
}
