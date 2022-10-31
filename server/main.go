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
		log.Fatal("erro ao pegar os dados da tabela", err)
	}
	aluno := model.Aluno{}
	for dados.Next() {
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
func (t *Service) Insert(args []string, reply *string) (err error) {
	conexao := conect.Conect()
	comando, err := conexao.Prepare("INSERT INTO aluno(nome,idade,matricula,curso) VALUES(?,?,?,?)")
	if err != nil {
		log.Println("erro com o SQL")
	}
	comando.Exec(args[0], args[1], args[2], args[3])

	return nil
}
func (t *Service) Delete(args string, reply *model.Alunos) (err error) {
	conexao := conect.Conect()
	dados, err := conexao.Query("SELECT * FROM aluno WHERE id =?", args)
	if err != nil {
		log.Println("error com o select")
	}
	for dados.Next() {
		aluno := model.Aluno{}
		var idade, matricula, id int
		var nome, curso string
		dados.Scan(&nome, &idade, &matricula, &curso, &id)
		aluno.Id = id
		aluno.Idade = idade
		aluno.Matricula = matricula
		aluno.Curso = curso
		aluno.Nome = nome
		reply.Alunos = append(reply.Alunos, aluno)
	}
	comando, err := conexao.Prepare("DELETE FROM aluno WHERE id = ?")
	if err != nil {
		log.Println("erro com o deletar")
	}
	comando.Exec(args)
	return nil
}
func (t *Service) Edit(args []string, reply *model.Alunos) (err error) {
	conexao := conect.Conect()
	dados, err := conexao.Query("SELECT * FROM aluno WHERE id=?", args[4])
	if err != nil {
		log.Println("erro com a coleta de dados inicial")
	}
	for dados.Next() {
		aluno := model.Aluno{}
		var idade, matricula, id int
		var nome, curso string
		dados.Scan(&nome, &idade, &matricula, &curso, &id)
		aluno.Idade = idade
		aluno.Id = id
		aluno.Matricula = matricula
		aluno.Nome = nome
		aluno.Curso = curso
		reply.Alunos = append(reply.Alunos, aluno)
	}
	comando, err := conexao.Prepare("UPDATE aluno SET nome=?,matricula=?,idade=?,curso=? WHERE id =?")
	if err != nil {
		log.Println("error ao editar os dados")
	}
	comando.Exec(args[0], args[1], args[2], args[3], args[4])
	dados, err = conexao.Query("SELECT * FROM aluno WHERE id=?", args[4])
	if err != nil {
		log.Println("erro com a coleta de dados final")
	}
	for dados.Next() {
		aluno := model.Aluno{}
		var idade, matricula, id int
		var nome, curso string
		dados.Scan(&nome, &idade, &matricula, &curso, &id)
		aluno.Idade = idade
		aluno.Id = id
		aluno.Matricula = matricula
		aluno.Nome = nome
		aluno.Curso = curso
		reply.Alunos = append(reply.Alunos, aluno)
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
	log.Printf("esperando requisições pela porta :3030")
	http.Serve(listener, nil)
}
