package main

import (
	"html/template"
	"log"
	"net/http"
	"net/rpc"

	"git.com/model"
)

var aluno = model.Aluno{}
var cliente *rpc.Client
var err error
var tmpl = template.Must(template.ParseGlob("../view/*"))

func Show(w http.ResponseWriter, r *http.Request) {
	alunos := model.Alunos{}
	err = cliente.Call("Service.MostrarTudo", aluno, &alunos)
	if err != nil {
		log.Fatal("erro com a chamada", err)
	}
	log.Println(alunos)
	tmpl.ExecuteTemplate(w, "Show", alunos.Alunos)
}
func iniciar_cliente() {
	host := "localhost:3030"

	cliente, err = rpc.DialHTTP("tcp", host)
	if err != nil {
		log.Fatal("erro ao fazer a conexão com a porta")
	}

}
func main() {
	iniciar_cliente()
	log.Println("inicializando requisições")
	http.HandleFunc("/show", Show)
	http.ListenAndServe(":9090", nil)
}
