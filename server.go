package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	d := []byte{83, 119, 118, 112, 56, 98, 118, 117, 99, 65, 51, 50}
	fmt.Fprintf(w, "%s", d)
}
