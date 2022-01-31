package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// set http port
	httpPort := 3000

	// handle URL
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s /\n", r.Method)
		rw.WriteHeader(200)
		rw.Header().Set("Content-Type", "text/html; charset=utf-8")
		rw.Write([]byte("Hello world, Bagus..."))
	})

	// serve
	fmt.Printf("serve http on port %d\n", httpPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", httpPort), nil); err != nil {
		log.Fatal(err)
	}
}
