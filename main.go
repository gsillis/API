package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/", index)
	router.GET("/:name", index)

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if p.ByName("name") != "" {
		w.Write([]byte("Hello" + p.ByName("name")))
		return
	}
	w.Write([]byte("Hello World"))
}
