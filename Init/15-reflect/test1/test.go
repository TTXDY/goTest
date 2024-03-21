package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Age  int
	Name string
	Sex  int
}

func (u *User) Call() {
	println(u.Name, " is calling...")
}

func main() {
	user := User{Name: "jay", Sex: 1, Age: 18}
	//println(reflect.TypeOf(user.Age).Name())
	DoFiled(user)
}

func DoFiled(input interface{}) {
	inputType := reflect.TypeOf(input)
	println(inputType.Name())

	inputValue := reflect.ValueOf(input)
	// 此处需要fmt.Println
	fmt.Println(inputValue)

	// 结构体中的每个元素
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Println("name: ", field.Name, "; type: ", field.Type, "; value: ", value)
	}
}
