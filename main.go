package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fileOutput, err := os.ReadFile("buffer.txt")
		if err != nil {
			http.Error(w, "os.ReadFile() failed with "+string(err.Error()), http.StatusInternalServerError)
			return
		}
		var dummyUser user
		err = json.Unmarshal(fileOutput, &dummyUser)
		if err != nil {
			http.Error(w, "Incorrect Format Information in buffer err:"+string(err.Error()), http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(dummyUser)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	} else if r.Method == "POST" {
		var dummyUser user
		err := json.NewDecoder(r.Body).Decode((&dummyUser))
		if err != nil {
			http.Error(w, "Incorrect User struct format", http.StatusBadRequest)
		}
		marshalledUser, err := json.Marshal(dummyUser)
		if err != nil {
			http.Error(w, "Could not Marshal data err:"+string(err.Error()), http.StatusInternalServerError)
			return
		}
		err = os.WriteFile("buffer.txt", marshalledUser, 0777)
		if err != nil {
			http.Error(w, "Could not Write to file err:"+string(err.Error()), http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Unsupported Method", http.StatusBadRequest)
	}
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
	http.HandleFunc("/file", fileHandler)
	http.HandleFunc("/", myHandler)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
