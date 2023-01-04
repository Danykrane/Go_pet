package main

import "fmt"

func main() {

	//стандартное объявление с предопределенным значение
	var v1 int //v1 = 0
	fmt.Println(v1)

	var v2 = 2 // v2 = 2
	fmt.Println(v2)

	//сокращенный вариант без var с автоматическим опредедением типа
	v3 := 5 //var v3 int = 5
	fmt.Println(v3)

	v4 := 4
	//v4 := 4 // нельзя объявлять и инициализировать два раза одну переменную
	v4 = 10
	fmt.Println(v4)

	var v5, v6 int
	v5, v6 = 12, 13
	fmt.Println(v5)
	fmt.Println(v6)

	//замена значений
	v5, v6 = v6, v5
	fmt.Println(v5)
	fmt.Println(v6)

	//блочный режим
	var (
		v7 = 1
		v8 = "str"
	)

	fmt.Println(v8)

	//Механизм игнорирования
	_ = v7
}
