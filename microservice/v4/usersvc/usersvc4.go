package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	pb "github.com/nurali-techie/meetup-golang/microservice/v4/model"
	"google.golang.org/grpc"
)

func main() {
	log.Println("user service v4")
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
	conn, _ := grpc.Dial("localhost:8090", grpc.WithInsecure())
	defer conn.Close()

	client := pb.NewRoleServiceClient(conn)

	req := pb.GetRoleRequest{UserID: userID}
	reply, _ := client.GetRole(context.Background(), &req)

	return reply.Role
}
