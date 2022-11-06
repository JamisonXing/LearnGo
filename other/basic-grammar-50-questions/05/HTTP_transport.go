package main

import (
	"fmt"
	"io"
	"net/http"
	test04 "other/basic-grammar-50-questions/04"
)

// 你可以创建一个自定义配置的 HTTP transport 客户端，用来取消 HTTP 全局的复用连接
func main() {
	tr := http.Transport{
		DisableKeepAlives: true,
	}
	client := http.Client{Transport: &tr}

	resp, err := client.Get("http://www.baidu.com")
	if resp != nil {
		defer resp.Body.Close()
	}
	test04.CheckError(err)

	fmt.Println(resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	test04.CheckError(err)

	fmt.Println(string(body))

}
