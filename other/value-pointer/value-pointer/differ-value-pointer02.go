package value_pointer

// Animal 新增接口
type Animal interface {
	GetName() string
	GetAge() int
}

type Person02 struct {
	Name string
	Age  int
}

// GetName 值接受者
func (p Person02) GetName() string {
	return p.Name
}

// GetAge 指针接受者
func (p *Person02) GetAge() int {
	return p.Age
}
