package main

import (
	"fmt"
	"go-micro/src/gateway"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

//自建http网关服务
func main(){
	routes := gateway.InitConfig()
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if r := routes.Match(request);r != nil{
			remote,_ := url.Parse(r.Url)//反向代理9000端口的内网http服务
			//访问http://localhost:8090/v1/test代理访问=>http://localhost:9000/v1/test
			//最后去掉v1，成功访问http://localhost:9000/test
			request.URL.Path = strings.Replace(request.URL.Path,"/v1","",-1)
			p := httputil.NewSingleHostReverseProxy(remote)
			p.ServeHTTP(writer,request)
		}else{
			writer.WriteHeader(http.StatusNotFound)
		}

	})
	fmt.Println("服务启动成功")
	http.ListenAndServe(":8090",nil)
}
