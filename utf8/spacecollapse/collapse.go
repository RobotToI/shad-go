//go:build !solution

package spacecollapse

import (
	"unicode"
	"unicode/utf8"
)

func CollapseSpaces(input string) string {
	// Оценка максимальной длины результирующей строки
	n := len(input)
	if n == 0 {
		return ""
	}

	// Массив рун для хранения результата
	result := make([]rune, 0, n) // Резервируем место для всех символов

	spaceAdded := false

	for i := 0; i < n; {
		r, size := utf8.DecodeRuneInString(input[i:])
		if r == utf8.RuneError && size == 1 {
			// Некорректная последовательность, добавляем символ замены
			result = append(result, '\ufffd')
			i++
			spaceAdded = false
			continue
		}

		if unicode.IsSpace(r) {
			if !spaceAdded {
				result = append(result, ' ')
				spaceAdded = true
			}
		} else {
			result = append(result, r)
			spaceAdded = false
		}
		i += size
	}

	return string(result)
}
