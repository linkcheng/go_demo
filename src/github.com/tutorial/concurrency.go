package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func test() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func serve(addr string, handler http.Handler, stop <-chan struct{}) error {
	s := http.Server{
		Addr: addr, 
		Handler: handler,
	}
	go func() {
		<-stop
		s.Shutdown(context.Background())
	}()

	return s.ListenAndServe()
}

func serveApp(stop <-chan struct{}) error {
	addr := ":8080"
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})
	return serve(addr, mux, stop)
}

func serveDebug(stop <-chan struct{}) error {
	addr := ":8081"
	return serve(addr, http.DefaultServeMux, stop)
}

func main() {
	done := make(chan error, 2)
	stop := make(chan struct{})
	
	go func() {
		done <- serveApp(stop)
	}()

	go func() {
		done <- serveDebug(stop)
	}()

	var stopped bool
	for i:=0; i<cap(done); i++ {
		if err := <-done; err != nil {
			fmt.Printf("error: %v", err)
		}
		if !stopped {
			stopped = true
			close(stop)
		}
	}
}
