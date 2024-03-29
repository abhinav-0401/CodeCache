package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

type Person struct {
	Id        int64
	FirstName string
	LastName  string
	Email     sql.NullString
	Gender    string
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("idk")

	rows, err := app.db.Query("SELECT id, first_name, last_name, gender, email FROM person;")
	if err != nil {
		app.errorLog.Println(err)
		return
	}
	defer rows.Close()

	var people = []*Person{}

	for rows.Next() {
		var person = &Person{}
		err = rows.Scan(&person.Id, &person.FirstName, &person.LastName, &person.Gender, &person.Email)
		if err != nil {
			app.errorLog.Println(err)
			return
		}

		people = append(people, person)
		s := fmt.Sprintf("%v %v %v %v %v", person.Id, person.FirstName, person.LastName, person.Gender, person.Email.String)
		fmt.Fprintln(w, s)
	}
}
