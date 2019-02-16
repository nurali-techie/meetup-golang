package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/rpc"

	"github.com/nurali-techie/meetup-golang/microservice/v2/model"
)

func main() {
	log.Println("user service v2")
	http.HandleFunc("/users", GetUser)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func GetUser(res http.ResponseWriter, req *http.Request) {
	userID := req.URL.Query().Get("id")
	log.Printf("got request for user id:%s\n", userID)

	user := users[userID]
	user.Role = getRole(user.ID)
	content, _ := json.Marshal(user)
	res.Write(content)
}

type User struct {
	ID   string
	Name string
	Role string
}

var users = map[string]*User{
	"1": &User{ID: "1", Name: "Mani"},
	"2": &User{ID: "2", Name: "Hemal"},
}

func getRole(userID string) string {
	client, _ := rpc.DialHTTP("tcp", "localhost:8090")

	req := model.GetRoleRequest{UserID: userID}
	reply := model.GetRoleReply{}
	client.Call("RoleService.GetRole", req, &reply)

	return reply.Role
}
