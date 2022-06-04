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
	Insert(CreatePerson())
	fmt.Println(FindAll())
}

func CreatePerson() model.Person {
	p := model.Person{}
	fmt.Print("First name: ")
	fmt.Scanf("%s", &p.First_name)
	fmt.Print("Middle name: ")
	fmt.Scanf("%s", &p.Middle_name)
	fmt.Print("Last name: ")
	fmt.Scanf("%s", &p.Last_name)
	fmt.Print("Contact: ")
	fmt.Scanf("%s", &p.Contacts)
	return p
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
