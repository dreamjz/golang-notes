package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 注册路由
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)

	// 监听并启动 HTTP 服务
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.path = %q\n", req.URL.Path)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	// 输出所有 http 请求头 request.Header
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] =  %q\n", k, v)
	}
}
