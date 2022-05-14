package main

import (
       "fmt"
       "log"
       "net/http"
       "json"
)

var counter =10000000 // crore

func homePage(w http.ResponseWriter, r * http.Request){
     fmt.Fprintf(w, "Welcome to the Crore darood!")
     fmt.Println("Endpoint Hit: homePage")
 }

func getCounter(w http.ResponseWriter, r* http.Request) {
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

func decrementCounter(w http.ResponseWriter, r *http.Request){
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
     http.HandleFunc("/api/v1/counter", decrementCounter).Methods("PUT")
     http.HandleFunc("/api/v1/counter", getCounter).Methods("GET")
     //http.HandleFunc("/", homePage)
     fmt.Println("server startging at 3010")
     http.ListenAndServe(":3010", nil)
     fmt.Println("started server at 3010")
}

func main() {
     handleRequests()
}
