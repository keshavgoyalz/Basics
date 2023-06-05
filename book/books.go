package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

type Books struct {
	ID        int
	NAME      string
	SUBJECT   string
	AUTHOR    string
	PUBLISHER string
	DATE      string
}

const dbAddress = "root:CSD@mysql-1872@(host.docker.internal:3306)/book_details?parseTime=True&loc=Local"

func initializeRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/books", GetBooks).Methods("GET")
	r.HandleFunc("/books", CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}

func initialMigration() {
	DB, err = gorm.Open(mysql.Open(dbAddress), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Database not valid")
	}

	DB.AutoMigrate(&Books{})
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var books []Books
	DB.Find(&books)
	json.NewEncoder(w).Encode(books)

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var books Books
	json.NewDecoder(r.Body).Decode(&books)
	DB.Create(&books)
	json.NewEncoder(w).Encode(books)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var books Books
	id := mux.Vars(r)
	DB.First(&books, id["id"])
	json.NewDecoder(r.Body).Decode(&books)
	DB.Save(&books)
	json.NewEncoder(w).Encode(books)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var books Books
	id := mux.Vars(r)
	DB.Delete(&books, id["id"])
	json.NewEncoder(w).Encode(books)

}
