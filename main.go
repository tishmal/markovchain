package main

import (
	"flag"
	"fmt"
	"markovchain/logic"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Error: no input text")
		os.Exit(1)
	}

	// Флаги командной строки
	wordCount := flag.Int("w", 100, "Maximum number of words to generate")
	prefix := flag.String("p", "", "Starting prefix")
	prefixLength := flag.Int("l", 2, "Prefix length (default 2)")
	showHelp := flag.Bool("help", false, "Show usage information")

	flag.Parse()

	// Показ справки
	if *showHelp {
		logic.Help()
		return
	}

	// Проверка значений флагов
	if *wordCount < 0 || *wordCount > 10000 {
		fmt.Println("Error: word count must be between 0 and 10,000")
		return
	}

	if *prefixLength < 0 || *prefixLength > 5 {
		fmt.Println("Error: prefix length must be between 0 and 5")
		return
	}

	// Чтение входного текста
	words, err := logic.ReadInput()
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Проверка наличия текста
	if len(words) == 0 {
		fmt.Println("Error: no input text")
		return
	}

	// Генерация текста по алгоритму Маркова
	generatedText, err := logic.GenerateMarkovChain(words, *wordCount, *prefix, *prefixLength)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Вывод результата
	fmt.Println(strings.Join(generatedText, " "))
}
