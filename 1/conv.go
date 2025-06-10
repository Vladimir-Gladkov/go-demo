package main

import "fmt"

func main() {
	var dollar, rub float64

	fmt.Print("Введите курс Доллара: ")
	fmt.Scan(&dollar)

	fmt.Print("Введите Сумма в Рублях: ")
	fmt.Scan(&rub)

	a := rub / dollar
	fmt.Printf("У Вас есть: %.2f долларов", a)
}
