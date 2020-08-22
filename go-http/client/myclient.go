package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("*** HTTP Client Demo ***")

	var req *http.Request
	var res *http.Response
	var err error

	tranport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}
	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: tranport,
	}

	// home
	printAction("Home")
	res, err = client.Get("http://localhost:8080")
	// client := http.DefaultClient
	// res, err = client.Get("http://localhost:8080")
	checkErr(err, "Get failed")
	processResponse(res)
	// ---

	// view item
	printAction("View Items")
	req, err = http.NewRequest(http.MethodGet, "http://localhost:8080/items?category=tech", nil)
	checkErr(err, "New items request failed")
	res, err = client.Do(req)
	checkErr(err, "Get items failed")
	processResponse(res)
	// ---

	// login
	printAction("Login")
	req, err = http.NewRequest(http.MethodGet, "http://localhost:8080/login", nil)
	checkErr(err, "New login request failed")
	req.SetBasicAuth("nurali", "abcd1234")

	res, err = client.Do(req)
	checkErr(err, "Login failed")
	cookie := res.Cookies()[0]
	fmt.Println("cookie=", cookie)
	// ---

	// buy item-1
	printAction("Buy Item")
	req, err = http.NewRequest(http.MethodPost, "http://localhost:8080/buy", strings.NewReader(`{"ID":"I001", "Quantity":2}`))
	req.AddCookie(cookie)
	checkErr(err, "New buy request failed")

	res, err = client.Do(req)
	checkErr(err, "Buy item failed")
	fmt.Println("OK")
	// ---

	// buy item-2
	printAction("Buy Item")
	req, err = http.NewRequest(http.MethodPost, "http://localhost:8080/buy", strings.NewReader(`{"ID":"I002", "Quantity":1}`))
	req.AddCookie(cookie)
	checkErr(err, "New buy request failed")

	_, err = client.Do(req)
	checkErr(err, "Buy item failed")
	fmt.Println("OK")
	// ---
}

func printAction(action string) {
	fmt.Println()
	fmt.Println(">>", action)
}
func checkErr(err error, msg string) {
	if err != nil {
		fmt.Printf("[CLIENT] %s, err:%v\n", msg, err)
		os.Exit(1)
	}
}

func processResponse(res *http.Response) {
	checkCode(res, "")
	body, err := ioutil.ReadAll(res.Body)
	checkErr(err, "Read body failed")
	fmt.Println(string(body))
	fmt.Println()
}

func checkCode(res *http.Response, msg string) {
	if res.StatusCode != 200 {
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Printf("[CLIENT] %s, code:%d, status:%s, err:%s\n", msg, res.StatusCode, res.Status, string(body))
		os.Exit(1)
	}
}
