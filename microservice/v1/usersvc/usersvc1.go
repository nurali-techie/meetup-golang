package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	log.Println("monolith user service v1")
	http.HandleFunc("/users", GetUser)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func GetUser(res http.ResponseWriter, req *http.Request) {
	userID := req.URL.Query().Get("id")
	log.Printf("got request for user id:%s\n", userID)

	user := users[userID]
	content, _ := json.Marshal(user)
	res.Write(content)
}

type User struct {
	ID   string
	Name string
	Role string
}

var users = map[string]*User{
	"1": &User{ID: "1", Name: "Mani", Role: "developer"},
	"2": &User{ID: "2", Name: "Hemal", Role: "admin"},
}
