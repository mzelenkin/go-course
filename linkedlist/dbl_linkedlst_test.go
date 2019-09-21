package linkedlist

import (
	"testing"
)

func TestPushBack(t *testing.T) {
	list := NewDblLinkedList()
	list.PushBack(1)
	list.PushBack(2)

	if list.First().Value() != 1 && list.First().Next().Value() != 2 {
		t.Errorf("Ошибка при получении элементов списка в нужной последовательности")
	}
}

func TestPushFront(t *testing.T) {
	list := NewDblLinkedList()
	list.PushFront(1)
	list.PushFront(2)

	if list.First().Value() != 2 && list.First().Next().Value() != 1 {
		t.Errorf("Ошибка при получении элементов списка в нужной последовательности")
	}
}

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

func TestGetTail(t *testing.T) {
	list := NewDblLinkedList()
	list.PushBack(1)
	list.PushBack(2)

	if list.Last().Value() != 2 {
		t.Errorf("Ошибка при получении элементов списка в нужной последовательности")
	}
}

func TestRemove(t *testing.T) {
	list := NewDblLinkedList()
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.PushBack(4)

	list.Remove(list.Last().Prev())

	if list.Last().Value() != 4 || list.Last().Prev().Value() != 2 {
		t.Errorf("Ошибка при получении элементов списка в нужной последовательности")
	}
}
