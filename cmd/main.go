package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
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
