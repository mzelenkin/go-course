package linkedlist

import (
	"testing"
)

var testValues = []int{1, 2, 3, 4, 5, 6}

func TestPushBackSingleItem(t *testing.T) {
	// Проверка только для 1 элемента
	list := NewDblLinkedList()
	list.PushBack(1)
	if list.front == nil || list.back == nil || list.back != list.front {
		t.Errorf("list.front не ссылается на 1 элемент")
	}
}

// Проверка для 2 элементов
func TestPushBackTwoItems(t *testing.T) {
	list := NewDblLinkedList()
	list.PushBack(1)
	list.PushBack(2)

	if list.First().Value() != 1 && list.First().Next().Value() != 2 {
		t.Errorf("Ошибка при получении элементов списка в нужной последовательности")
	}

	// т.к. у нас только 2 элемента, то проверим как работают ссылки front/back
	if list.First().next != list.back {
		t.Errorf("list.back не ссылается на 2 элемент")
	}
	if list.back.previous != list.front {
		t.Errorf("list.front не ссылается на 1 элемент")
	}
}

// Произвольное кол-во элементов
func TestPushBack(t *testing.T) {
	list := NewDblLinkedList()
	for _, v := range testValues {
		list.PushBack(v)
	}

	currentItem := list.front
	for _, v := range testValues {
		if currentItem.Value() != v {
			t.Errorf("Нарушена последовательность элементов в списке")
			break
		}

		currentItem = currentItem.Next()
	}
}

// Проверка PushFront только для 1 элемента
func TestPushFrontSingleItem(t *testing.T) {
	list := NewDblLinkedList()
	list.PushFront(1)
	if list.front == nil || list.back == nil || list.back != list.front {
		t.Errorf("list.back и push.front не ссылается на один и тот же элемент")
	}
}

func TestPushFront(t *testing.T) {
	list := NewDblLinkedList()
	list.PushFront(1)
	list.PushFront(2)

	if list.First().Value() != 2 && list.First().Next().Value() != 1 {
		t.Errorf("Ошибка при получении элементов списка в нужной последовательности")
	}

	// т.к. у нас только 2 элемента, то проверим как работают ссылки front/back
	if list.First().next != list.back {
		t.Errorf("list.back не ссылается на 2 элемент")
	}
	if list.back.previous != list.front {
		t.Errorf("list.front не ссылается на 1 элемент")
	}

	//------------------------------
	// Произвольное кол-во элементов
	list3 := NewDblLinkedList()
	for _, v := range testValues {
		list3.PushFront(v)
	}
	// Важно учесть тот факт, что список будет задом наперед
	currentItem := list3.front
	lenTestValues := len(testValues)
	for i := range testValues {
		val := testValues[lenTestValues-i-1]
		if currentItem.Value() != val {
			t.Errorf("Нарушена последовательность элементов в списке")
		}

		currentItem = currentItem.Next()
	}
}

// Тестирует одновременное добавление назад и вперед
func TestPushFrontAndBack(t *testing.T) {
	list := NewDblLinkedList()
	list.PushFront(1)
	list.PushFront(2)
	list.PushBack(3)

	if list.First().Value() != 2 &&
		list.First().Next().Value() != 1 &&
		list.First().Next().Value() != 3 {
		t.Errorf("Ошибка при получении элементов списка в нужной последовательности")
	}
}

func TestRemove(t *testing.T) {
	// Создаем список и заполняем его значениями
	list := NewDblLinkedList()
	for _, v := range testValues {
		list.PushBack(v)
	}

	// Удаляем предпоследний элемент
	list.Remove(*list.Last().Prev())

	// Вычисляем длинну тестовых значений (понадобится ниже)
	testValuesLen := len(testValues)

	// Проверям список
	currentItem := list.front
	for k, v := range testValues {

		// Предпоследний элемент игнорируем (мы его выше удалили)
		if k == testValuesLen-2 {
			continue
		}

		if currentItem.Value() != v {
			t.Errorf("Нарушена последовательность элементов в списке")
			break
		}

		currentItem = currentItem.Next()
	}

	if list.Len() != testValuesLen-1 {
		t.Errorf("Ошибка при подсчете длины != %d", testValuesLen-1)
	}
}

// Тест удаления последнего элемента списка
func TestRemoveLastItem(t *testing.T) {
	list := NewDblLinkedList()

	list.PushBack(1)
	list.Remove(list.Last())

	// В итоге здесь наш список пустой
	if list.front != nil || list.back != nil {
		t.Errorf("Неверное отрабатывается удаление последнего элемента")

	}

	if list.Len() != 0 {
		t.Errorf("Ошибка при подсчете длины != 0")
	}
}

// Тест удаления крайнего элемента списка
func TestRemoveLastOfTwoItem(t *testing.T) {
	list := NewDblLinkedList()

	list.PushBack(1)
	list.PushBack(2)
	list.Remove(list.Last())

	// В итоге имеет 1 элемент (1)
	// front и back должны ссылаться на один и тот же элемент
	if list.front == nil || list.back == nil || list.back != list.front {
		t.Errorf("list.front и list.back не ссылаются на один и тот же элемент")
	}

	if list.Len() != 1 {
		t.Errorf("Ошибка при подсчете длины != 0")
	}
}
