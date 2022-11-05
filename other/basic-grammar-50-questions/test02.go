package main

import "fmt"

//string类型可以更改吗
/*不能，尝试使用索引遍历字符串，来更新字符串中的个别字符，是不允许的。
string 类型的值是只读的二进制 byte slice，如果真要修改字符串中的字符，将 string 转为 []byte 修改后，再转为 string 即可。
*/
func main() {
	var name string
	name = "jamison"
	//name[0] = 'd'   错误

	nameBytes := []byte(name)
	nameBytes[0] = 'k'
	fmt.Println(string(nameBytes))
}
