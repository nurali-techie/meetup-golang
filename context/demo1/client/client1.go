package main

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second) // run1
	// ctx, cancel := context.WithTimeout(ctx, 7*time.Second)	// run2
	defer cancel()

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
