/**
 *
 * @package       main
 * @author        YuanZhiGang <zackyuan@yeah.net>
 * @version       1.0.0
 * @copyright (c) 2013-2023, YuanZhiGang
 */

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

//自定义Server示例
//func httpResponseUserLogin(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Hello, World!!!!!") // 响应 "Hello, World!"
//}
//
//func main() {
//
//	server := &http.Server{
//		Addr:    ":8080",
//		Handler: http.HandlerFunc(httpResponseUserLogin),
//	}
//
//	err := server.ListenAndServe()
//	if err != nil {
//		fmt.Println("server began failed,err:", err)
//		return
//	}
//
//}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	// 直接输出字符串
	//str := "hfa;gfhsd;lfgjadp;grj"
	//_, err := w.Write([]byte(str))

	// 从文件中读取，内容可以是html,css,js
	data, err := os.ReadFile("./github-recovery-codes.txt")
	if err != nil {
		fmt.Println("#1response failed,err:", err)
	}

	_, err = w.Write(data)
	if err != nil {
		fmt.Println("#2response failed,err:", err)
	}

}

func handlerFuncClient(w http.ResponseWriter, r *http.Request) {
	// client 请求响应
	defer r.Body.Close()

	fmt.Println(r.URL.Query().Get("name")) // 获取url query 的 param
	fmt.Println(r.URL.Query().Get("age"))
	fmt.Println(r.URL)
	fmt.Println(r.Method)
	fmt.Println(io.ReadAll(r.Body))

	str := "返回给client的数据:请求收到了，ok!"
	_, err := w.Write([]byte(str))
	if err != nil {
		fmt.Println("#1response failed,err:", err)
	}
}

func handlerFuncPostClient(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// 1. 请求类型是application/x-www-form-urlencoded时解析form数据
	r.ParseForm()
	fmt.Println(r.PostForm) // 打印form数据
	fmt.Println(r.PostForm.Get("name"), r.PostForm.Get("age"))
	// 2. 请求类型是application/json时从r.Body读取数据
	data, _ := io.ReadAll(r.Body)
	fmt.Println(string(data))

	answer := `{"status": "ok"}`
	w.Write([]byte(answer))
}

func main() {
	http.HandleFunc("/cs", handlerFunc)
	http.HandleFunc("/client_cs", handlerFuncClient)
	http.HandleFunc("/client_cs_post", handlerFuncPostClient)
	err := http.ListenAndServe("127.0.0.1:9000", nil)
	if err != nil {
		fmt.Println("http failed,err:", err)
		return
	}
}
