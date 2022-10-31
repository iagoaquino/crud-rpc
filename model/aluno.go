package model

type Aluno struct {
	Nome      string
	Idade     int
	Matricula int
	Curso     string
	Id        int
}

type Alunos struct {
	Alunos []Aluno
}
