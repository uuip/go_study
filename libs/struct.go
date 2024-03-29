package study

type User struct {
	Name string
	Age  int32
}

func (u *User) MyAge() int32 {
	return u.Age
}

func (u *User) SetAge(age int32) {
	u.Age = age
}
func (u *User) SetName(name string) {
	u.Name = name
}

type UpdateName interface {
	SetName(name string)
}
type UpdateAge interface {
	SetAge(age int32)
}
type Action interface {
	UpdateName
	UpdateAge
}
