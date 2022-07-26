package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type user struct {
	Name string
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		dummyUser := user{"aayush"}
		json.NewEncoder(w).Encode(dummyUser)
	}
}
func main() {
	fmt.Printf("Hello World 2")
	http.HandleFunc("/", myHandler)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
