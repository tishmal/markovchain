package main

import (
	"flag"              // Пакет для обработки флагов командной строки
	"fmt"               // Пакет для форматированного вывода
	"markovchain/logic" // Логика алгоритма Маркова (ваша собственная реализация)
	"os"                // Пакет для работы с операционной системой
	"strings"           // Пакет для работы со строками
)

func main() {
	// Считываем аргументы командной строки
	args := os.Args
	// Если аргументов меньше 2 (отсутствует входной текст), выводим ошибку и завершаем программу
	if len(args) < 2 {
		fmt.Println("Error: no input text")
		os.Exit(1)
	}

	// Определяем флаги командной строки
	wordCount := flag.Int("w", 100, "Maximum number of words to generate") // Флаг для указания количества слов
	prefix := flag.String("p", "", "Starting prefix")                      // Флаг для начального префикса
	prefixLength := flag.Int("l", 2, "Prefix length (default 2)")          // Флаг для длины префикса
	showHelp := flag.Bool("help", false, "Show usage information")         // Флаг для отображения справки

	// Разбираем флаги
	flag.Parse()

	// Если установлен флаг помощи, показываем справочную информацию и выходим
	if *showHelp {
		logic.Help() // Функция, которая выводит информацию о программе
		return
	}

	// Проверяем допустимость значения флага wordCount
	if *wordCount < 0 || *wordCount > 10000 {
		fmt.Println("Error: word count must be between 0 and 10,000") // Ошибка, если количество слов выходит за пределы
		return
	}

	// Проверяем допустимость значения флага prefixLength
	if *prefixLength < 0 || *prefixLength > 5 {
		fmt.Println("Error: prefix length must be between 0 and 5") // Ошибка, если длина префикса выходит за пределы
		return
	}

	// Считываем входной текст с помощью логики, реализованной в пакете logic
	words, err := logic.ReadInput()
	if err != nil {
		fmt.Println("Error reading input:", err) // Если произошла ошибка при чтении, выводим её
		return
	}

	// Проверяем, что текст не пустой
	if len(words) == 0 {
		fmt.Println("Error: no input text") // Ошибка, если текст пустой
		return
	}

	// Генерируем текст с использованием алгоритма Маркова
	generatedText, err := logic.GenerateMarkovChain(words, *wordCount, *prefix, *prefixLength)
	if err != nil {
		fmt.Println("Error:", err) // Если произошла ошибка в процессе генерации, выводим её
		return
	}

	// Выводим сгенерированный текст, объединяя слова через пробел
	fmt.Println(strings.Join(generatedText, " "))
}
