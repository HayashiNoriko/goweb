package main

import (
	"fmt"
	"net/http"
	"os"
)

func main1() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("http server failed,", err)
		return
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 1. 直接写入字符串
	// fmt.Fprintln(w, "<h1>Hello World!</h1>")

	// 2.读取文件
	b, _ := os.ReadFile("./hello.html")
	fmt.Fprintln(w, string(b))
}
