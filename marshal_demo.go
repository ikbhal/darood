package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name        string
	age         int
	active      bool
	lastLoginAt string
}

func main() {
	u2 := User{Name: "Bob", age: 10, active: true, lastLoginAt: "today"}
	u, err := json.Marshal(&u2)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(u)) // {"Name":"Bob","Age":10,"Active":true}
}
