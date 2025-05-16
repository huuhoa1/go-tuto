package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/* {
	"user_id": 12345,
	"name":"User A",
	"age": 35,
	"password": "my-password",
	"roles": ["admin", "collaborator"]
} */

type User struct {
	ID          int      `json:"user_id"`
	Name        string   `json:"name,omitempty"`
	Age         int      `json:"age"`
	Password    string   `json:"-"`
	Permissions []string `json:"roles"`
}

type Person struct {
	Name   string `json:"username"`
	Age    int    `json:"age"`
	Active bool   `json:"is_active"`
}

func main() {
	person := Person{
		Name:   "Can Huynh",
		Age:    59,
		Active: true,
	}

	f, err := os.Create("output.json")
	if err != nil {
		fmt.Println("Error creating file!", err)
		panic(err)
	}
	defer f.Close()

	encoder := json.NewEncoder(f) // takes as parameter a pointer to a struct that implements io.Writer
	err = encoder.Encode(person)  //serialize the object; encode the object in characters
	if err != nil {
		fmt.Println("Error encoding person:", err)
		panic(err)
	}
}
