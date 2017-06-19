// server start: go run simple_http_server.go
// access to localhost:8080

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	host := flag.String("b", "127.0.0.1", "Binds to the specified IP.")
	port := flag.String("p", "3000", "Runs on the specified port.")
	flag.Parse()

	appRoot := fmt.Sprintf("%v:%v", *host, *port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`
		<html><head><title>Simple HTTP Server</title></head><body>Hello World!</body></html>
		`))
		if err != nil {
			log.Fatal(err)
		}
	})

	log.Printf("application starting on %v\n", appRoot)

	if err := http.ListenAndServe(appRoot, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
