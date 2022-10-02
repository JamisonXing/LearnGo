package main

type user struct {
	Name string
	Age  int
}

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

	/*	//遍历
		str := "niHaoMa"        //ascii
		str1 := "你好吗"           //unicode 一般为中文只能用for range才能读取到正确字符
		for k, s := range str { //for range
			fmt.Printf("unicode:%c %d,%d\n", s, s, k)
		}
		fmt.Print("__________")

		for i := 0; i < len(str); i++ {
			fmt.Printf("ascii:%c %d\n", str[i], str[i])
		}

		fmt.Print("__________")
		for _, m := range str1 {
			fmt.Printf("ascii:%c %d\n", m, m)
		}*/

	/*	//字符串的格式化
		str := "你好，xxj"
		result := fmt.Sprint(string([]rune(str)[2]))
		fmt.Println(result)
		//结果：，
		fmt.Printf("动态类型：%T\n", str) //动态类型：%T 输出结果:动态类型：string

		userInfo := user{
			Name: "xxj",
			Age:  20,
		}
		fmt.Printf("输出value:%v\n", userInfo)      //结果：输出value:{xxj 20}
		fmt.Printf("输出key-value:%+v\n", userInfo) //结果：输出key-value:{Name:xxj Age:20}

		fmt.Printf("userInfo的十六进制指针：%p\n", &userInfo) //结果：userInfo的十六进制指针：0xc000008078
		fmt.Printf("6的二进制表示：%b\n", 6) */ //6的二进制表示：110

	/*	//字符串的查找 正向与反向
		str := "我来了,我走了,bye bye"
		//正向
		index := strings.Index(str, "bye")
		fmt.Printf("bye 所在的位置：%d\n", index) //结果：20

		index1 := strings.Index(str, "你好")
		fmt.Printf("你好 所在的位置：%d\n", index1) //结果：-1

		index2 := strings.Index(str[index:], "e")
		fmt.Printf("e 所在的位置：%d\n", index2) //结果：2

		//反向类似
		index3 := strings.LastIndex(str, "by")
		fmt.Printf("by 所在的位置：%d\n", index3)*/

	//常量和iota修改器
	/*	const name string = "xxj"
		const pi = 2.3 //type可以省略

		const ( //与变量定义一样可以批量声明多个常量
			e = 232
			p = 888
		)

		var ( //变量批量声明
			name_ string
			age   int
		)*/

	/*	const (
			a = 1
			b
			c = 3
			d
		)
		fmt.Println(a, b, c, d) //1,1,3,3*/

	/*	const (
			a = 4
			b
			c = iota
			d
		)
		fmt.Print(a, b, c, d) //4，4，2，3*/

	/*	//指针
		var room int = 10
		var ptr = &room

		fmt.Printf("变量地址：%p\n", &room)
		fmt.Printf("指针变量地址：%p\n", &ptr)
		fmt.Printf("指针指向的值：%d\n", *ptr)
		fmt.Printf("变量地址：%p", ptr)
		变量地址：0xc0000a6058
		指针变量地址：0xc0000ca018
		指针指向的值：10
		变量地址：0xc0000a6058*/

	/*	//变量逃逸test
		f()
		fmt.Print(*global) //仍然能够取出x的值*/

	/*	//类型别名
		type myInt = int
		var age myInt = 22
		var height int = 43
		fmt.Printf("age的数据类型：%T\n", age)
		fmt.Printf("height的数据类型：%T\n", height)
			//结果：
			//age的数据类型：int
			//height的数据类型：int*/

	/*	//字符串和其他数据类型的转换
		//1.字符串 《--》整数
		str := "45"
		intValue, _ := strconv.Atoi(str)
		fmt.Printf("str转为整数后：%d\n", intValue)
		//结果：str转为整数后：45
		intValue1 := 2
		strValue := strconv.Itoa(intValue1)
		fmt.Printf("整数转为字符串后：%s, 类型为：%T\n", strValue, strValue)
		//结果：整数转为字符串后：2, 类型为：string

		//2. 字符串 《--》 浮点
		str1 := "3.14"
		parseFloat, _ := strconv.ParseFloat(str1, 32)
		fmt.Printf("字符串转浮点：%f,类型：%T\n", parseFloat, parseFloat)
		//结果：字符串转浮点：3.140000,类型：float64
		floatValue := 3.14
		str2 := strconv.FormatFloat(floatValue, 'f', 2, 32)
		fmt.Printf("浮点数转字符串：%s,类型：%T\n", str2, str2)
		//结果：浮点数转字符串：3.14,类型：string*/
}

/*//变量逃逸
var global *int

func f() {
	x := 1
	global = &x
}*/
