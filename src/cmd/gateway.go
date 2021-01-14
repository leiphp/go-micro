package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	http.HandleFunc("/",func(writer http.ResponseWriter, request *http.Request) {
		remote, _ := url.Parse("http://"+"localhost:9000")
		p := httputil.NewSingleHostReverseProxy(remote)
		p.ServeHTTP(writer, request)
	})
	http.ListenAndServe(":8123",nil)
}