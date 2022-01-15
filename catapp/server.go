package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("Please provide a port as a command line argument")
	}
	port := args[0]
	mux := http.NewServeMux()
	mux.HandleFunc("/", cat)
	mux.HandleFunc("/favicon.ico", errorFavi)
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal("ListenAndServe failed with error: ", err)
	}
}

type catJSON []struct {
	Breeds []interface{} `json:"breeds"`
	ID     string        `json:"id"`
	URL    string        `json:"url"`
	Width  int           `json:"width"`
	Height int           `json:"height"`
}

func cat(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("index.gohtml")
		if err != nil {
			log.Fatal("Parse failed: ", err)
		}
		response, err := http.Get("https://api.thecatapi.com/v1/images/search")
		jsonObj := catJSON{}
		json.NewDecoder(response.Body).Decode(&jsonObj)
		println(jsonObj[0].URL)
		if err != nil {
			log.Fatal(err)
		}
		if err := t.ExecuteTemplate(w, "index", jsonObj[0]); err != nil {
			log.Fatal("ExecuteTemplate failed:", err)
		}
	}
}
func errorFavi(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
