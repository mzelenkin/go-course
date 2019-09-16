package freqanalysis

import (
	"bufio"
	"io"
	"sort"
	"strings"
	"unicode"
)

// Структура с данными статистики по слову
type wordFreq struct {
	Word      string
	Frequency int
}

// Analyze - анализирует частоту слов в потоке
func Analyze(r io.Reader) (ret []string, err error) {
	words, err := countWords(r)
	if err != nil {
		return
	}

	sortedWords := sortByFreq(words)
	ret = getTop(sortedWords, 10)

	return
}

// countWords подсчитывает количество слов и возвращает результат в виде map[слово]кол-во
func countWords(r io.Reader) (stat []wordFreq, err error) {
	var textRune rune
	var buff = strings.Builder{}
	var words = map[string]int{}

	buffReader := bufio.NewReader(r)
	for {
		textRune, _, err = buffReader.ReadRune()
		if err == io.EOF {
			// При EOF записываем слово и прерываем цикл
			word := strings.ToLower(buff.String())
			words[word] = words[word] + 1
			err = nil
			break
		} else if err != nil {
			// Если что-то идет не так просто заканчиваем работу
			// Вызывающий метод получит err и разберется что с ним делать
			return
		}

		if isEndOfWord(textRune) {
			// если это разделитель слов,
			// преобразуем в нижний регистр и добавляем в map
			if buff.Len() > 0 {
				word := strings.ToLower(buff.String())
				words[word] = words[word] + 1
				buff.Reset()
			}
		} else {
			buff.WriteRune(textRune)
		}
	}

	return convertMapToSlice(words), nil
}

// isEndOfWord проверяет является ли символ признаком конца слова
func isEndOfWord(r rune) bool {
	return !(unicode.IsLetter(r) || unicode.IsDigit(r))
}

// convertMapToSlice конвертирует map слово-частота в слайс
func convertMapToSlice(words map[string]int) []wordFreq {
	var ret []wordFreq

	for word, freq := range words {
		ret = append(ret, wordFreq{
			Word:      word,
			Frequency: freq,
		})
	}

	return ret
}

// sortByFreq сортирует слова по частоте использования
func sortByFreq(words []wordFreq) []wordFreq {
	sort.Slice(words, func(i, j int) bool {
		return words[i].Frequency > words[j].Frequency
	})

	return words
}

// getTop возвращает топ (max) часто используемых слов
func getTop(sortedWords []wordFreq, max int) (ret []string) {
	ret = []string{}
	for idx, w := range sortedWords {
		ret = append(ret, w.Word)
		if idx > max {
			break
		}
	}

	return
}
