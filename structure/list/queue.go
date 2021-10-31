package list

type Queue struct {
	elements []interface{}
	length   int
}

func NewQueue(elements ...interface{}) *Queue {
	q := &Queue{nil, 0}
	for _, data := range elements {
		q.Enqueue(data)
	}
	return q
}

func (q *Queue) Enqueue(data interface{}) {
	q.elements = append(q.elements, data)
	q.length++
}

func (q *Queue) Dequeue() interface{} {
	data := q.elements[:1]
	q.elements = q.elements[1:]
	q.length--
	return data
}

func (q *Queue) Length() int {
	return q.length
}

func (q *Queue) IsEmpty() bool {
	return q.length == 0
}

func (q *Queue) Clear() {
	q.elements = nil
	q.length = 0
}
