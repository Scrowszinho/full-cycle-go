package main

import (
	"html/template"
	"net/http"
	"os"
	"strings"
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
	t := template.New("content.html")
	t.Funcs(template.FuncMap{"ToUpper": strings.ToUpper})
	t = template.Must(t.ParseFiles(templates...))
	err := t.Execute(os.Stdout, Cursos{
		{"Test", 12},
		{"Teste", 15},
		{"Teste", 18},
	})
	if err != nil {
		panic(err)
	}
	http.ListenAndServe(":8080", nil)
}
