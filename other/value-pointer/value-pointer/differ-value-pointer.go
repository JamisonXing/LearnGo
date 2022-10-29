package value_pointer

//Go中的值接收者与指针接收者有什么关系与区别，该怎么选？

type Person struct {
	Name string
	Age  int
}

// GetName 值接受者
func (p Person) GetName() string {
	return p.Name
}

// GetAge 指针接受者
func (p *Person) GetAge() int {
	return p.Age
}

/*
/private/var/folders/7d/1n18573d01j7cx_9wk3cc7800000gp/T/GoLand/___go_build_value_pointer
xxj
20
*/

/*
从使用过程看，值类型的变量，可以调用该类型的值接收者方法，
也可以调用指针接收者方法。反之，我们可以定义一个指针类型，
然后看看调用结果：见01
*/
