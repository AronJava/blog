package main

import(
	"net/http"
)

func main() {
	http.HandleFunc("/", index)

	http.ListenAndServe(":11000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}
