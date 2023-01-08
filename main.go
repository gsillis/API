package main

import "net/http"

func main() {
	http.ListenAndServe(":3000", Handler{})
}

type Handler struct{}

func (Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
