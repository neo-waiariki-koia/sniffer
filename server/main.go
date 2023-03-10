package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
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

	mux.HandleFunc("/sleep", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Minute)
		w.Write([]byte(`{"message":"finished sleeping"}`))
	})

	mux.HandleFunc("/post_success", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request coming in on `/post_success`")
		resp, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		fmt.Println(string(resp))
		w.Write([]byte(`{"message":"successful"}`))
	})

	mux.HandleFunc("/post_fail", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request coming in on `/post_fail`")
		resp, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		fmt.Println(string(resp))
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"message":"failed"}`))
	})

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	log.Printf("Starting server on %s", addr)
	log.Println(server.ListenAndServe())
}
