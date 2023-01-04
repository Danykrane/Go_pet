package main

import "fmt"

func main() {

	var v1 = 12
	_ = v1

	// & - адреса
	// * - полечение значения
	var ptr = &v1
	fmt.Println(&v1)
	fmt.Println(ptr)

	*ptr = 24
	fmt.Println(v1)

}
