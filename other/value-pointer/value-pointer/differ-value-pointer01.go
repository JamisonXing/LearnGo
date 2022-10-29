package value_pointer

type Person01 struct {
	Name string
	Age  int
}

// GetName 值接受者
func (p Person01) GetName() string {
	return p.Name
}

// GetAge 指针接受者
func (p *Person01) GetAge() int {
	return p.Age
}
