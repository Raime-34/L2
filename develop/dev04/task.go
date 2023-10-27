package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	for _, v := range getAnagrams([]string{"пятак", "пятка", "тяпка"}) {
		fmt.Println(v)
	}
}

func getAnagrams(words []string) (result map[string]map[string]struct{}) {

	result = make(map[string]map[string]struct{})

	for _, word := range words {

		runes := strings.Split(word, "")
		sort.Strings(runes)
		sortedWord := strings.Join(runes, "")

		if result[sortedWord] == nil {
			result[sortedWord] = make(map[string]struct{})
		}

		result[sortedWord][word] = struct{}{}

	}

	for k, v := range result {
		if len(v) == 1 {
			delete(result, k)
		}
	}

	return
}
