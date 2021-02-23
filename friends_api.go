package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
//structure for table
type Friends struct {
	gorm.Model
	Name      string
	ContactNo string
}
//gorm variable
var db *gorm.DB
var err error

func main() {
	router := mux.NewRouter()
       // opening database
	db, err = gorm.Open("mysql", "root:lucifer007@tcp(127.0.0.1:3306)/batch2020?charset=utf8&parseTime=True")

	if err != nil {
		panic(err)
	}
       //closing database
	defer db.Close()

	db.AutoMigrate(&Friends{})

	router.HandleFunc("/friends", Getfriends).Methods("GET")
	router.HandleFunc("/friends/{id}", Getfriend).Methods("GET")
	router.HandleFunc("/friends", Addfriends).Methods("POST")
	router.HandleFunc("/friends/{id}", Removefriends).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":1956", router))
}
//getting all friends details
func Getfriends(w http.ResponseWriter, r *http.Request) {
	var friend []Friends
	db.Find(&friend)
	json.NewEncoder(w).Encode(&friend)
}
//getting a perticular friend details using id
func Getfriend(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var friend Friends
	db.First(&friend, params["id"])
	json.NewEncoder(w).Encode(&friend)
}
//adding new friend details
func Addfriends(w http.ResponseWriter, r *http.Request) {
	var friend Friends
	json.NewDecoder(r.Body).Decode(&friend)
	db.Create(&friend)
	json.NewEncoder(w).Encode(&friend)
}
//removing a friend details using id
func Removefriends(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var friend Friends
	db.First(&friend, params["id"])
	db.Delete(&friend)

	var friend1 []Friends
	db.Find(&friend)
	json.NewEncoder(w).Encode(&friend1)
}
