package main

import (
	"fmt"
	"io"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	fmt.Println("hello world")

	http.HandleFunc("/healthz", Headlthz)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("listen failure")
	}

}

func Headlthz(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Headlthz", r.Header, r.Body)
	fmt.Println("Remote addr", r.RemoteAddr)

	for k, vs := range r.Header {
		for _, v := range vs {
			w.Header().Set(k, v)
		}
	}
	w.Header().Set("VERSION", "v1")
	
	w.WriteHeader(404)
	io.WriteString(w, "OK")
}
