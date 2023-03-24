package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := t.Execute(os.Stdout, Cursos{
		{"Test", 12},
		{"Teste", 15},
		{"Teste", 18},
	})
	if err != nil {
		panic(err)
	}

}
