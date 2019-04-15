package main

import (
	"encoding/json"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/nurali-techie/meetup-golang/microservice-auth/auth-service/token"
	log "github.com/sirupsen/logrus"
)

type User struct {
	ID       string
	Username string
	Password string
	Email    string
}

var users map[string]*User

func main() {
	initData()

	http.HandleFunc("/api/login", login)
	addr := "localhost:8081"
	log.Infof("auth service running at:%s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("failed to start auth service, error:%v", err)
	}
}

func login(rw http.ResponseWriter, r *http.Request) {
	log.Infof("login request")
	if r.Method != http.MethodPost {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var userIn User
	err := decoder.Decode(&userIn)
	if err != nil {
		log.Errorf("json to user failed, err:%v", err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	user := users[userIn.Username]
	if user == nil {
		log.Errorf("user not found, username:%s", userIn.Username)
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	if user.Password != userIn.Password {
		log.Errorf("invalid password")
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	rawToken := getToken(user)
	rw.Write(rawToken)
}

func getToken(user *User) []byte {
	claim := jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
		"role":  "customer",
	}
	token := token.GenerateToken(claim)
	if token == nil {
		return nil
	}
	rawToken, err := json.Marshal(token)
	if err != nil {
		return nil
	}
	return rawToken
}

func initData() {
	users = make(map[string]*User)

	users["ali"] = &User{ID: "1", Username: "ali", Password: "abcd1234", Email: "ali@gmail.com"}
	users["mary"] = &User{ID: "2", Username: "mary", Password: "test1234", Email: "mary@yahoo.com"}
}
