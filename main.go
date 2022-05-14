package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

const file string = "darood.db"

type Darood struct {
	Id          int
	Counter     int
	Name        string
	Description string
}

var counter = 10000000 // crore
const croreDaroodName = "crore darood"

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Crore darood!")
	fmt.Println("Endpoint Hit: homePage")
}

func getCounterDbHelper(name string) (Darood, error) {
	db, err := sql.Open("sqlite3", file)

	if err != nil {
		fmt.Println("Unable to open darood.db sqlite db")
	}

	fmt.Println("inside getCounter Db Helper")

	row, err := db.Query("SELECT * FROM darood WHERE name=?", name)
	darood := Darood{}
	if err != nil {
		return Darood{}, err
	}

	daroodId := 1
	daroodCounter := 100000000
	daroodName := "crore darood"
	daroodDescription := "Reach target by decrement"

	if err = row.Scan(&daroodId, &daroodCounter, &daroodName, &daroodDescription); err == sql.ErrNoRows {
		fmt.Printf(" no darood row found matching requirements")
		return Darood{}, sql.ErrNoRows
	}

	darood.Id = daroodId
	darood.Counter = daroodCounter
	darood.Name = daroodName
	darood.Description = daroodDescription

	return darood, err
}

func getCounterDb(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside getCounter Db")

	darood, err := getCounterDbHelper(croreDaroodName)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	jsonResp, err := json.Marshal(darood)
	if err != nil {
		log.Fatalf("Error happend in db darood table ;counter retrieve JSON Marshal. Err:%s", err)
	}
	w.Write(jsonResp)
	return
}

func getCounter(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]interface{})
	resp["message"] = "Counter value retrieve"
	resp["counter"] = counter
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happend in counter retrieve JSON Marshal. Err:%s", err)
	}
	w.Write(jsonResp)
	return
}

func decrementCounter(w http.ResponseWriter, r *http.Request) {
	counter--
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]interface{})
	resp["message"] = "Counter decremented"
	resp["counter"] = counter
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happend in JSON marshal. Err: %s", err)

	}
	w.Write(jsonResp)
	return
}

func handleRequests() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// http.HandleFunc("/api/v1/counter", decrementCounter)
	// http.HandleFunc("/api/v1/counter", decrementCounter).Methods("PUT")
	// http.HandleFunc("/api/v1/counter", getCounter).Methods("GET")
	// http.HandleFunc("/api/v2/counter", getCounterDb).Methods("GET")

	http.HandleFunc("/api/v3/test", homePage)

	//http.HandleFunc("/", homePage)
	fmt.Println("server startging at 3010")
	http.ListenAndServe(":3010", nil)
	fmt.Println("started server at 3010")
}

func main() {
	handleRequests()
}
