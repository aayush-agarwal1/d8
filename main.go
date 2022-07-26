package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		dummyUser := user{"aayush", 21}
		err := json.NewEncoder(w).Encode(dummyUser)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	} else if r.Method == "POST" {
		var dummyUser user
		err := json.NewDecoder(r.Body).Decode((&dummyUser))
		fmt.Println(dummyUser)
		if err != nil {
			http.Error(w, "Incorrect User struct format", http.StatusBadRequest)
		}
	} else {
		http.Error(w, "Unsupported Method", http.StatusBadRequest)
	}
}
func main() {
	http.HandleFunc("/", myHandler)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
