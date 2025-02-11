package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/mojocn/base64Captcha"
)

type Hello struct {
	Message string `json:"message"`
}

func genCaptcha64(w http.ResponseWriter, r *http.Request) {
	c := base64Captcha.NewCaptcha(base64Captcha.DefaultDriverDigit, base64Captcha.DefaultMemStore)
	id, b64s, answer, err := c.Generate()
	fmt.Printf("id: %s answer: %s", id, answer)
	body := map[string]interface{}{"code": 1, "data": b64s, "captchaId": id, "msg": "success"}
	if err != nil {
		body = map[string]interface{}{"code": 0, "msg": err.Error()}
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(body)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(Hello{Message: "Hello, World!"})
	})
	mux.HandleFunc("GET /get-captcha", genCaptcha64)
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
