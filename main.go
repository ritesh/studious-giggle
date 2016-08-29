package main

import "net/http"

func main() {
	http.HandleFunc("/", sayHello)
	http.ListenAndServe(":5050", nil)
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
