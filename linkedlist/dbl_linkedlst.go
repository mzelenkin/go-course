package linkedlist

type DblLinkedList struct {
	front *Item
	back  *Item
}

// Конструктор
func NewDblLinkedList() DblLinkedList {
	return DblLinkedList{}
}

// Len возвращает длину списка
// Конечно можно было бы просто завести поле len int в нашей структуре
// и изменять его при добавлении и удалении элемента, а здесь возвращать его значение
// Но я думаю в учебных целях интереснее пройтись по списку, хотя и медленнее O(n)
func (d *DblLinkedList) Len() (len int) {
	// Если front не указывает на элемент, то длина списка 0
	if d.front == nil {
		return 0
	}

	len = 1         // потому что последний элемент не посчитается в цикле
	item := d.front // текущий Item берем с головы
	for item != d.back {
		len++
		item = item.next
	}

	return len
}

// First возвращает первый элемент списка, либо nil если список пуст
func (d DblLinkedList) First() Item {
	return *d.front
}

// Last возвращает последний элемент списка, либо nil если список пуст
func (d DblLinkedList) Last() Item {
	return *d.back
}

// PushFront добавляет элемент v в начало списка
func (d *DblLinkedList) PushFront(v interface{}) {
	// Даже если d.front = nil, это не сильно меняет ситуацию
	current := d.front
	d.front = &Item{next: current, value: v}

	if current != nil {
		current.previous = d.front
	}

	// Корректируем back если список был пуст
	if d.back == nil {
		d.back = d.front
	}
}

// PushBack добавляет элемент v в конец списка
func (d *DblLinkedList) PushBack(v interface{}) {
	lastItem := d.back

	if lastItem == nil {
		// Если список пуст
		d.back = &Item{value: v}
		d.front = d.back
	} else {
		// Если в back что-то есть, т.е. список не пуст
		newItem := Item{previous: lastItem, value: v}
		lastItem.next = &newItem // Пользуемся тем что Item в нашем модуле. Эти поля не доступны пользователю модуля
		d.back = &newItem
	}
}

// Remove удаляет элемент v из списка
func (d *DblLinkedList) Remove(i Item) {

	if i.Prev() == nil {
		// Первый элемент
		// Просто двигаем front в сторону back
		d.front = i.Next()
		if i.Next() == nil {
			// Если это был последний элемент, корректируем back
			d.back = nil
		}
	} else if i.Next() == nil {
		// Последний элемент
		// Просто двигаем back в сторону front
		d.back = i.Prev()
		if i.Prev() == nil {
			// Если это был последний элемент, корректируем front
			d.front = nil
		}
	} else {
		// Где-то в списке
		// Берем next предыдущего элемента и записываем ссылку из next переданного элемента
		// и наоборот в prev следующего записываем Prev из переданного
		// В общем, связываем наш предыдущий элемент и наш следующий
		i.Prev().next = i.Next()
		i.Next().previous = i.Prev()
	}
}
