package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"git.com/conect"
	"git.com/model"
)

type Service string

func (t *Service) MostrarTudo(args *model.Aluno, tabela *model.Alunos) (err error) {
	conexao := conect.Conect()

	dados, err := conexao.Query("SELECT * FROM aluno")
	if err != nil {
		log.Println("erro ao pegar os dados da tabela")
	}
	aluno := model.Aluno{}
	for dados.Next() {
		log.Println("entrei aqui")
		var idade, id, matricula int
		var nome, curso string
		dados.Scan(&nome, &idade, &matricula, &curso, &id)
		aluno.Id = id
		aluno.Idade = idade
		aluno.Matricula = matricula
		aluno.Nome = nome
		aluno.Curso = curso
		tabela.Alunos = append(tabela.Alunos, aluno)
	}
	return nil
}
func main() {
	servico := new(Service)
	rpc.Register(servico)
	rpc.HandleHTTP()

	port := ":3030"
	listener, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatal("error ao escutar a porta")
	}
	log.Printf("eu to te ouvindo arrombado")
	http.Serve(listener, nil)
}
