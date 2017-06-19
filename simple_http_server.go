// server start: go run simple_http_server.go
// access to localhost:8080

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func newHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`
		<html><head><title>Simple HTTP Server</title></head><body>Hello World!</body></html>
		`))
		if err != nil {
			log.Println(err)
		}
	})
	return mux
}

func main() {
	host := flag.String("b", "127.0.0.1", "Binds to the specified IP.")
	port := flag.String("p", "3000", "Runs on the specified port.")
	flag.Parse()

	address := fmt.Sprintf("%v:%v", *host, *port)
	server := &http.Server{
		Handler: newHandler(),
		Addr:    address,
	}

	log.Printf("application starting on %v\n", address)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()

	sigCh := make(chan os.Signal, 1)
	// signal.Notify(c chan<- os.Signal, sig ...os.Signal)
	// キャッチするシグナルをチャネルに登録する
	signal.Notify(sigCh,
		syscall.SIGTERM,
	)
	<-sigCh
	log.Printf("application stopped on %v\n", address)

	// シグナルを受け取ったらシャットダウンする
	// 処理の終了を最大5秒まで待つ
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
