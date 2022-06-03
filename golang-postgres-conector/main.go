package main

import (
	"database/sql"
	"fmt"
	"golang-training-repository/golang-postgres-conector/model"

	_ "github.com/lib/pq"
)

func main() {
	people := []model.Person{}
	//Data source information.
	ds := "user=postgres dbname=postgres password=postgres sslmode=disable"
	//openning a database conection
	conn, err := sql.Open("postgres", ds)

	//Throwing erro
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	if rows, err := conn.Query("select * from person"); err != nil {
		panic(err)
	} else {

		for rows.Next() {
			person := model.Person{}
			rows.Scan(&person.Id, &person.First_name, &person.Middle_name, &person.Last_name, &person.Contacts)
			people = append(people, person)
		}

	}

	fmt.Println(people)
}
