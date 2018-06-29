package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type content struct {
	Id   int            `db:"content_id"`
	Txt  string         `db:"txt"`
	Txt1 sql.NullString `db:"txt1"`
	Txt2 sql.NullString `db:"txt2"`
	Txt3 sql.NullString `db:"txt3"`
}

func main() {
	//_, err := sqlx.Connect("mysql", "user=root passowrd=cerx123 dbname=ceex ")
	db, err := sqlx.Open("mysql", "root:cerx123@tcp(127.0.0.1:3306)/ceex")
	if err != nil {
		log.Fatalln(err)
	}

	txts := []content{}
	err = db.Select(&txts, "SELECT * FROM jc_content_txt")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	//one := txts[0]
	fmt.Println(len(txts))
}
