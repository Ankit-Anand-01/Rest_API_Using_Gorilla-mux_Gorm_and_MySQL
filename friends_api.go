package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Friends struct {
	gorm.Model
	Name      string
	ContactNo string
}

var db *gorm.DB
var err error

func main() {
	router := mux.NewRouter()

	db, err = gorm.Open("mysql", "root:lucifer007@tcp(127.0.0.1:3306)/batch2020?charset=utf8&parseTime=True")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	db.AutoMigrate(&Friends{})

	router.HandleFunc("/friends", Getfriends).Methods("GET")
	router.HandleFunc("/friends/{id}", Getfriend).Methods("GET")
	router.HandleFunc("/friends", Addfriends).Methods("POST")
	router.HandleFunc("/friends/{id}", Removefriends).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":1956", router))
}

func Getfriends(w http.ResponseWriter, r *http.Request) {
	var friend []Friends
	db.Find(&friend)
	json.NewEncoder(w).Encode(&friend)
}

func Getfriend(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var friend Friends
	db.First(&friend, params["id"])
	json.NewEncoder(w).Encode(&friend)
}

func Addfriends(w http.ResponseWriter, r *http.Request) {
	var friend Friends
	json.NewDecoder(r.Body).Decode(&friend)
	db.Create(&friend)
	json.NewEncoder(w).Encode(&friend)
}

func Removefriends(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var friend Friends
	db.First(&friend, params["id"])
	db.Delete(&friend)

	var friend1 []Friends
	db.Find(&friend)
	json.NewEncoder(w).Encode(&friend1)
}
