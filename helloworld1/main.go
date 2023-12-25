package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!222")
	remoteAddr := r.RemoteAddr
	fmt.Fprintf(w, "Remote Address: %s", remoteAddr)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}
