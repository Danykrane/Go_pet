package main

import "fmt"

func countdig(number int) int {
	var count_dig int = 0
	for number != 0 {
		number /= 10
		count_dig++
	}
	return count_dig
}

func reverseFirst(number int) {
	col := countdig(number)
	revn := 0
	for i := 0; i < col; i++ {
		temp := number % 10
		number /= 10
		fmt.Println("Последний разряд - ", temp)
		revn += temp
		revn *= 10
	}
	revn /= 10
	fmt.Println("Реверсированное число: ", revn)
}

func reverseNumber(num int) int {
	res := 0
	for num > 0 {
		remainder := num % 10
		res = (res * 10) + remainder
		num /= 10
	}
	return res
}

func palindrome(val1 int, val2 int) {
	if val1 == val2 {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Не Палиндром")
	}
}

// 12345678

// 8 * 10 + 7
// 87 * 10 + 7

func main() {
	sum := 0
	for i := 1; i < 10; i++ {
		sum += i
		fmt.Println(sum, "->", i+1)
	}
	fmt.Println(sum)

	// цикл while
	n := 0
	for n != 6 {
		fmt.Print(n, " ")
		n += 1
	}

	// проверка палиндрома
	number := 123321

	rev := reverseNumber(number)
	palindrome(number, rev)

}
