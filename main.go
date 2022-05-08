package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	DbConnection, _ := sql.Open("mysql", "go_test:password@tcp(db:3306)/go_database")
	defer DbConnection.Close()
	cmd := `CREATE TABLE IF NOT EXISTS person(
				name varchar(14),
				age int)`
	_, err := DbConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("test")
}
