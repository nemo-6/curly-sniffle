package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/potato", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Method: "+r.Method+"\n")
		fmt.Fprintf(w, "URL: "+r.URL.Host+"\n")
		fmt.Fprintf(w, "Host: "+r.Host+"\n")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
