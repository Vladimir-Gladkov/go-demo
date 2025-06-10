package main

import "fmt"

func main() {
	var num int

	fmt.Printf("Введите число:  ")
	fmt.Scan(&num)
	if num > 0 {
		fmt.Println("Положительное число")
	} else if num < 0 {
		fmt.Println("Отричательное число")
	} else {
		fmt.Println("Ноль")
	}
	if num%2 == 0 {
		fmt.Println("Четное")
	} else {
		fmt.Println("Нечетное")
	}
}
