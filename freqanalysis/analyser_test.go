package freqanalysis

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

// TestAnalyze тестирует основную функцию пакета Analyze - для вывода самых часто встречаемых слов
func TestAnalyze(t *testing.T) {
	r, err := os.Open("chukigek.txt")
	defer r.Close() // На случай если что-то идет не так

	if err != nil {
		t.Fatal("Can't open file agent008.txt")
	}

	result, err := Analyze(r)
	if err != nil {
		t.Fatalf("Read error: %s", err.Error())
	}

	if len(result) < 10 {
		t.Error("Количество возвращенных слов не равно кол-ву слов для сравнения")
		t.FailNow()
	}

	ethalon := []string{"и", "на", "в", "не", "что", "он", "а", "с", "гек", "мать", "чук", "то"}
	if !reflect.DeepEqual(ethalon, result) {
		t.Errorf("Слайсы результата и эталонный не совпадают\n Результат: %#v\n Эталон: %#v\n", result, ethalon)
	}
}

// TestIsEndOfWord это тест функции определения конца слова
func TestIsEndOfWord(t *testing.T) {
	// Список разделителей
	wordsDelimiters := []rune{' ', ',', '.', 13, 10}

	for _, symbol := range wordsDelimiters {
		if !isEndOfWord(symbol) {
			t.Errorf("Символ с кодом %d не воспринялся как конец слова", symbol)
		}
	}
}

// Тест функции подсчета слов
func TestCountWords(t *testing.T) {
	ethalons := map[string]int{
		"аптека": 3,
		"фонарь": 2,
		"улица":  1,
	}

	r := strings.NewReader("аптека улица фонарь Аптека фонарь аптека")
	results, err := countWords(r)
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(results) != len(ethalons) {
		t.Fatal("Количество возвращенных слов не равно кол-ву слов для сравнения")
	}

	for _, wordStat := range results {
		frequency, ok := ethalons[wordStat.Word]
		if !ok {
			t.Errorf("Слово '%s' не найдено в списке для проверки", wordStat.Word)
		}

		if wordStat.Frequency != frequency || !ok {
			t.Errorf("Частота для слова '%s' вычисленна на верно - %d (должно быть %d)",
				wordStat.Word, wordStat.Frequency, frequency)
		}
	}
}

// Тест сортировки слов по частоте
func TestSortByFrequency(t *testing.T) {
	r := strings.NewReader("аптека улица фонарь Аптека фонарь аптека")
	words, err := countWords(r)
	if err != nil {
		t.Fatal(err.Error())
	}

	ethalon := []wordFreq{
		wordFreq{Word: "аптека", Frequency: 3},
		wordFreq{Word: "фонарь", Frequency: 2},
		wordFreq{Word: "улица", Frequency: 1},
	}

	result := sortByFreq(words)

	for index, _ := range ethalon {
		if result[index].Word != ethalon[index].Word ||
			result[index].Frequency != ethalon[index].Frequency {
			t.Errorf("Элементы ответа не совпадают. Результат '%s': %d, должно быть '%s':%d ",
				result[index].Word, result[index].Frequency, ethalon[index].Word, ethalon[index].Frequency)
		}
	}
}
