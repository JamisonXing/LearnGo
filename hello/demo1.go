package main

func main() {
	/*a := 100
	b := 200
	a, b = b, a
	fmt.Printf("a = %d, b = %d", a, b)*/

	/*//类型转换
	c := 4.0
	d := int(c)
	fmt.Print(d, c)*/

	/*	//修改字符串
		a := "xxj"
		//var b []byte = []byte(a)
		b := []byte(a)
		b[0] = 'k'
		fmt.Print(b)
		fmt.Print(string(b))*/

	/*	//转义字符
		str := "he\n\toll"
		fmt.Println(str)
		str1 := "\u4e16"
		fmt.Println(str1)

		str2 := `hel\nlo`	//反引号原样输出
		fmt.Println(str2) */

	/*	//替换字符串
		str1 := "hello,xxj"
		target := "lili"
		char1 := []byte(str1)*/

	/*	//多个字符组成字符串
		str1 := "hello"
		char1 := [5]byte{104, 101, 108, 108, 111}
		fmt.Printf("str1:%s\n", str1)
		fmt.Printf("char1:%s", char1)*/

	/*	//字母和汉字字节长度
		str1 := "abc"
		str2 := "你好"
		fmt.Println(len(str1)) //3
		fmt.Println(len(str2)) //6*/

	/*		//字符串实际上就是字节数组
			str1 := "xxj"
			fmt.Print(str1[0])
			//结果：120*/

	/*	//获取字符串中某个字节的地址属于非法行为
		str := "xxj"
		fmt.Print(&str[2])*/

	/*	//计算字符串长度
		str1 := "在吗"
		fmt.Println(len(str1))
		fmt.Println(utf8.RuneCountInString(str1))
		//结果：6	2*/

	/*	//拼接字符串"+"
		str := "xxj" + "你好"
		str += "吗"
		fmt.Println(str)
		//拼接大量的字符串时候
		str1 := "abc,,,,,"
		str2 := "def,,,,,"
		var stringBuilder bytes.Buffer
		stringBuilder.WriteString(str1)
		stringBuilder.WriteString(str2)
		fmt.Printf("拼接之后的字符串：%s", stringBuilder.String())*/

	/*	//从字符串”xxj,你好“中获取"好"
		str1 := "abc" //对应ASCII码表，底层为[]byte数组
		fmt.Printf("%c\n", str1[2])
		str2 := "xxj,你好" //有中文，对应unicode字符集，会变成[]rune数组
		//fmt.Printf("%c")   //直接使用索引无效，可以使用string方法进行转换
		fmt.Print(string([]rune(str2)[5]))
		//结果:c 好*/

	//遍历

}
