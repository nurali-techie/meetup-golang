package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var reqCnt = 1

func main() {
	log.Println("server is up at localhost:8080")

	http.HandleFunc("/order", placeOrder)
	http.ListenAndServe("localhost:8080", nil)
}

func placeOrder(res http.ResponseWriter, req *http.Request) {
	log.Println("==============================")
	log.Println("got req:", reqCnt)
	reqCnt = reqCnt + 1

	ctx := req.Context()

	select {
	case <-ctx.Done():
		log.Println("ctx.Done() is called")
	case <-time.After(5 * time.Second):
		log.Println("order ok")
		fmt.Fprintf(res, "order placed\n")
	}

}
