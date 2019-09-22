package linkedlist

// Item - элемент списка
type Item struct {
	value    interface{}
	previous *Item
	next     *Item
}

// Value возвращает значение элемента
func (i Item) Value() interface{} {
	return i.value
}

// Prev возвращает ссылку на предыдущий элемент
func (i Item) Prev() *Item {
	return i.previous
}

// Next возвращает ссылку на следующий элемент
func (i Item) Next() *Item {
	return i.next
}
