package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Используем bufio для корректного чтения русских букв
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Выйти? (да/нет): ")

		// Читаем всю строку
		answer, _ := reader.ReadString('\n')

		// Убираем символы переноса строки и пробелы
		answer = strings.TrimSpace(answer)
		answer = strings.ToLower(answer)

		if answer == "да" {
			fmt.Println("Программа завершена.")
			break
		} else if answer == "нет" {
			// Продолжаем цикл
		} else {
			fmt.Println("Пожалуйста, введите только 'да' или 'нет'")
		}
	}
}
