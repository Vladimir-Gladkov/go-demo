package main

import "fmt"

func main() {
	var n int

	fmt.Printf("Введите число: ")
	fmt.Scan(&n)

	if n > 0 {
		fmt.Println("Прямой отсчет ")
		for count := 1; count <= n; count++ {
			fmt.Println(count)
		}
	} else if n < 0 {
		fmt.Println("Обратный отсчет ")
		for count := -1; count >= n; count-- {
			fmt.Println(count)
		}
	} else {
		fmt.Println("Вы ввели 0!")
	}

	fmt.Println("Готово!")

}
