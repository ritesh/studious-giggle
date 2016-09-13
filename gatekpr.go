package main

import "net/http"

func init() {
	http.HandleFunc("/", sayHello)
}

const (
	msgHello = `<?xml version="1.0" encoding="UTF-8"?>
				<Response> 
				<Say>Something something</Say>
				</Response> `
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/xml")
	w.Write([]byte(msgHello))
}
