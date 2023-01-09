package main

import "fmt"

func main() {
	fmt.Println()
	var str1 string // ""
	_ = str1
	var str string = "Hello"
	_ = str
	fmt.Println(str)

	str2 := "hello world!"
	fmt.Println(str2)

	var str3 string = "my name is Artyom!"
	fmt.Println(str3[0])   //109 - код первого символа
	fmt.Println(str3[1])   //121
	fmt.Println(str3[2])   //32
	fmt.Println(len(str3)) // 18

	//str[from:to]
	var str4 string = str3[11:17] // Artyom
	var str5 string = str3[:11]   // my name is
	var str6 string = str3[11:]   // Artyom!
	fmt.Println(str4)
	fmt.Println(str5)
	fmt.Println(str6)
	fmt.Println()

	str6 += " " + str5
	fmt.Println(str6)
	//str6[0] = "w" // ошибка
	fmt.Println(str6[0])

	fmt.Println(len("abcde!"))
	fmt.Println(str6)

	str6 = "a" + str6[1:]
	fmt.Println(str6)

}
