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

	//Throwing erro
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	rows, err := conn.Query("select * from person")

	if err != nil {
		panic(err)
	}

	people := []Person{}
	for rows.Next() {
		person := Person{}
		rows.Scan(&person.Id, &person.First_name, &person.Middle_name, &person.Last_name, &person.Contacts)
		people = append(people, person)
	}

	fmt.Println(people)
}

type Person struct {
	Id          int
	First_name  string
	Middle_name string
	Last_name   string
	Contacts    string
}
