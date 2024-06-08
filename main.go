package main

import (
	"net/http"
	"strings"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", GeraTemplate)

	http.ListenAndServe(":8080", mux)
}
func GeraTemplate(w http.ResponseWriter, r *http.Request) {
	templates := []string{
		"content.html",
		"header.html",
		"footer.html",
	}

	// Criando o template apontando para o meu template principal
	tmp := template.New("content.html")

	// Atribuindo uma função Go para ser executada no html, pode ser qualquer função criada
	tmp.Funcs(template.FuncMap{"ToUpper": strings.ToUpper})
	tmp = template.Must(tmp.ParseFiles(templates...))

	err := tmp.Execute(w, []Curso{
		{"Go", 40},
		{"Java", 45},
		{"Pyton", 50},
	})
	if err != nil {
		panic(err)
	}
}
