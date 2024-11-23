package logic

import (
	"bufio"
	"os"
)

// ReadInput считывает текст из stdin и возвращает его как слайс слов words и nil, если нет ошибок
func ReadInput() ([]string, error) {
	scanner := bufio.NewScanner(os.Stdin) // создаём объект для постраничного чтения данных из потока

	scanner.Split(bufio.ScanWords) // этому объекту задаём режим разделения на отдельные слова

	var words []string

	for scanner.Scan() { // читает элементы объекта
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}
