package logic

import (
	"bufio" // Пакет для буферизованного ввода
	"os"    // Пакет для работы с операционной системой
)

// ReadInput считывает текст из stdin и возвращает его как слайс слов
func ReadInput() ([]string, error) {
	// Создаем новый сканер для чтения из стандартного ввода
	scanner := bufio.NewScanner(os.Stdin)
	// Устанавливаем режим разбора текста: разбиваем ввод по словам
	scanner.Split(bufio.ScanWords)

	// Инициализируем пустой слайс для хранения слов
	var words []string
	// Построчно читаем ввод, добавляя каждое слово в слайс
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	// Проверяем наличие ошибок во время чтения
	if err := scanner.Err(); err != nil {
		// Если ошибка есть, возвращаем её
		return nil, err
	}

	// Возвращаем слайс слов и отсутствие ошибок
	return words, nil
}
