package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

const file string = "./darood.db"

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

	defer db.Close()

	fmt.Println("inside getCounter Db Helper")

	rows, err := db.Query("SELECT id, counter, name, description FROM darood WHERE name='crore darood'")
	defer rows.Close()
	rows.Next()
	darood := Darood{}
	if err != nil {
		return Darood{}, err
	}
	if err = rows.Scan(&darood.Id, &darood.Counter, &darood.Name, &darood.Description); err == sql.ErrNoRows {
		fmt.Printf(" no darood row found matching requirements")
		return Darood{}, sql.ErrNoRows
	}
	return darood, err
}

func getCounterDb(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside getCounter Db")

	darood, err := getCounterDbHelper(croreDaroodName)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Printf("with fmt print darood %+v", darood)
	jsonResp, err := json.Marshal(&darood)
	fmt.Println("Json marshal darood print")
	fmt.Println(string(jsonResp))
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

	http.HandleFunc("/api/v1/counter/decr", decrementCounter)
	// http.HandleFunc("/api/v1/counter", decrementCounter).Methods("PUT")
	http.HandleFunc("/api/v1/counter/get", getCounter)
	http.HandleFunc("/api/v2/counter/get", getCounterDb)

	http.HandleFunc("/api/v3/test", homePage)

	//http.HandleFunc("/", homePage)
	fmt.Println("server startging at 3010")
	http.ListenAndServe(":3010", nil)
	fmt.Println("started server at 3010")
}

func main() {
	handleRequests()
}
