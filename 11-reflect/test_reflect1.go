package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this User) Call() {
	fmt.Println(this)
}

func main() {
	user := User{123, "dddd", 23}
	userType := reflect.TypeOf(user)
	userValue := reflect.ValueOf(user)
	fmt.Println("type = ", userType.Name())
	fmt.Println("value = ", userValue)
	for i := 0; i < userType.NumField(); i++ {
		filed := userType.Field(i)
		value := userValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", filed.Name, filed.Type, value)
	}
	for i := 0; i < userType.NumMethod(); i++ {
		m := userType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}
