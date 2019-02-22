package main

import (
	"database/sql"
	"encoding/json"
)

//PersonRepository requires a database connection to be built
type PersonRepository struct {
	database *sql.DB
}

//uses our database connection to return a list of Person structs representing the data in our database
func (repository *PersonRepository) FindAll() []*Person {
	rows, _ := repository.database.Query(
		`SELECT id, name, age FROM people;`,
	)
	defer rows.Close()

	people := []*Person{}

	for rows.Next() {
		var (
			id   int
			name string
			age  int
		)

		rows.Scan(&id, &name, &age)

		people = append(people, &Person{
			Id:   id,
			Name: name,
			Age:  age,
		})
	}

	return people
}

func NewPersonRepository(database *sql.DB) *PersonRepository {
	return &PersonRepository{database: database}
}
