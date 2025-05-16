package main

import (
	"encoding/json"
	"fmt"
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
	Name       string   `json:"full_name"`
	Age        int      `json:"years_old,omitempty"`
	Occupation string   `json:"occupation,-"`
	Languages  []string `json:"spoken_languages"`
}

func main() {
	jsonData := `{"full_name":"Jane Doe", "years_old":25, "spoken_languages":["French","English"], "occupation":"teacher"}` //`` syntax to avoid \"

	var person Person
	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		panic(err)
	}
	fmt.Println("Name", person.Name)
	fmt.Println("Age", person.Age)
	fmt.Println("Languages", person.Languages)
	fmt.Println("Occupation", person.Occupation)

	/* u := User{
		ID:          1,
		Name:        "Can Huynh",
		Age:         20,
		Password:    "my-password",
		Permissions: []string{"admin", "group-member"},
	}
	b, err := json.Marshal(u)
	if err != nil {
		fmt.Println("error marshalling JSON: ", err)
		panic(err)
	}
	fmt.Println(string(b)) */
}
