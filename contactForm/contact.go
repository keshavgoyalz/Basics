package main

import (
	"database/sql"
	"fmt"
	_ "go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
)

type Details struct {
	NAME    string
	EMAIL   string
	SUBJECT string
	MESSAGE string
	PHONE   string
}

func main() {

	/*
		SQL DATABASE
	*/
	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/database_name")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query :=
		`CREATE TABLE form_details (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    subject VARCHAR(255) NOT NULL,
    message TEXT NOT NULL,
    phone VARCHAR(20) NOT NULL
);`

	_, err = db.Exec(query)

	var list []Details
	temp := template.Must(template.ParseFiles("forms.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			temp.Execute(w, nil)
			return
		}
		details := Details{
			NAME:    r.FormValue("name"),
			EMAIL:   r.FormValue("email"),
			SUBJECT: r.FormValue("subject"),
			MESSAGE: r.FormValue("message"),
			PHONE:   r.FormValue("phoneNumber"),
		}
		list = append(list, details)
		result, err := db.Exec(`INSERT INTO form_details(name, email, subject, message, phoneNumber) VALUES (?, ?, ?, ?, ?)`, details.NAME, details.EMAIL, details.SUBJECT, details.MESSAGE, details.PHONE)
		if err != nil {
			log.Fatal(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Inserted data with ID: %d\n", id)
		temp.Execute(w, struct{ Success bool }{true})

	})

	http.ListenAndServe(":80", nil)
}
