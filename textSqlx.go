package main


import (
	"log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type content struct {
	Id string `db:"content_id"`
	txt string `db:"txt"`
}

func main() {
	//_, err := sqlx.Connect("mysql", "user=root passowrd=cerx123 dbname=ceex ")
	_, err := sqlx.Open("mysql", "root:cerx123@tcp(127.0.0.1:3306)/ceex")
	if err != nil {
		log.Fatalln(err)
	}
}