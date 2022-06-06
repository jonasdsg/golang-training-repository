package main

import (
	"database/sql"
	"fmt"
	"golang-training-repository/golang-postgres-conector/service"
	"strings"

	_ "github.com/lib/pq"
)

var Conn *sql.DB

func main() {
	var choice string
	ps := service.PersonService{}
	ps.New()
	fmt.Print("Choose one action\n[I]nsert\n[D]elete\n[L]ist\nType :")
	fmt.Scanf("%s", &choice)
	switch strings.ToUpper(choice) {
	case "I":
		ps.Save(service.CreatePerson())
	case "D":
		ps.DeletePerson()
	case "L":
		ps.ListPeople()
	}
}
