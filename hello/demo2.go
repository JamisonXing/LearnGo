package main

import "fmt"

func main() {
	/*	//获取键盘输入字符
		reader := bufio.NewReader(os.Stdin)
		readString, _ := reader.ReadString('\n')
		fmt.Println(readString)*/

	/*	//数组
		scores := [5]int{1, 2, 3, 4, 5}
		for index, value := range scores {
			fmt.Printf("索引：%d, 得分：%d\n", index, value)
		}

		scores1 := [3]int{2: 3} //只初始化第三个值，且设置为3
		fmt.Print(scores1)*/

	/*	//多维数组
		arr := [4][2]int{1: {12, 21}, 2: {43, 34}}
		for index, value := range arr {
			fmt.Printf("索引：%d,值：%d\n", index, value)
		}
		//结果：
		//索引：0,值：[0 0]
		//索引：1,值：[12 21]
		//索引：2,值：[43 34]
		//索引：3,值：[0 0]

		//对多维数组元素的简单应用
		var arr1 [2]int = arr[1]
		fmt.Print(arr1)
		var a int = arr[1][0]
		fmt.Print(a)*/

	/*	//切片
		arr := [...]int{32, 54, 76, 12, 76}
		fmt.Print(arr, arr[1:4]) //从数组中生成切片
		//结果：[32 54 76 12 76] [54 76 12]
		fmt.Print(arr[:])
		//空切片
		fmt.Print(arr[0:0])
		//自定义切片,添加元素
		var nameLists []string
		nameLists = append(nameLists, "xxj", "kk")
		fmt.Print(nameLists)
		//切片判空
		fmt.Println(nameLists == nil)*/

	/*	//make声明切片
		ints := make([]int, 3, 10)
		fmt.Print(len(ints)) //3
		ints = append(ints, 9, 6, 3)
		fmt.Print(ints) //[0 0 0 9 6 3]*/

	/*	//切片初始化
		var s []int            //切片声明，len = cap = 0
		s = []int{}            //初始化，len = cap = 0
		s = make([]int, 3)     //初始化，len = cap = 3
		s = make([]int, 3, 5)  //初始化，len = 3, cap = 5
		s = []int{1, 2, 34, 5} //len = cap =4
		s2d := [][]int{
			{1}, {2, 4}, //二维数组各行的列数相等，但二维切片各行的len可以不等
		}*/

	coef_cap()
}

// append函数扩容策略
func coef_cap() {
	s := make([]int, 0, 5)
	nowCap := cap(s)
	for i := 0; i < 100; i++ {
		s = append(s, i)
		afterCap := cap(s)
		if afterCap > nowCap {
			fmt.Printf("cap %d ---> %d\n", nowCap, afterCap)
			nowCap = afterCap
		}
	}
}
