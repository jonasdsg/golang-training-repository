package service

import (
	"database/sql"
	"fmt"
	"golang-training-repository/golang-postgres-conector/model"
	"golang-training-repository/golang-postgres-conector/service/repository"
)

type PersonService struct {
	pr repository.PersonRepository
	db *sql.DB
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
	//openning a database conection
	if conn, err := sql.Open("postgres", ds); err != nil {
		panic(err)
	} else {
		return conn
	}
}

func (ps *PersonService) Save(p model.Person) (int, error) {
	return ps.pr.Save(p)
}

func (ps *PersonService) DeletePerson() {
	var id int
	ps.ListPeople()
	fmt.Println("Choose an Id: ")
	fmt.Scanf("%d", &id)
	ps.pr.Delete(id)
}

func (ps *PersonService) ListPeople() {
	if pp, err := ps.pr.FindAll(); err != nil {
		panic(err)
	} else {

		for _, p := range pp {

			fmt.Println(p.Id, "-", p.First_name, p.Middle_name, p.Last_name, ",", p.Contacts)
		}
	}
}

func (ps *PersonService) New() {
	pr := repository.PersonRepository{}
	ps.db = Connect()
	pr.New(ps.db)
	ps.pr = pr
}
