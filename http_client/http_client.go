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
	"strings"
)

// 1.响应数据读取完毕之后需要关闭Response.Body.close()
// 2. keep-alive 设置

// http（GET）基础请求...
//func main() {
//	resp, err := http.Get("http://127.0.0.1:9000/client_cs?name=zackyuan&age=35")
//	if err != nil {
//		fmt.Println("request failed,err:", err)
//	}
//
//	// http://127.0.0.1:9000/client_cs?name=zackyuan&age=35
//	defer resp.Body.Close()
//
//	data, err := io.ReadAll(resp.Body)
//	if err != nil {
//		fmt.Println("response data read failed,err:", err)
//	}
//
//	fmt.Println("响应数据:", string(data))
//
//}

// 带参数的GET请求示例
//func main() {
//	urlObj, err := url.Parse("http://127.0.0.1:9000/client_cs")
//	if err != nil {
//		fmt.Println("parse  requestUrl failed,err:", err)
//	}
//
//	data := url.Values{}
//	data.Set("name", "===zackyuan???!@#@$#$")
//	data.Set("age", "35")
//
//	strObj := data.Encode()
//	urlObj.RawQuery = strObj
//	fmt.Println(urlObj.String())
//
//	// 请求方式1
//	//http.Get()
//
//	// 请求方式2
//	req, err := http.NewRequest("GET", urlObj.String(), nil) // 获取一个请求客户端对象
//	if err != nil {
//		fmt.Println("get http request failed,err:", err)
//		return
//	}
//	//http.DefaultClient.Do(req) // 发送请求
//
//	// 设置禁用长连接
//	client := http.Client{
//		Transport: &http.Transport{
//			DisableKeepAlives: true,
//		},
//	}
//
//	resp, err := client.Do(req) // 发送请求
//	if err != nil {
//		fmt.Println("http response failed,err:", err)
//		return
//	}
//
//	defer resp.Body.Close()
//
//	// 读取响应数据
//	resp_data, err := io.ReadAll(resp.Body)
//	if err != nil {
//		fmt.Printf("get resp failed, err:%v\n", err)
//		return
//	}
//
//	fmt.Println(string(resp_data))
//
//}

// Post请求示例1
//func main() {
//	urlObj, err := url.Parse("http://127.0.0.1:9000/client_cs_post")
//	if err != nil {
//		fmt.Println("parse  requestUrl failed,err:", err)
//	}
//
//	data := `{"name":"小王子","age":18}`
//	//contentType := "application/json"
//	//http.Post(urlObj.String(), contentType, strings.NewReader(data))
//
//	// 请求方式2
//	req, err := http.NewRequest("POST", urlObj.String(), strings.NewReader(data)) // 获取一个请求客户端对象
//	if err != nil {
//		fmt.Println("post http request failed,err:", err)
//		return
//	}
//
//	req.Header.Set("Content-Type", "application/json")
//
//	//http.DefaultClient.Do(req) // 发送请求
//
//	// 设置禁用长连接
//	client := http.Client{
//		Transport: &http.Transport{
//			DisableKeepAlives: true,
//		},
//	}
//
//	resp, err := client.Do(req) // 发送请求
//	if err != nil {
//		fmt.Println("http response failed,err:", err)
//		return
//	}
//
//	defer resp.Body.Close()
//
//	// 读取响应数据
//	resp_data, err := io.ReadAll(resp.Body)
//	if err != nil {
//		fmt.Printf("get resp failed, err:%v\n", err)
//		return
//	}
//
//	fmt.Println(string(resp_data))
//
//}

// Post请求示例2
func main() {
	url := "http://127.0.0.1:9000/client_cs_post"
	// 表单数据
	contentType := "application/x-www-form-urlencoded"
	data := "name=小王子&age=18"
	// json
	//contentType := "application/json"
	//data := `{"name":"小王子","age":18}`
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}
