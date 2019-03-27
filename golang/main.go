package main

import (
	"reflect"
	"fmt"
)
type temp int

func main() {
	// a := 1
	var b temp = 1

	// t := reflect.TypeOf(a)
	var a int
	a = reflect.ValueOf(b).Interface()
	fmt.Println(a)

	
}