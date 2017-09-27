package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID       string
	Username string
	Fullname string
	Email    string
	Bio      string
	Avatar   string
}

var userList = []User{
	User{ID: "1", Username: "jdoe", Fullname: "John Doe", Email: "jdoe@example.com", Bio: "We shine together", Avatar: "https://google.ca"},
	User{ID: "2", Username: "mprice", Fullname: "John Doe", Email: "mprice@example.com", Bio: "We shine together", Avatar: "https://google.ca"},
	User{ID: "3", Username: "dtow", Fullname: "John Doe", Email: "dtow@example.com", Bio: "We shine together", Avatar: "https://google.ca"},
	User{ID: "4", Username: "wjackson", Fullname: "John Doe", Email: "wjackson@example.com", Bio: "We shine together", Avatar: "https://google.ca"},
	User{ID: "5", Username: "tswift", Fullname: "John Doe", Email: "tswift@example.com", Bio: "We shine together", Avatar: "https://google.ca"},
	User{ID: "6", Username: "dpeter", Fullname: "John Doe", Email: "dpeter@example.com", Bio: "We shine together", Avatar: "https://google.ca"},
}

var GetTalentAllHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// Here we are converting the slice of products to json
	payload, _ := json.Marshal(userList)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
})

// var RegisterHandler = http.HandleFunc(func(w http.ResponseWriter, r *http.Request)) {
// 	vars := mux.Vars(r)
// 	username := vars["username"]
// 	password := vars["password"]

// }

var GetTalentIDHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var user User
	vars := mux.Vars(r) // mux.Vars()
	id := vars["id"]

	for _, u := range userList {
		if u.ID == id {
			user = u
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if user.ID != "" {
		payload, _ := json.Marshal(userList)
		w.Write([]byte(payload))
	} else {
		w.Write([]byte("Talent Not Found"))
	}
})

func main() {

	// new instance gorilla mux router
	r := mux.NewRouter()

	r.Handle("/talent/{id}", GetTalentIDHandler).Methods("GET")
	r.Handle("/talents", GetTalentAllHandler).Methods("GET")

	// r.Handle("/register", RegisterHandler).Methods("POST")

	http.ListenAndServe(":8080", r)
}
