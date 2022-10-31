package conect

import (
	"database/sql" // Pacote Database SQL para realizar Query
	"log"

	_ "github.com/go-sql-driver/mysql" // Driver Mysql para Go
)

func Conect() (bd *sql.DB) {
	driver := "mysql"
	name := "alunos"

	bd, err := sql.Open(driver, "root:teste123@tcp(172.17.0.2:3306)/"+name)

	if err != nil {
		log.Fatal("error com a conexão do banco", err)
	} else {
		log.Print("conectou ao banco")
	}
	return bd
}
