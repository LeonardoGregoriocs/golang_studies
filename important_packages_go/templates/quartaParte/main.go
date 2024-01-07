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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("template.html"))
		err := t.Execute(w, Cursos{
			{"Go", 40},
			{"Java", 30},
			{"Python", 20},
		})
		if err != nil {
			panic(err)
		}

	})

	http.HandleFunc("/go", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("go.html").ParseFiles("go.html"))
		err := t.Execute(w, nil)
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8082", nil)

}
