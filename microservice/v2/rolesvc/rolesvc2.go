package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/nurali-techie/meetup-golang/microservice/v2/model"
)

func main() {
	log.Println("role service v2")
	rs := new(RoleService) // Call "RoleService.GetRole"
	rpc.Register(rs)
	rpc.HandleHTTP()
	l, _ := net.Listen("tcp", ":8090")
	log.Fatal(http.Serve(l, nil))
}

type RoleService struct {
}

func (r *RoleService) GetRole(req model.GetRoleRequest, reply *model.GetRoleReply) error {
	log.Printf("got request for user id:%s\n", req.UserID)
	userRole := roles[req.UserID]
	reply.Role = userRole.Role
	return nil
}

type UserRole struct {
	UserID string
	Role   string
}

var roles = map[string]*UserRole{
	"1": &UserRole{UserID: "1", Role: "developer"},
	"2": &UserRole{UserID: "2", Role: "admin"},
}
