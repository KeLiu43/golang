package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

//健康检查函数
func healthz(w http.ResponseWriter, r *http.Request) {
	//Fprintf：来格式化并输出到 io.Writers 而不是 os.Stdout。
	fmt.Fprintf(w, "working")
}

//业务处理
func index(w http.ResponseWriter, r *http.Request) {
	// 03.设置version
	os.Setenv("VERSION", "v0.0.1")
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)
	fmt.Printf("os version: %s \n", version)
	for k, v := range r.Header {
		for _, vv := range v {
			fmt.Printf("Header key: %s, Header value: %s \n", k, v)
			w.Header().Set(k, vv)
		}
	}

	// 04.记录日志并输出
	clientip := getCurrentIP(r)
	log.Printf("Success! Response code: %d", 200)
	log.Printf("Success! clientip: %s", clientip)
}

func getCurrentIP(r *http.Request) string {
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		// 当请求头不存在即不存在代理时直接获取ip
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	return ip
}

func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
		return ip
	}
	return ""
}

func main() {
	fmt.Println("exec main functions")
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/healthz", healthz)
	if err := http.ListenAndServe(":8088", mux); err != nil {
		log.Fatalf("start http server failed, error: %s\n", err.Error())
	} else {
		log.Default().Println("visit OK")
	}

	// fmt.Println("hello world")
}
