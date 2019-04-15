package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/nurali-techie/meetup-golang/microservice-auth/api-gateway/token"

	log "github.com/sirupsen/logrus"
)

func main() {
	// register http endpoint and start http service
	http.HandleFunc("/", router)
	addr := "localhost:8080"
	log.Infof("api gateway running at:%s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("failed to start api gateway, error:%v", err)
	}
}

func router(rw http.ResponseWriter, r *http.Request) {
	reqPath := r.URL.Path
	log.Infof("reqeust path:'%s'", reqPath)

	switch reqPath {
	case "/api/login":
		forward(rw, r, "http://localhost:8081")
	case "/api/books":
		authOk := authenticate(rw, r)
		if authOk {
			forward(rw, r, "http://localhost:8082")
		}
	case "/api/reviews":
		authOk := authenticate(rw, r)
		if authOk {
			forward(rw, r, "http://localhost:8083")
		}
	default:
		rw.Write([]byte(fmt.Sprintf("routing not found, path:'%s'", reqPath)))
	}
}

func forward(rw http.ResponseWriter, r *http.Request, host string) {
	// create request
	fwURL := fmt.Sprintf("%s%s", host, r.RequestURI)
	fwReq, err := http.NewRequest(r.Method, fwURL, r.Body)
	if err != nil {
		log.Errorf("failed to create request, url:'%s'", fwURL)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	// set headers
	fwReq.Header = make(http.Header)
	for h, val := range r.Header {
		fwReq.Header[h] = val
	}

	// execute request
	httpClient := http.Client{}
	resp, err := httpClient.Do(fwReq)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadGateway)
		return
	}

	// process response
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	headerResp := strings.Join(resp.Header["Content-Type"], "")
	rw.Header().Set("Content-Type", headerResp)
	rw.Write([]byte(body))
}

func authenticate(rw http.ResponseWriter, r *http.Request) bool {
	tokenStr := extractToken(r.Header.Get("Authorization"))
	token := token.ValidateToken(tokenStr)
	if token == nil {
		rw.WriteHeader(http.StatusUnauthorized)
		return false
	}
	return true
}

func extractToken(authHeader string) string {
	parts := strings.Split(authHeader, " ")
	if len(parts) >= 2 {
		return parts[1]
	}
	return ""
}
