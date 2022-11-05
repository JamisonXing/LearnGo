package main

import (
	"fmt"
	"io"
	"net/http"
)

//如何关闭http的响应体
/*直接在处理http相应错误的代码块中*/

func main() {
	resp, err := http.Get("www.baidu.com")

	//正确关闭resp.Body
	if resp != nil {
		defer resp.Body.Close()
	}
	checkError(err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	checkError(err)
	fmt.Println(string(body))
}

func checkError(err error) {
	fmt.Println(err)
}
