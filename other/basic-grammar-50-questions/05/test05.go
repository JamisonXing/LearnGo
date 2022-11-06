package main

import (
	"fmt"
	"io"
	"net/http"
	test04 "other/basic-grammar-50-questions/04"
)

//你是否主动关闭过http连接，为啥要这样做
/*
有关闭，不关闭会程序可能会消耗完 socket 描述符。有如下2种关闭方式：
1.直接设置请求变量的 Close 字段值为 true，每次请求结束后就会主动关闭连接。
设置 Header 请求头部选项 Connection: close，然后服务器返回的响应头部也
会有这个选项，此时 HTTP 标准库会主动断开连接

2.你可以创建一个自定义配置的 HTTP transport 客户端，用来取消 HTTP 全局
的复用连接。
*/

func main() {
	req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	test04.CheckError(err)

	req.Close = true
	//req.Header.Add("Connection", "close") //等效的关闭方式

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		defer resp.Body.Close()
	}
	test04.CheckError(err)

	body, err := io.ReadAll(resp.Body)
	test04.CheckError(err)

	fmt.Println(string(body))
}
