package repository

import (
	"database/sql"
	"golang-training-repository/golang-postgres-conector/model"
)

type PersonRepository struct {
	db *sql.DB
}

func (pr *PersonRepository) insert(person model.Person) (int, error) {
	var id int
	var query string = "INSERT INTO person(first_name, middle_name, last_name, contacts) VALUES ( $1 ,$2 , $3, $4 ) RETURNING id"

	if tx, err := pr.db.Begin(); err != nil {
		tx.Rollback()

		return id, err
	} else {
		stmt, err := tx.Prepare(query)

		if err != nil {
			tx.Rollback()
		} else {
			stmt.QueryRow(person.First_name, person.Middle_name, person.Last_name, person.Contacts).Scan(&id)
			tx.Commit()
			stmt.Close()
		}

		return id, err
	}
}

func (pr *PersonRepository) Delete(id int) (bool, error) {
	if _, err := pr.db.Exec("DELETE FROM person WHERE id = $1", id); err != nil {
		return false, err
	}
	return true, nil
}

func (pr *PersonRepository) FindAll() ([]model.Person, error) {
	people := []model.Person{}

	if rows, err := pr.db.Query("SELECT * FROM person"); err != nil {
		return nil, err
	} else {

		for rows.Next() {
			person := model.Person{}
			rows.Scan(&person.Id, &person.First_name, &person.Middle_name, &person.Last_name, &person.Contacts)
			people = append(people, person)
		}

	}
	return people, nil
}

func (pr *PersonRepository) FindById(id int) (model.Person, error) {
	row := pr.db.QueryRow("SELECT * FROM person WHERE id = $id", id)
	p := model.Person{}

	if err := row.Scan(&p.Id, &p.First_name, &p.Middle_name, &p.Last_name, &p.Contacts); err != nil {
		return p, err
	}

	return p, nil
}

func (pr *PersonRepository) Save(person model.Person) (int, error) {
	query := "UPDATE person SET first_name = $1, middle_name = $2, last_name = $3, contact = $4 WHERE id = $5"
	if p, err := pr.FindById(person.Id); err != nil {
		return pr.insert(person)
	} else {
		var id int
		pr.db.QueryRow(query, person.First_name, person.Middle_name, person.Last_name, person.Contacts, p.Id).Scan(&id)
		return id, nil
	}
}

func (p *PersonRepository) New(db *sql.DB) {
	p.db = db
}
