package main

import (
	"database/sql"
	"fmt"
	"golang-training-repository/golang-postgres-conector/model"
	"log"

	_ "github.com/lib/pq"
)

var Conn *sql.DB

func main() {
	Conn = Connect()
	Insert(model.Person{
		Id:          0,
		First_name:  "Cath",
		Middle_name: "Little",
		Last_name:   "Pop",
		Contacts:    "catlittlepop@candies.co",
	})
	fmt.Println(FindAll())
}

func Connect() *sql.DB {
	//Data source information.
	ds := "user=postgres dbname=postgres password=postgres sslmode=disable"
	defer ErrorHandler()
	//openning a database conection
	if conn, err := sql.Open("postgres", ds); err != nil {
		panic(err)
	} else {
		return conn
	}
}

func Insert(person model.Person) bool {
	//Inserting data into person table
	defer ErrorHandler()
	if _, err := Conn.Exec("insert into person(first_name, middle_name, last_name, contacts) values ( $1 ,$2 , $3, $4 )", person.First_name, person.Middle_name, person.Last_name, person.Contacts); err != nil {
		panic(err)
	}
	return true
}

func FindAll() []model.Person {
	people := []model.Person{}

	defer ErrorHandler()
	if rows, err := Conn.Query("select * from person"); err != nil {
		panic(err)
	} else {

		for rows.Next() {
			person := model.Person{}
			rows.Scan(&person.Id, &person.First_name, &person.Middle_name, &person.Last_name, &person.Contacts)
			people = append(people, person)
		}

	}
	return people
}

func ErrorHandler() {
	//Handle an errer when it occurs
	err := recover()
	if err != nil {
		log.Fatal("Something went wrong..", err)
		Conn.Close()
	}
}
