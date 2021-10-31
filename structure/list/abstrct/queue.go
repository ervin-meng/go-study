package abstrct

type IQueue interface {
	Enqueue(data interface{})
	Dequeue() interface{}
	Length() int
	IsEmpty() bool
	Clear()
}
