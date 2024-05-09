package main

import (
	"fmt"
	"log"
	"net/http"
)


func main() {
	var port = 3000

	srv := http.Server{
		Addr: fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("handling")
			w.Write([]byte("hello"))
		}),
	}

	log.Println("Listening on ", port)
	srv.ListenAndServe()

	
}