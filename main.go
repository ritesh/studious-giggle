package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":5050"
	} else {
		port = ":" + port
	}
	log.Printf("%s", port)
	http.HandleFunc("/", sayHello)
	log.Fatal(http.ListenAndServe(port, nil))
}

const (
	msgHello = `<?xml version="1.0" encoding="UTF-8"?>
<Response> 
<Say> What's the password? </Say>
</Response> `
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/xml")
	w.Write([]byte(msgHello))
}
