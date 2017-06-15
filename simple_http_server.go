// server start: go run simple_http_server.go
// access to localhost:8080

package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`
		<html><head><title>Simple HTTP Server</title></head><body>Hello World!</body></html>
		`))
		if err != nil {
			log.Fatal(err)
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
