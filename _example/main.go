package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/uoula/go-nulltype"
)

func main() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var nt nulltype.NullTime
	err = db.QueryRow("select current_timestamp").Scan(&nt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(nt)
}
