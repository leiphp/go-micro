package main

import (
	"go-micro/src/gateway"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func main() {
	routes := gateway.InitConfig()
	//fmt.Println(routes[0],routes[1])
	http.HandleFunc("/",func(writer http.ResponseWriter, request *http.Request) {
		if r := routes.Match(request);r!=nil {
			remote, _ := url.Parse(r.Url)
			//request.URL.Path /v1/detail/1 => /detail/1去掉v1
			request.URL.Path = strings.Replace(request.URL.Path,"/v1","",-1)
			p := httputil.NewSingleHostReverseProxy(remote)
			p.ServeHTTP(writer, request)
		}else{
			writer.WriteHeader(http.StatusNotFound)
		}

	})
	http.ListenAndServe(":8123",nil)
}