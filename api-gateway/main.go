package main

import (
	"github.com/brandonbyrd6/api-gateway/reverseproxy"
	"github.com/gorilla/mux"
)

func main() {

	proxy := &reverseproxy.ReverseProxy{}
	r := mux.NewRouter()
	r.Host("localhost").PathPrefix("/api")
	proxy.AddTarget("http://localhost:8081", r)

	proxy.AddTarget("http://localhost:8001", nil)

	proxy.Start()
}
