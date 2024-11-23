package logic

import (
	"fmt"
	"os"
)

func Validation(help *bool, maxWords *int, prefix *string, prefixLength *int) {
	// Флаг --help
	if *help {
		Help()
	}

	// Проверки на корректность параметров
	if *maxWords < 1 || *maxWords > 10000 {
		PrintError("number of words must be in between 1 and 10000")
	}

	if *prefixLength < 1 || *prefixLength > 5 {
		PrintError("length of prefix must be in between 1 and 5")
	}

	// эта проверка служит для того, чтобы гарантировать, что программа не будет запущена без входных данных (когда она ожидает ввод через конвейер).
	file, _ := os.Stdin.Stat()
	if (file.Mode() & os.ModeCharDevice) != 0 { // если результат побитовой операции не равен нулю, это означает, что стандартный ввод это уст-во символов.
		fmt.Print("Error: no input text\n")
		os.Exit(1)
	}
}
