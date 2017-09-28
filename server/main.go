package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	. "./models"
	"github.com/gorilla/mux"
	"github.com/thermokarst/jwt"
	mgo "gopkg.in/mgo.v2"
)

//const MongoDb details

var info = &mgo.DialInfo{
	Addrs:    []string{hosts},
	Timeout:  60 * time.Second,
	Database: database,
	Username: username,
	Password: password,
}

var session, errDial = mgo.DialWithInfo(info)
var col = session.DB("stalents").C("talents")

var GetTalentAllHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	// Get all items in talent collection
	if errDial != nil {
		panic(errDial)
	}

	var results []Talent

	err := col.Find(nil).All(&results)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Results All: ", results)
	}
	payload, _ := json.Marshal(results)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
})

// var RegisterHandler = http.HandleFunc(func(w http.ResponseWriter, r *http.Request)) {
// 	vars := mux.Vars(r)
// 	username := vars["username"]
// 	password := vars["password"]

// }

var GetTalentIDHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var talent Talent
	vars := mux.Vars(r) // mux.Vars()
	id := vars["id"]

	if errDial != nil {
		panic(errDial)
	}

	var results []Talent

	err := col.Find(nil).All(&results)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Results All: ", results)
	}

	for _, u := range results {
		if u.ID == id {
			talent = u
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if talent.ID != "" {
		payload, _ := json.Marshal(talent)
		w.Write([]byte(payload))
	} else {
		w.Write([]byte("Talent Not Found"))
	}
})

func auth(email string, password string) error {
	// Hard-code a user
	if email != "test" || password != "test" {
		return errors.New("invalid credentials")
	}
	return nil
}

func setClaims(id string) (map[string]interface{}, error) {
	currentTime := time.Now()
	return map[string]interface{}{
		"iat": currentTime.Unix(),
		"exp": currentTime.Add(time.Minute * 60 * 24).Unix(),
	}, nil
}

func verifyClaims(claims []byte, r *http.Request) error {
	currentTime := time.Now()
	var c struct {
		Iat int64
		Exp int64
	}
	_ = json.Unmarshal(claims, &c)
	if currentTime.After(time.Unix(c.Exp, 0)) {
		return errors.New("this token has expired")
	}
	return nil
}

func main() {

	// clear all the talent records
	col.RemoveAll(nil)

	// Add new talents records
	err := col.Insert(
		&Talent{ID: "1", Username: "jdoe", Fullname: "John Doe", Email: "jdoe@example.com", Bio: "We shine together", Avatar: "https://google.ca"},
		&Talent{ID: "2", Username: "mprice", Fullname: "John Doe", Email: "mprice@example.com", Bio: "We shine together", Avatar: "https://google.ca"},
		&Talent{ID: "3", Username: "dtow", Fullname: "John Doe", Email: "dtow@example.com", Bio: "We shine together", Avatar: "https://google.ca"},
		&Talent{ID: "4", Username: "wjackson", Fullname: "John Doe", Email: "wjackson@example.com", Bio: "We shine together", Avatar: "https://google.ca"},
		&Talent{ID: "5", Username: "tswift", Fullname: "John Doe", Email: "tswift@example.com", Bio: "We shine together", Avatar: "https://google.ca"},
		&Talent{ID: "6", Username: "dpeter", Fullname: "John Doe", Email: "dpeter@example.com", Bio: "We shine together", Avatar: "https://google.ca"},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Count all items in talent collection
	count, err2 := col.Count()
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(fmt.Sprintf("Talents count: %d", count))

	// Setup JWT
	config := &jwt.Config{
		Secret: "password",
		Auth:   auth,
		Claims: setClaims,
	}
	j, errJWT := jwt.New(config)
	if err != nil {
		panic(errJWT)
	}

	// new instance gorilla mux router
	r := mux.NewRouter()

	r.Handle("/auth", j.Authenticate())
	r.Handle("/talent/{id}", GetTalentIDHandler).Methods("GET")
	r.Handle("/talents", GetTalentAllHandler).Methods("GET")

	// r.Handle("/register", RegisterHandler).Methods("POST")

	http.ListenAndServe(":8080", r)
}
