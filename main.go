package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	// "gorilla/mux"
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

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to the Crore darood!")
// 	fmt.Println("Endpoint Hit: homePage")
// }

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

func decrementCounterDbHelper(name string) {
	fmt.Println("inside decremenetCounterDbHelper")
	db, err := sql.Open("sqlite3", file)

	if err != nil {
		fmt.Println("Unable to open darood.db sqlite db")
	}

	defer db.Close()

	stmt, err := db.Prepare("UPDATE darood set counter = counter - 1 where name= ?")
	defer stmt.Close()

	stmt.Exec(name)
}

func decrementCounterDb(w http.ResponseWriter, r *http.Request) {
	decrementCounterDbHelper(croreDaroodName) // for now crore darood
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

func incrementCounterDbHelper(name string) {
	fmt.Println("inside incrementCounterDbHelper")
	db, err := sql.Open("sqlite3", file)

	if err != nil {
		fmt.Println("Unable to open darood.db sqlite db")
	}

	defer db.Close()

	stmt, err := db.Prepare("UPDATE darood set counter = counter + 1 where name= ?")
	defer stmt.Close()

	stmt.Exec(name)
}

func incrementCounterDb(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside incrementCounterDb")
	incrementCounterDbHelper(croreDaroodName) // for now crore darood
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]interface{})
	resp["message"] = "Counter incremeneted"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happend in JSON marshal. Err: %s", err)

	}
	w.Write(jsonResp)
	return
}

func decrementCounter(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]interface{})
	resp["message"] = "Counter decremented"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happend in JSON marshal. Err: %s", err)

	}
	w.Write(jsonResp)
	return
}

func registerUserFormShow(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside registerUserFormHandler")

	fmt.Fprintf(w, "User registration form show")
	// redirect to  user_register form in static
}

func registerUserFormHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside registerUserFormHandler")
	// create row in user table

	fmt.Fprintf(w, "User registration form handler")
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
	first_name := r.FormValue("first_name")
	last_name := r.FormValue("last_name")
	username := r.FormValue("username")
	mobile := r.FormValue("mobile")
	email := r.FormValue("email")
	fmt.Fprintf(w, "First Name = %s\n", first_name)
	fmt.Fprintf(w, "Last Name = %s\n", last_name)
	fmt.Fprintf(w, "Username = %s\n", username)
	fmt.Fprintf(w, "mobile = %s\n", mobile)
	fmt.Fprintf(w, "email = %s\n", email)

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside home page function")
	const newUrl = "http://darood.tauqeer.in/register.html"
	http.Redirect(w, r, newUrl, http.StatusSeeOther)
}

// var router = mux.NewRouter()

func handleRequests() {

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/api/v1/counter/decr", decrementCounter)
	http.HandleFunc("/api/v2/counter/decr", decrementCounterDb)
	http.HandleFunc("/api/v2/counter/incr", incrementCounterDb)
	// http.HandleFunc("/api/v1/counter", decrementCounter).Methods("PUT")
	http.HandleFunc("/api/v1/counter/get", getCounter)
	http.HandleFunc("/api/v2/counter/get", getCounterDb)

	http.HandleFunc("/api/v3/test", homePage)

	// user regiser
	http.HandleFunc("/api/v1/users/register/form", registerUserFormShow)
	http.HandleFunc("/api/v1/users/handleRegister", registerUserFormHandler)

	http.HandleFunc("/home", homePage)
	// user login

	//http.HandleFunc("/", homePage)
	fmt.Println("server startging at 3010")
	http.ListenAndServe(":3010", nil)
	fmt.Println("started server at 3010")
}

func main() {
	handleRequests()
}
