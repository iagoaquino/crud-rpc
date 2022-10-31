package main

import (
	"html/template"
	"log"
	"net/http"
	"net/rpc"

	"git.com/model"
)

var cliente *rpc.Client
var err error
var tmpl = template.Must(template.ParseGlob("../view/*"))

func Show(w http.ResponseWriter, r *http.Request) {
	alunos := model.Alunos{}
	aluno := model.Aluno{}
	err = cliente.Call("Service.MostrarTudo", aluno, &alunos)
	if err != nil {
		log.Fatal("erro com a chamada", err)
	}
	tmpl.ExecuteTemplate(w, "Show", alunos.Alunos)
}
func Delete(w http.ResponseWriter, r *http.Request) {
	alunos := model.Alunos{}
	var id = r.URL.Query().Get("id")

	err = cliente.Call("Service.Delete", id, &alunos)
	if err != nil {
		log.Fatal("erro com a chamada", err)
	}
	tmpl.ExecuteTemplate(w, "delete", alunos.Alunos)
}
func Edit(w http.ResponseWriter, r *http.Request) {
	alunos := model.Alunos{}
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		idade := r.FormValue("idade")
		matricula := r.FormValue("matricula")
		curso := r.FormValue("curso")
		id := r.FormValue("id")
		var dados []string
		dados = append(dados, nome)
		dados = append(dados, matricula)
		dados = append(dados, idade)
		dados = append(dados, curso)
		dados = append(dados, id)
		cliente.Call("Service.Edit", dados, &alunos)
	}
	tmpl.ExecuteTemplate(w, "edit", alunos.Alunos)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nomeI")
		idade := r.FormValue("idadeI")
		matricula := r.FormValue("matriculaI")
		curso := r.FormValue("cursoI")
		var dados []string
		dados = append(dados, nome)
		dados = append(dados, idade)
		dados = append(dados, matricula)
		dados = append(dados, curso)
		err = cliente.Call("Service.Insert", dados, nil)
		if err != nil {
			log.Fatal("erro com a chamada", err)
		}
		http.Redirect(w, r, "/show", 301)
	}
}
func iniciar_cliente() {
	host := "localhost:3030"

	cliente, err = rpc.DialHTTP("tcp", host)
	if err != nil {
		log.Fatal("erro ao fazer a conexão com a porta")
	}
}
func Greet(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "greet", "olá")
}
func main() {
	iniciar_cliente()
	log.Println("inicializando requisições")
	http.HandleFunc("/show", Show)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/delete", Delete)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/", Greet)
	http.ListenAndServe(":9090", nil)
}
