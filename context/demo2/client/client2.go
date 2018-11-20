package main

import (
	"bufio"
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	test1()
}

func test1() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		log.Println("press 'enter' to cancel client request before server response arrives")
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		cancel()
	}()

	payload := `{"user":"hemal", "item": "thinkpad", "otp":"220011"}`
	req, _ := http.NewRequest("POST", "http://localhost:8080/order", bytes.NewBuffer([]byte(payload)))
	req = req.WithContext(ctx)

	log.Println("placing order")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	resBody, _ := ioutil.ReadAll(res.Body)
	log.Printf("response: %s", string(resBody))
}
