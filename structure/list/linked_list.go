package list

import (
	"errors"
)

type LinkedListElement struct {
	next *LinkedListElement
	prev *LinkedListElement
	data interface{}
}

type LinkedList struct {
	length int
	head   *LinkedListElement
	tail   *LinkedListElement
}

func NewLinkedList(elements ...interface{}) LinkedList {

	l := LinkedList{0, nil, nil}

	node := &LinkedListElement{}

	for key, data := range elements {
		if key == 0 {
			l.head = &LinkedListElement{
				next: nil,
				prev: nil,
				data: data,
			}
			node = l.head
		} else {
			node.next = &LinkedListElement{
				next: nil,
				prev: node,
				data: data,
			}
			node = node.next
		}
		l.length++
	}

	l.tail = node

	return l
}

func (l *LinkedList) checkIndex(index int) error {

	if index < 0 {
		return errors.New("index不能小于0")
	}

	if index >= (l.length) {
		return errors.New("index已经超出范围")
	}

	return nil
}

func (l *LinkedList) locateElement(index int) *LinkedListElement {
	if index == 0 {
		return l.head
	} else if index == l.length-1 {
		return l.tail
	} else {
		node := l.head
		for i := 0; i <= index; i++ {
			if i != index {
				node = node.next
			}
		}
		return node
	}
}

func (l *LinkedList) Clear() {
	l.length = 0
	l.head = nil
	l.tail = nil
}

func (l *LinkedList) IsEmpty() bool {
	if l.length == 0 {
		return true
	} else {
		return false
	}
}

func (l *LinkedList) Length() int {
	return l.length
}

func (l *LinkedList) Append(data interface{}) {
	newNode := &LinkedListElement{prev: l.tail, next: nil, data: data}
	l.tail = newNode
	l.length++
}

func (l *LinkedList) Insert(index int, data interface{}) error {

	if err := l.checkIndex(index); err != nil {
		return err
	}

	node := l.locateElement(index)

	newNode := &LinkedListElement{node, nil, data}

	if index == 0 {
		l.head = newNode
	} else if index == (l.length - 1) {
		newNode.prev = node.prev
		node.prev.next = newNode
	} else {
		newNode.prev = node.prev
	}

	node.prev = newNode

	l.length++

	return nil
}

func (l *LinkedList) Delete(index int) error {

	if err := l.checkIndex(index); err != nil {
		return err
	}

	node := l.locateElement(index)

	if index == 0 {
		node = node.next
		node.next.prev = nil
	} else if index == l.length-1 {
		node = node.prev
		node.prev.next = nil
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}

	l.length--

	return nil
}

func (l *LinkedList) Update(index int, data interface{}) error {

	if err := l.checkIndex(index); err != nil {
		return err
	}

	node := l.locateElement(index)
	node.data = data

	return nil
}

func (l *LinkedList) Get(index int) (interface{}, error) {

	if err := l.checkIndex(index); err != nil {
		return nil, err
	}

	node := l.locateElement(index)

	return node.data, nil
}

func (l *LinkedList) Scan() []interface{} {
	var result []interface{}

	node := l.head

	for node != nil {
		result = append(result, node.data)
		node = node.next
	}

	return result
}
