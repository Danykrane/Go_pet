package main

import "fmt"

func main() {
	var val5 = 15
	/*
		if val5 { //нельзя в if записывать int
			fmt.Println("Same")
		}
	*/
	if val5 != 0 { //нельзя в if записывать int
		fmt.Println("more than 0")
	} else {
		fmt.Println("no more than 0")
	}

}
