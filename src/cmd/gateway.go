package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main(){
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		remote,_ := url.Parse("http://"+"localhost:9000")//反向代理9000端口的内网http服务
		p := httputil.NewSingleHostReverseProxy(remote)
		p.ServeHTTP(writer,request)
	})
	http.ListenAndServe(":8090",nil)
}
