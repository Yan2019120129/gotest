package model

type User struct {
	name string
	age  int
	sex  string
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) SetName(Name string) {
	u.name = Name
}

func (u *User) GetAge() int {
	return u.age
}

func (u *User) SetAge(Age int) {
	u.age = Age
}

func (u *User) GetSex() string {
	return u.sex
}

func (u *User) SetSex(sex string) {
	u.sex = sex
}
