package demo

type User struct {
	Id   int
	Name string
}

type option func(*User)

func (u *User) Option(opts ...option) {
	for _, opt := range opts {
		opt(u)
	}
}

func WithId(id int) option {
	return func(u *User) {
		u.Id = id
	}
}

func WithName(name string) option {
	return func(u *User) {
		u.Name = name
	}
}
