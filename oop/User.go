package main

import "fmt"

//类属性定义
type User struct {
	first_name 	string
	last_name 	string
	age 		int
}

//类方法定义
func (user *User) GetName() string {
	return user.first_name + " " + user.last_name
}

func (user *User) GetAge() int  {
	return user.age
}

func (user *User) SetFirstName(firstname string)  {
	user.first_name = firstname
}

func (user *User) SetLastName(lastname string)  {
	user.last_name = lastname
}

func (user *User) SetAge(age int)  {
	user.age = age
}

func main()  {

	user := User{first_name:"S", last_name:"S", age:36}

	user.SetFirstName("Stranley")
	user.SetLastName("Sun")
	user.SetAge(36)

	fmt.Println(user)

}
