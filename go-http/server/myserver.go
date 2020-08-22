package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// Item struct
type Item struct {
	ID    string
	Name  string
	Price string
}

var items = []*Item{
	{ID: "I001", Name: "Mobile", Price: "10000"},
	{ID: "I002", Name: "Laptop", Price: "40000"},
}

// Order struct
type Order struct {
	ID       string
	Quantity int
}

func main() {
	fmt.Println("*** HTTP Server Demo ***")

	// http.HandleFunc("/", homeHandler)
	// http.HandleFunc("/items", itemsHandler)
	router := http.NewServeMux()
	router.HandleFunc("/", homeHandler)
	router.Handle("/login", http.HandlerFunc(loginHandler))
	router.HandleFunc("/items", itemsHandler)
	router.Handle("/buy", http.HandlerFunc(buyHandler))

	// http.ListenAndServe("0.0.0.0:8080", router)
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: router,
	}
	server.ListenAndServe()
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	printAction("homeHandler()")
	// time.Sleep(time.Second * 5)
	res := `-> Flipkart
		-> Login
		-> View Items
		-> Buy Items`
	w.Write([]byte(res))
	fmt.Println("return menu")
}

func itemsHandler(w http.ResponseWriter, r *http.Request) {
	printAction("itemsHandler()")

	category := r.URL.Query().Get("category")
	fmt.Println("category=", category)
	if category == "" {
		http.Error(w, "ERR-001: missing category", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(items)
	checkErr(err, "Failed to marshal items")
	w.Write(res)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	printAction("loginHandler()")

	auth := r.Header.Get("Authorization")
	fmt.Println("auth=", auth)

	cookieValue := fmt.Sprintf("S-%d", time.Now().Nanosecond())
	fmt.Println("set-cookie=", cookieValue)
	cookie := &http.Cookie{
		Name:    "session-id",
		Value:   cookieValue,
		Expires: time.Now().Add(time.Minute * 10),
	}
	http.SetCookie(w, cookie)
}

func buyHandler(w http.ResponseWriter, r *http.Request) {
	printAction("buyHandler()")

	cookie, err := r.Cookie("session-id")
	checkErr(err, "Cookie not found")
	fmt.Println("get-cookie=", cookie.Value)

	body, err := ioutil.ReadAll(r.Body)
	checkErr(err, "Failed to read payload")

	order := &Order{}
	err = json.Unmarshal(body, order)
	checkErr(err, "Failed to unmarshal order")

	fmt.Println("order =>", order.ID, order.Quantity)
}

func checkErr(err error, msg string) {
	if err != nil {
		fmt.Printf("[SERVER] %s, err:%v\n", msg, err)
		os.Exit(1)
	}
}

func printAction(action string) {
	fmt.Println()
	fmt.Println(">>", action)
}
