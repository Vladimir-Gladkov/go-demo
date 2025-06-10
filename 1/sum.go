package main

import "fmt"

func main() {
	var x, sum int

	fmt.Print("Введите число: ")
	fmt.Scan(&x)

	for i := 1; i <= x; i++ {
		sum += i
	}
	fmt.Printf("Сумма от 1 до %d = %d\n", x, sum)
}
