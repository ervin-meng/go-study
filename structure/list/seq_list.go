package list

import (
	"errors"
)

type SeqList struct {
	elements []interface{}
	length   int
}

func NewSeqList(elements ...interface{}) SeqList {

	l := SeqList{nil, 0}

	for _, data := range elements {
		l.elements = append(l.elements, data)
		l.length++
	}

	return l
}

func (l *SeqList) checkIndex(index int) error {

	if index < 0 {
		return errors.New("index不能小于0")
	}

	if index >= (l.length) {
		return errors.New("index已经超出范围")
	}

	return nil
}

func (l *SeqList) Clear() {
	l.length = 0
	l.elements = nil
}

func (l *SeqList) IsEmpty() bool {
	if l.length == 0 {
		return true
	} else {
		return false
	}
}

func (l *SeqList) Length() int {
	return l.length
}

func (l *SeqList) Append(data interface{}) {
	l.elements = append(l.elements, data)
	l.length++
}

func (l *SeqList) Insert(index int, data interface{}) error {

	if err := l.checkIndex(index); err != nil {
		return err
	}

	newElements := make([]interface{}, 1+l.length)

	copy(newElements, l.elements[:index])
	copy(newElements[index:], []interface{}{data})
	copy(newElements[index+1:], l.elements[index:])

	l.elements = newElements
	l.length++

	return nil
}

func (l *SeqList) Delete(index int) error {

	if err := l.checkIndex(index); err != nil {
		return err
	}

	l.elements = append(l.elements[:index], l.elements[index+1:]...)
	l.length--

	return nil
}

func (l *SeqList) Update(index int, data interface{}) error {

	if err := l.checkIndex(index); err != nil {
		return err
	}

	l.elements[index] = data

	return nil
}

func (l *SeqList) Get(index int) (interface{}, error) {

	if err := l.checkIndex(index); err != nil {
		return nil, err
	}

	return l.elements[index], nil
}
