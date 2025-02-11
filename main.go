package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
)

type Hello struct {
	Message string `json:"message"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(Hello{Message: "Hello, World!"})
	})
	mux.HandleFunc("GET /query", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Query().Encode()))
	})
	mux.HandleFunc("GET /path/{path}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.PathValue("path")))
	})

	println(`
                                             ▄▄  
  ▄▄█▀▀▀█▄█               ██                 ██  
▄██▀     ▀█              ▄██▄                    
██▀       ▀  ▄██▀██▄    ▄█▀██▄   ▀████████▄▀███  
██          ██▀   ▀██  ▄█  ▀██     ██   ▀██  ██  
██▄    ▀██████     ██  ████████    ██    ██  ██  
▀██▄     ██ ██▄   ▄██ █▀      ██   ██   ▄██  ██  
  ▀▀███████  ▀█████▀▄███▄   ▄████▄ ██████▀ ▄████▄
                                   ██            
                                 ▄████▄          
    `)
	log.Default().Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
	print("Hello, World!")
}
