package list

import (
	"errors"
)

type linkedListElement struct {
	next *linkedListElement
	prev *linkedListElement
	data interface{}
}

type linkedList struct {
	length int
	head   *linkedListElement
	tail   *linkedListElement
}

func NewLinkedList(elements ...interface{}) linkedList {

	l := linkedList{0, nil, nil}

	node := &linkedListElement{}

	for key, data := range elements {
		if key == 0 {
			l.head = &linkedListElement{
				next: nil,
				prev: nil,
				data: data,
			}
			node = l.head
		} else {
			node.next = &linkedListElement{
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

func (l *linkedList) checkIndex(index int) error {

	if index < 0 {
		return errors.New("index不能小于0")
	}

	if index >= (l.length) {
		return errors.New("index已经超出范围")
	}

	return nil
}

func (l *linkedList) locateElement(index int) *linkedListElement {
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

func (l *linkedList) Clear() {
	l.length = 0
	l.head = nil
	l.tail = nil
}

func (l *linkedList) IsEmpty() bool {
	if l.length == 0 {
		return true
	} else {
		return false
	}
}

func (l *linkedList) Length() int {
	return l.length
}

func (l *linkedList) Append(data interface{}) {
	newNode := &linkedListElement{prev: l.tail, next: nil, data: data}
	l.tail = newNode
	l.length++
}

func (l *linkedList) Insert(index int, data interface{}) error {

	if err := l.checkIndex(index); err != nil {
		return err
	}

	node := l.locateElement(index)

	newNode := &linkedListElement{node, nil, data}

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

func (l *linkedList) Delete(index int) error {

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

func (l *linkedList) Update(index int, data interface{}) error {

	if err := l.checkIndex(index); err != nil {
		return err
	}

	node := l.locateElement(index)
	node.data = data

	return nil
}

func (l *linkedList) Get(index int) (interface{}, error) {

	if err := l.checkIndex(index); err != nil {
		return nil, err
	}

	node := l.locateElement(index)

	return node.data, nil
}

func (l *linkedList) Scan() []interface{} {
	var result []interface{}

	node := l.head

	for node != nil {
		result = append(result, node.data)
		node = node.next
	}

	return result
}
