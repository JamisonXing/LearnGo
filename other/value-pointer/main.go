package main

import "value-pointer/value-pointer"

func main() {

	//diff-value_pointer00
	//定义一个值类型
	t := value_pointer.Person{
		Name: "xxj",
		Age:  20,
	}
	//调用值方法
	println(t.GetName())
	//调用指针方法
	println(t.GetAge())
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

	//diff-value_pointer01
	//定义一个指针类型
	t1 := &value_pointer.Person01{
		Name: "xxj",
		Age:  20,
	}
	//调用值方法
	println(t1.GetName())
	//调用指针方法
	println(t1.GetAge())
	/*
		xxj
		20

		这段代码告诉我们，指针类型的变量，
		可以调用该类型的值接收者方法，也可
		以调用指针接收者方法。

		看起来好像两者对等的，并没有差别。那么二者真的没有差别吗？
		只是一种表达形式上的差异？其实不然，如果引入接口类型后，我
		们再来看看：见02。
	*/

	//diff-value_pointer02
	//定义接口变量
	var ani value_pointer.Animal

	//person02实现了animal接口，赋值给了ani变量
	//这里会报错：cannot use value_pointer.Person02{…} (value of type value_pointer.Person02) as type value_pointer.Animal in assignment:
	//	value_pointer.Person02 does not implement value_pointer.Animal (GetAge method has pointer receiver)

	/*ani = value_pointer.Person02{	//这里会报错
		Name: "xxj",
		Age:  30,
	}

	ani.GetName()
	ani.GetAge()*/
	print(ani)

	/*
			为什么会报错呢？错误提醒很明显了：Person 没有实现
			Animal 的 GetName 方法。因为在上面的代码中，我们
			实现 GetName 方法的是 (*Person) 类型。

			但是为什么GetAge 方法不报错呢？那是因为 Go 里边对于
			(Type)Method的方法，会自动让他拥有 (*Type)Method
			方法的能力。

		总结一下，实现接口时有下面的约束：
		如果定义的是 (Type)Method，则该类型会隐式的声明一个 (*Type)Method；
		如果定义的是 (*Type)Method ，则不会隐式什么一个 (Type)Method。

		至于为什么不也隐式申明一个  (Type)Method ，我觉得有一个原因是，
		我们一般采用指针接收者时，方法内部改变的值，接收者本身也会改变，
		那么此时如果隐式有这样一个申明，外部使用值类型时，这个改变就不会生效，
		语义上就会非常奇怪。
	*/

}
