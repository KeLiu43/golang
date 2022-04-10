package main

import (
	"io"
	"log"
	"net/http"
)

//健康检查函数
func healthz(w http.ResponseWriter, r *http.Request) {
	//Fprintf：来格式化并输出到 io.Writers 而不是 os.Stdout。
	io.WriteString(w, "200")
}

//业务处理
func bussinesshandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("VERSION", "HTTP 2.0")
	for hk, _ := range r.Header {
		w.Header().Set(hk, "hv")
	}
	log.Printf("Response code: %d", 200)
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/healthz", healthz)
	server.HandleFunc("/", bussinesshandle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("start server failed", err)
	}

	// fmt.Println("hello world")
}
