package main

import (
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "http.ResponseWriter sample\n")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

// $ curl -Lv 'localhost:8080'
// * Rebuilt URL to: localhost:8080/
// *   Trying ::1...
// * TCP_NODELAY set
// * Connected to localhost (::1) port 8080 (#0)
// > GET / HTTP/1.1
// > Host: localhost:8080
// > User-Agent: curl/7.54.0
// > Accept: */*
// >
// < HTTP/1.1 200 OK
// < Date: Sat, 17 Feb 2018 07:13:18 GMT
// < Content-Length: 27
// < Content-Type: text/plain; charset=utf-8
// <
// http.ResponseWriter sample
// * Connection #0 to host localhost left intact
