package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	addr := os.Getenv("ADDR")

	if addr == "" {
		addr = ":8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request coming in on `/`")
		w.Write([]byte(`{"message":"root"}`))
	})
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request coming in on `/hello`")
		w.Write([]byte(`{"message":"hello world"}`))
	})

	mux.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request coming in on `/pod`")
		resp, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		fmt.Println(string(resp))
		w.Write([]byte(`{"message":"successful"}`))
	})

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	log.Printf("Starting server on %s", addr)
	log.Println(server.ListenAndServe())
}
