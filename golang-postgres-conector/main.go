package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	//Data source information.
	ds := "user=postgres dbname=postgres password=postgres sslmode=disable"
	//openning a database conection
	conn, err := sql.Open("postgres", ds)
	if err == nil {
		fmt.Println(conn.Stats())
	}

	defer conn.Close()
}
