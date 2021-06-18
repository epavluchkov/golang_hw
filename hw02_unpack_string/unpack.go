package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var sb strings.Builder

	var prev, cur rune

	for _, cur = range s {
		// Проверяем, что строка не начинается с цифры и не содержит более одной цифры подряд
		if unicode.IsDigit(cur) && (unicode.IsDigit(prev) || prev == 0) {
			return "", ErrInvalidString
		}

		// Записываем предыдущий символ в целевую строку (если предыдущий символ не цифра):
		// если текущий символ не цифра - один раз
		// если текущий симовл цира - столько раз, какая у нас цифра в текущем символе.
		if !unicode.IsDigit(prev) && prev != 0 {
			if !unicode.IsDigit(cur) {
				sb.WriteRune(prev)
			} else {
				i, err := strconv.Atoi(string(cur))
				if err != nil {
					return "", err
				}
				sb.WriteString(strings.Repeat(string(prev), i))
			}
		}

		prev = cur
	}

	// Добавляем последний символ исходной строки, если он не цифра
	// и если исходная строка не пустая
	if !unicode.IsDigit(cur) && len(s) > 0 {
		sb.WriteRune(cur)
	}

	return sb.String(), nil
}
