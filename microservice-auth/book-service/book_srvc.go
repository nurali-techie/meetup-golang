package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type Book struct {
	ID      string
	Name    string
	Author  string
	Price   float32
	Reviews []*Review
}

type Review struct {
	Reviewer string
	Rating   int
	Comment  string
}

var books map[string]*Book

func main() {
	initData()

	// register http endpoint and start http service
	http.HandleFunc("/api/books", showBooks)
	addr := "localhost:8082"
	log.Infof("book service running at:%s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("failed to start book service, error:%v", err)
	}
}

func showBooks(rw http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	log.Infof("show books request, id:%s", id)

	book := books[id]
	if book == nil {
		log.Errorf("book not found with id '%s'", id)
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	book.Reviews = getReviews(id)

	payload, err := json.Marshal(book)
	if err != nil {
		log.Errorf("book to json failed, error:%v", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.Write([]byte(payload))
}

func getReviews(id string) []*Review {
	url := fmt.Sprintf("http://localhost:8083/api/reviews?bookId=%s&limit=2", id)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Errorf("create reqeust failed, err:%v", err)
		return nil
	}

	httpClient := http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Errorf("review request failed, err:%v", err)
		return nil
	}

	// process response
	decoder := json.NewDecoder(resp.Body)
	var reviews []*Review
	err = decoder.Decode(&reviews)
	if err != nil {
		log.Errorf("json to reviews failed, err:%v", err)
		return nil
	}
	return reviews
}

func initData() {
	books = make(map[string]*Book)

	books["1"] = &Book{
		ID:     "1",
		Name:   "Microservices Patterns",
		Author: "Chris Richardson",
		Price:  50.0,
	}

	books["2"] = &Book{
		ID:     "2",
		Name:   "Go in Action",
		Author: "William Kennedy",
		Price:  40.0,
	}
}
