package linkedlist

type DblLinkedList struct {
	front *Item
	back  *Item
	len   int
}

// Конструктор
func NewDblLinkedList() DblLinkedList {
	return DblLinkedList{}
}

// Len возвращает длинну списка
func (d *DblLinkedList) Len() int {
	return d.len
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
	current := d.front
	d.front = &Item{next: current, value: v}

	if current == nil {
		d.back = d.front
	} else {
		d.back = current
	}
}

// PushBack добавляет элемент v в конец списка
func (d *DblLinkedList) PushBack(v interface{}) {
	lastItem := d.back
	if lastItem == nil {
		d.back = &Item{value: v}
		d.front = d.back
	} else {
		newItem := Item{previous: lastItem, value: v}
		lastItem.SetNext(&newItem)
		d.back = &newItem
	}
}

// InsertAfter добавляет элемент v после указанного элемента
func (d *DblLinkedList) InsertAfter(v interface{}, i *Item) {

}

// Remove удаляет элемент v из списка
func (d *DblLinkedList) Remove(i *Item) {
	if i.Prev() == nil {
		// Первый элемент
		d.front = i.Next()
	} else if i.Next() == nil {
		// Последний элемент
		d.back = i.Prev()
	} else {
		next := i.Next()
		i.Prev().SetNext(next)
		next.SetPrevious(i.Prev())
	}
}
