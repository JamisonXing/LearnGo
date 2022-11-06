package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//解析 JSON 数据时，默认将数字当做哪种类型
/*
在 encode/decode JSON 数据时，Go 默认会将数值当做 float64 处理。
*/

func main() {
	var data = []byte(`{"status": 200}`)
	var result map[string]any

	if err := json.Unmarshal(data, &result); err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("200的数据类型为：%T\n", result["status"])
}

/*
result:
200的数据类型为：float64
*/
