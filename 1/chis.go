package main

import "fmt"

func main() {
	var x, y, calc int

	fmt.Print("Введите число: ")
	fmt.Scan(&x)

	fmt.Print("Введите число: ")
	fmt.Scan(&y)

	fmt.Print("Введите оперцию (+, -, *, /, %): ")
	fmt.Scan(&calc)

	switch calc {
	case sum := x + y:
		fmt.Println("Сумма: %d\n", sum)
	case diff := x - y:
		fmt.Println("Разность: %d\n", diff)
	case prod := x * y:
		fmt.Println("Произведение: %d\n", prod)
	case qu := x / y :
		fmt.Println("Целочисленное деление: %d\n", qu)
	case rm := x % y :
		fmt.Println("Остаток от деления: %d\n", rm)
	case floatDiv := float64(x) / float64(y) //дробное деление
	fmt.Println("Дробное деление: %.2f\n", floatDiv)
	default:
		fmt.Println("Некорректный ввод")
	}
	

	if y == 0 {
		fmt.Println("Ошибка деления на ноль!")
		return
	}

	                       
	
/*
	fmt.Printf("Сумма: %d\n", sum)
	fmt.Printf("Разность: %d\n", diff)
	fmt.Printf("Произведение: %d\n", prod)
	fmt.Printf("Целочисленное деление: %d\n", qu)
	fmt.Printf("Остаток от деления: %d\n", rm)
	fmt.Printf("Дробное деление: %.2f\n", floatDiv)
*/
}
