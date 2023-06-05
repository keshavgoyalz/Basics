package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

type Employees struct {
	NAME    string
	EMAIL   string
	PHONE   string
	ADDRESS string
}

const dbaddress = "root:CSD@mysql-1872@(127.0.0.1:3306)/employees_details?charset=utf8mb4&parseTime=True&loc=Local"

func initializeRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/employees", GetEmployees).Methods("GET")
	r.HandleFunc("/employees/{id}", GetEmployeesID).Methods("GET")
	r.HandleFunc("/employees", CreateEmployees).Methods("POST")
	r.HandleFunc("/employees/{id}", UpdateEmployees).Methods("PUT")
	r.HandleFunc("/employees/{id}", DeleteEmployees).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
func initialMigration() {
	DB, err = gorm.Open(mysql.Open(dbaddress), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("database connection not valid")
	}

	DB.AutoMigrate(&Employees{})
}

func CreateEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employees Employees
	json.NewDecoder(r.Body).Decode(&employees)
	DB.Create(&employees)
	json.NewEncoder(w).Encode(employees)
}

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employees Employees
	DB.Find(&employees)
	json.NewEncoder(w).Encode(employees)
}

func GetEmployeesID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employees Employees
	id := mux.Vars(r)
	DB.First(&employees, id["id"])
	json.NewEncoder(w).Encode(employees)
}

func UpdateEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employees Employees
	id := mux.Vars(r)
	DB.First(&employees, id["id"])
	json.NewDecoder(r.Body).Decode(&employees)
	DB.Save(&employees)
	json.NewEncoder(w).Encode(employees)
}

func DeleteEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employees Employees
	id := mux.Vars(r)
	DB.Delete(&employees, id["id"])
	json.NewEncoder(w).Encode("Deleted")
}
