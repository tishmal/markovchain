package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	// Чтение входных данных
	scanner := bufio.NewScanner(os.Stdin)
	var text string

	// Считываем весь текст из stdin
	for scanner.Scan() {
		text += scanner.Text() + " "
	}

	// Если текст не был передан, выводим ошибку
	if text == "" {
		fmt.Println("Error: no input text")
		return
	}

	// Разбиваем текст на слова
	words := strings.Fields(text)

	// Создаем модель переходов для Маркова
	// Мапа: (префикс 2 слова) -> список возможных следующих слов
	model := make(map[string][]string)

	// Заполняем модель переходов
	for i := 0; i < len(words)-2; i++ {
		prefix := words[i] + " " + words[i+1]
		nextWord := words[i+2]
		model[prefix] = append(model[prefix], nextWord)
	}

	// Начальный префикс (первые два слова)
	if len(words) < 2 {
		fmt.Println("Error: not enough words to form a prefix")
		return
	}

	// Генерация текста (максимум 100 слов)
	rand.Seed(time.Now().UnixNano())

	prefix := words[0] + " " + words[1]
	result := []string{words[0], words[1]}
	for len(result) < 100 {
		// Находим следующее слово для текущего префикса
		nextWords, exists := model[prefix]
		if !exists {
			break
		}
		// Выбираем случайное следующее слово
		nextWord := nextWords[rand.Intn(len(nextWords))]
		result = append(result, nextWord)

		// Обновляем префикс
		prefix = result[len(result)-2] + " " + result[len(result)-1]
	}

	// Печатаем результат
	fmt.Println(strings.Join(result, " "))
}
