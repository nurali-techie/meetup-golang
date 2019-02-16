package main

import (
	"context"
	"log"
	"net"

	pb "github.com/nurali-techie/meetup-golang/microservice/v3/model"
	"google.golang.org/grpc"
)

func main() {
	log.Println("role service v3")
	l, _ := net.Listen("tcp", ":8090")
	s := grpc.NewServer()
	pb.RegisterRoleServiceServer(s, &RoleService{})
	log.Fatal(s.Serve(l))
}

type RoleService struct {
}

func (r *RoleService) GetRole(ctx context.Context, req *pb.GetRoleRequest) (*pb.GetRoleReply, error) {
	log.Printf("got request for user id:%s\n", req.UserID)
	userRole := roles[req.UserID]
	return &pb.GetRoleReply{Role: userRole.Role}, nil
}

type UserRole struct {
	UserID string
	Role   string
}

var roles = map[string]*UserRole{
	"1": &UserRole{UserID: "1", Role: "developer"},
	"2": &UserRole{UserID: "2", Role: "admin"},
}
