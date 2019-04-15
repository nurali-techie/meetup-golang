package main

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

type Review struct {
	Reviewer string
	Rating   int
	Comment  string
}

var bookReviews map[string][]*Review

func main() {
	initData()

	// register http endpoint and start http service
	http.HandleFunc("/api/reviews", showReviews)
	addr := "localhost:8083"
	log.Infof("review service running at:%s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("failed to start review service, error:%v", err)
	}
}

func showReviews(rw http.ResponseWriter, r *http.Request) {
	bookID := r.URL.Query().Get("bookId")
	limit := getLimit(r.URL.Query().Get("limit"))
	log.Infof("show reviews request, bookID:%s, limit:%d", bookID, limit)

	reviews := bookReviews[bookID]
	if limit != 0 && limit < len(reviews) {
		reviews = reviews[:limit]
	}

	payload, err := json.Marshal(reviews)
	if err != nil {
		log.Errorf("reviews to json failed, error:%v", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.Write([]byte(payload))
}

func getLimit(limit string) int {
	if limit != "" {
		return cast.ToInt(limit)
	}
	return 0
}

func initData() {
	bookReviews = make(map[string][]*Review)

	bookReviews["1"] = []*Review{
		&Review{Reviewer: "Hemal", Rating: 5, Comment: "Great book with real life example"},
		&Review{Reviewer: "Mani", Rating: 5, Comment: "Nice content and explaination"},
		&Review{Reviewer: "John", Rating: 4, Comment: "Good book but some repetation"},
		&Review{Reviewer: "Lacy", Rating: 5, Comment: "Love this book"},
	}

	bookReviews["2"] = []*Review{
		&Review{Reviewer: "Adil", Rating: 5, Comment: "Go basics nicely explained"},
		&Review{Reviewer: "Bunty", Rating: 4, Comment: "Good Go topics coverage"},
	}
}
