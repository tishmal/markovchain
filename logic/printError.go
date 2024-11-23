package logic

import (
	"fmt"
	"os"
)

// Функция для избежания дублирования кода. Вывод сообщения об конкретной ошибке и выход со статусом 1.
func PrintError(msg string) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}
